package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/JIAWea/erpServer/api/erp"
	"github.com/JIAWea/erpServer/config"
	"github.com/JIAWea/erpServer/pkg/utils"
	"github.com/ml444/gkit/errorx"
	log "github.com/ml444/glog"

	"github.com/xuri/excelize/v2"
)

const maxUploadSize = 1024 * 1024 * 10 // 10M

var File = &fileService{}

type fileService struct{}

func (f *fileService) UploadFile(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("file required"))
		return
	}
	defer func() { _ = file.Close() }()

	if header.Size > maxUploadSize {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("file size limit"))
		return
	}

	buf, err := io.ReadAll(file)
	if err != nil {
		log.Error("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("read file error"))
		return
	}

	md5 := utils.EncodeMD5(buf)
	resp := fmt.Sprintf(fmt.Sprintf(`{"md5":"%s","file_name":"%s"}`, md5, header.Filename))

	savePath := filepath.Join(config.DefaultConfig.AssetDir, "detail", md5)
	exist, err := utils.IsPathExist(savePath)
	if err != nil {
		log.Error("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("server error"))
		return
	}
	if exist {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp))
		return
	}

	err = utils.SaveFile(bytes.NewReader(buf), savePath)
	if err != nil {
		log.Error("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("save error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
	return
}

func (f *fileService) ImportExpense(w http.ResponseWriter, r *http.Request) {
	userId := GetUserId(r)

	accIdList, err := dbUserAccount.GetIdListByUserId(r.Context(), userId)
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "未分配可操作的账户")
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		utils.RspBadError(w, "file required")
		return
	}
	defer func() { _ = file.Close() }()

	if header.Size > maxUploadSize {
		utils.RspBadError(w, "file size limit 10M")
		return
	}

	pathDir := filepath.Join(config.DefaultConfig.AssetDir, "import", time.Now().Format("20060102"))
	exist, _ := utils.IsPathExist(pathDir)
	if !exist {
		_ = os.MkdirAll(pathDir, 0777)
	}
	path := filepath.Join(pathDir, fmt.Sprintf("ep_%s_%s", utils.GenUUID(), header.Filename))
	err = utils.SaveFile(file, path)
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "save error")
		return
	}

	excel, err := excelize.OpenFile(path)
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "open excel error")
		return
	}
	defer func() { _ = excel.Close() }()

	rows, err := excel.GetRows("Sheet1")
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "get excel rows error")
		return
	}

	// 第一行为表头
	if len(rows) < 2 {
		log.Error("err:", err)
		utils.RspError(w, "数据不能为空")
		return
	}

	var (
		errMsg      string
		batchRecord []*erp.ModelExpense
		accountMap  = make(map[string]*erp.ModelAccount)
	)

	writeRecord := func(row []string) error {
		date := strings.TrimSpace(row[0])
		uuid := strings.TrimSpace(row[1])
		categoryName := strings.TrimSpace(row[2])
		mark := strings.TrimSpace(row[3])
		money := strings.TrimSpace(row[4])
		accountName := strings.TrimSpace(row[5])
		ticket := strings.TrimSpace(row[6])
		handleBy := strings.TrimSpace(row[7])

		if uuid == "" {
			uuid = utils.GenUUID()
		}

		payAt, err := utils.StrToTime(date)
		if err != nil {
			errMsg = fmt.Sprintf("【%v】该日期格式错误", date)
			log.Error("err:", err)
			return err
		}

		mon, err := strconv.ParseFloat(money, 64)
		if err != nil {
			errMsg = fmt.Sprintf("【%v】该金额格式错误", money)
			log.Error("err:", err)
			return err
		}
		monFen := uint32(mon * 100)

		account, ok := accountMap[accountName]
		if !ok {
			var acc erp.ModelAccount
			err = dbAccount.newScope().
				SetNotFoundErr(erp.ErrNotFoundAccount).
				Eq(dbName, accountName).First(&acc)
			if err != nil {
				if !errorx.IsNotFoundErr(err, erp.ErrNotFoundAccount) {
					return err
				}
				account = &erp.ModelAccount{}
			}
			account = &acc
			accountMap[accountName] = &acc
		}
		if account.Id == 0 {
			errMsg = fmt.Sprintf("【%s】该账户不存在", accountName)
			return errorx.New(erp.ErrNotFoundAccount)
		}

		if !isInSliceUint64(account.Id, accIdList) {
			errMsg = fmt.Sprintf("没有【%s】该账户的操作权限", accountName)
			return errorx.New(erp.ErrNotPermissionForAccount)
		}

		cat, ok := erp.ExpenseCategoryMap[categoryName]
		if !ok {
			errMsg = fmt.Sprintf("没有【%s】该收入类目", categoryName)
			return errorx.New(erp.ErrExpenseCategoryInvalid)
		}

		batchRecord = append(batchRecord, &erp.ModelExpense{
			PayAt:     uint32(payAt.Unix()),
			Uuid:      uuid,
			Category:  cat,
			Mark:      mark,
			PayMoney:  monFen,
			AccountId: account.Id,
			Ticket:    ticket,
			HandleBy:  handleBy,
			UserId:    userId,
		})

		return nil
	}

	isHeader := true
	for _, row := range rows {
		if isHeader {
			isHeader = false
			continue
		}
		if len(row) != 8 {
			utils.RspError(w, "请检查表格列数")
			return
		}

		err = writeRecord(row)
		if err != nil {
			log.Error("err:", err)
			msg := "写入失败"
			if errors.Is(err, errorx.New(erp.ErrExpenseCategoryInvalid)) {
				msg = erp.ErrCodeMap[erp.ErrExpenseCategoryInvalid]
			} else if errors.Is(err, errorx.New(erp.ErrNotPermissionForAccount)) {
				msg = erp.ErrCodeMap[erp.ErrNotPermissionForAccount]
			}
			if errMsg != "" {
				msg = errMsg
			}
			utils.RspError(w, msg)
			return
		}
	}

	if len(batchRecord) > 0 {
		err = dbExpense.newScope().CreateInBatches(batchRecord, 100)
		if err != nil {
			log.Error("err:", err)
			utils.RspError(w, "write record error")
			return
		}
	}

	utils.RspOK(w)
	return
}

func (f *fileService) DownloadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fileName := r.FormValue("file_name")
	md5 := r.FormValue("md5")
	savePath := filepath.Join(config.DefaultConfig.AssetDir, "detail", utils.GetFileName(md5))

	file, err := os.Open(savePath)
	if err != nil {
		log.Error("文件打开失败", err)
		msg := fmt.Sprintf("文件打开失败，错误：%v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(msg))
		return
	}
	fileStat, _ := file.Stat()

	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Disposition", "attachment; filename=\""+url.QueryEscape(fileName)+"\"")
	w.Header().Set("Content-Length", strconv.FormatInt(fileStat.Size(), 10))

	file.Seek(0, 0)
	io.Copy(w, file)

	return
}

func (f *fileService) ImportIncome(w http.ResponseWriter, r *http.Request) {
	userId := GetUserId(r)

	accIdList, err := dbUserAccount.GetIdListByUserId(r.Context(), userId)
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "未分配可操作的账户")
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		utils.RspBadError(w, "file required")
		return
	}
	defer func() { _ = file.Close() }()

	if header.Size > maxUploadSize {
		utils.RspBadError(w, "file size limit 10M")
		return
	}

	pathDir := filepath.Join(config.DefaultConfig.AssetDir, "import", time.Now().Format("20060102"))
	exist, _ := utils.IsPathExist(pathDir)
	if !exist {
		_ = os.MkdirAll(pathDir, 0777)
	}
	path := filepath.Join(pathDir, fmt.Sprintf("ep_%s_%s", utils.GenUUID(), header.Filename))
	err = utils.SaveFile(file, path)
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "save error")
		return
	}

	excel, err := excelize.OpenFile(path)
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "open excel error")
		return
	}
	defer func() { _ = excel.Close() }()

	rows, err := excel.GetRows("Sheet1")
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "get excel rows error")
		return
	}

	// 第一行为表头
	if len(rows) < 2 {
		log.Error("err:", err)
		utils.RspError(w, "数据不能为空")
		return
	}

	var (
		errMsg      string
		batchRecord []*erp.ModelIncome
		accountMap  = make(map[string]*erp.ModelAccount)
	)

	writeRecord := func(row []string) error {
		date := strings.TrimSpace(row[0])
		uuid := strings.TrimSpace(row[1])
		categoryName := strings.TrimSpace(row[2])
		from := strings.TrimSpace(row[3])
		mark := strings.TrimSpace(row[4])
		money := strings.TrimSpace(row[5])
		accountName := strings.TrimSpace(row[6])
		handleBy := strings.TrimSpace(row[7])

		if uuid == "" {
			uuid = utils.GenUUID()
		}

		incomeAt, err := utils.StrToTime(date)
		if err != nil {
			errMsg = fmt.Sprintf("【%v】该日期格式错误", date)
			log.Error("err:", err)
			return err
		}

		mon, err := strconv.ParseFloat(money, 64)
		if err != nil {
			errMsg = fmt.Sprintf("【%v】该金额格式错误", money)
			log.Error("err:", err)
			return err
		}
		monFen := uint32(mon * 100)

		account, ok := accountMap[accountName]
		if !ok {
			var acc erp.ModelAccount
			err = dbAccount.newScope().
				SetNotFoundErr(erp.ErrNotFoundAccount).
				Eq(dbName, accountName).First(&acc)
			if err != nil {
				if !errorx.IsNotFoundErr(err, erp.ErrNotFoundAccount) {
					return err
				}
				account = &erp.ModelAccount{}
			}
			account = &acc
			accountMap[accountName] = &acc
		}
		if account.Id == 0 {
			errMsg = fmt.Sprintf("【%s】该账户不存在", accountName)
			return errorx.New(erp.ErrNotFoundAccount)
		}

		if !isInSliceUint64(account.Id, accIdList) {
			errMsg = fmt.Sprintf("没有【%s】该账户的操作权限", accountName)
			return errorx.New(erp.ErrNotPermissionForAccount)
		}

		cat, ok := erp.IncomeCategoryMap[categoryName]
		if !ok {
			errMsg = fmt.Sprintf("没有【%s】该收入类目", categoryName)
			return errorx.New(erp.ErrIncomeCategoryInvalid)
		}

		batchRecord = append(batchRecord, &erp.ModelIncome{
			IncomeAt:    uint32(incomeAt.Unix()),
			Uuid:        uuid,
			Category:    cat,
			Mark:        mark,
			IncomeMoney: monFen,
			AccountId:   account.Id,
			HandleBy:    handleBy,
			UserId:      userId,
			From:        from,
		})

		return nil
	}

	isHeader := true
	for _, row := range rows {
		if isHeader {
			isHeader = false
			continue
		}
		if len(row) != 8 {
			utils.RspError(w, "请检查表格列数")
			return
		}

		err = writeRecord(row)
		if err != nil {
			log.Error("err:", err)
			msg := "写入失败"
			if errors.Is(err, errorx.New(erp.ErrExpenseCategoryInvalid)) {
				msg = erp.ErrCodeMap[erp.ErrExpenseCategoryInvalid]
			} else if errors.Is(err, errorx.New(erp.ErrNotPermissionForAccount)) {
				msg = erp.ErrCodeMap[erp.ErrNotPermissionForAccount]
			}
			if errMsg != "" {
				msg = errMsg
			}
			utils.RspError(w, msg)
			return
		}
	}

	if len(batchRecord) > 0 {
		err = dbIncome.newScope().CreateInBatches(batchRecord, 100)
		if err != nil {
			log.Error("err:", err)
			utils.RspError(w, "write record error")
			return
		}
	}

	utils.RspOK(w)
	return
}

func isInSliceUint64(i uint64, s []uint64) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}

func (f *fileService) ExportExpense(w http.ResponseWriter, r *http.Request) {
	var req erp.ListExpenseReq

	ctx := ParseCtx(r)

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil && err != io.EOF {
		log.Error("err:", err)
		utils.RspBadError(w, "请求参数错误")
		return
	}

	if !req.IsExport {
		utils.RspBadError(w, "请求参数错误")
		return
	}

	listRsp, err := NewErpService().ListExpense(ctx, &req)
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "导出失败")
		return
	}

	if listRsp.Paginate.Total > 200000 {
		log.Error("err:", err)
		utils.RspError(w, "超出导出上限20万")
		return
	}

	excel := excelize.NewFile()
	defer excel.Close()
	index, err := excel.NewSheet("支出列表")
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "导出失败")
		return
	}
	excel.SetActiveSheet(index)

	excelWriter, err := excel.NewStreamWriter("支出列表")
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "导出失败")
		return
	}
	// 第一行写表头
	err = excelWriter.SetRow("A1", []interface{}{"日期", "支出ID", "科目", "摘要", "支出", "账户", "凭证", "经手人"})
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "导出失败")
		return
	}

	if len(listRsp.List) > 0 {
		// 第二行开始写数据
		rowNum := 2
		for _, v := range listRsp.List {
			payAt := time.Unix(int64(v.PayAt), 0).Format("2006-01-02 15:04:05")
			cateName := erp.ExpenseCategoryMapName[v.Category]

			var accName string
			if acc, ok := listRsp.AccountMap[v.AccountId]; ok {
				accName = acc.Name
			}

			money := float64(v.PayMoney) / 100

			record := []interface{}{payAt, v.Uuid, cateName, v.Mark, money, accName, v.Ticket, v.HandleBy}

			cell, _ := excelize.CoordinatesToCellName(1, rowNum)
			if err = excelWriter.SetRow(cell, record); err != nil {
				log.Error("err:", err)
				utils.RspError(w, "导出失败")
				return
			}
			rowNum++
		}

	}
	if err = excelWriter.Flush(); err != nil {
		log.Error("err:", err)
		utils.RspError(w, "导出失败")
		return
	}

	fileName := "支出列表.xlsx"
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Disposition", "attachment; filename=\""+url.QueryEscape(fileName)+"\"")

	err = excel.Write(w)
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "导出失败")
		return
	}

	return
}

func (f *fileService) ExportIncome(w http.ResponseWriter, r *http.Request) {
	var req erp.ListIncomeReq

	ctx := ParseCtx(r)

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil && err != io.EOF {
		log.Error("err:", err)
		utils.RspBadError(w, "请求参数错误")
		return
	}

	if !req.IsExport {
		utils.RspBadError(w, "请求参数错误")
		return
	}

	listRsp, err := NewErpService().ListIncome(ctx, &req)
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "导出失败")
		return
	}

	if listRsp.Paginate.Total > 200000 {
		log.Error("err:", err)
		utils.RspError(w, "超出导出上限20万")
		return
	}

	excel := excelize.NewFile()
	defer excel.Close()
	index, err := excel.NewSheet("收入列表")
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "导出失败")
		return
	}
	excel.SetActiveSheet(index)

	excelWriter, err := excel.NewStreamWriter("收入列表")
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "导出失败")
		return
	}
	// 第一行写表头
	err = excelWriter.SetRow("A1", []interface{}{"日期", "收款ID", "科目", "收入来源", "摘要", "收入", "账户", "经手人"})
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "导出失败")
		return
	}

	if len(listRsp.List) > 0 {
		// 第二行开始写数据
		rowNum := 2
		for _, v := range listRsp.List {
			incomeAt := time.Unix(int64(v.IncomeAt), 0).Format("2006-01-02 15:04:05")
			cateName := erp.IncomeCategoryMapName[v.Category]

			var accName string
			if acc, ok := listRsp.AccountMap[v.AccountId]; ok {
				accName = acc.Name
			}

			money := float64(v.IncomeMoney) / 100

			record := []interface{}{incomeAt, v.Uuid, cateName, v.From, v.Mark, money, accName, v.HandleBy}

			cell, _ := excelize.CoordinatesToCellName(1, rowNum)
			if err = excelWriter.SetRow(cell, record); err != nil {
				log.Error("err:", err)
				utils.RspError(w, "导出失败")
				return
			}
			rowNum++
		}

	}
	if err = excelWriter.Flush(); err != nil {
		log.Error("err:", err)
		utils.RspError(w, "导出失败")
		return
	}

	fileName := "收入列表.xlsx"
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Disposition", "attachment; filename=\""+url.QueryEscape(fileName)+"\"")

	err = excel.Write(w)
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "导出失败")
		return
	}

	return
}

func (f *fileService) ImportPlan(w http.ResponseWriter, r *http.Request) {
	userId := GetUserId(r)
	file, header, err := r.FormFile("file")
	if err != nil {
		utils.RspBadError(w, "file required")
		return
	}
	defer func() { _ = file.Close() }()

	typ := r.FormValue("type")
	planTyp, err := strconv.Atoi(typ)
	if err != nil {
		utils.RspBadError(w, "type invalid")
		return
	}

	if header.Size > maxUploadSize {
		utils.RspBadError(w, "file size limit 10M")
		return
	}

	pathDir := filepath.Join(config.DefaultConfig.AssetDir, "import", time.Now().Format("20060102"))
	exist, _ := utils.IsPathExist(pathDir)
	if !exist {
		_ = os.MkdirAll(pathDir, 0777)
	}
	path := filepath.Join(pathDir, fmt.Sprintf("plan_%s_%s_%s", typ, utils.GenUUID(), header.Filename))
	err = utils.SaveFile(file, path)
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "save error")
		return
	}

	excel, err := excelize.OpenFile(path)
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "open excel error")
		return
	}
	defer func() { _ = excel.Close() }()

	rows, err := excel.GetRows("Sheet1")
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "get excel rows error")
		return
	}

	// 第一行为表头
	if len(rows) < 2 {
		log.Error("err:", err)
		utils.RspError(w, "数据不能为空")
		return
	}

	var (
		errMsg      string
		batchRecord []*erp.ModelPlan
	)

	writeRecord := func(row []string) error {
		date := strings.TrimSpace(row[0])
		uuid := strings.TrimSpace(row[1])
		customer := strings.TrimSpace(row[2])
		mark := strings.TrimSpace(row[3])
		totalMoney := strings.TrimSpace(row[4])
		tradeMoney := strings.TrimSpace(row[5])
		balanceMoney := strings.TrimSpace(row[6])
		comment := strings.TrimSpace(row[7])

		if uuid == "" {
			uuid = utils.GenUUID()
		}

		planAt, err := utils.StrToTime(date)
		if err != nil {
			errMsg = fmt.Sprintf("【%v】该日期格式错误", date)
			log.Error("err:", err)
			return err
		}

		totalMon, err := strconv.ParseFloat(totalMoney, 64)
		if err != nil {
			errMsg = fmt.Sprintf("【%v】该金额格式错误", totalMoney)
			log.Error("err:", err)
			return err
		}
		totalMonFen := uint32(totalMon * 100)

		tradeMon, err := strconv.ParseFloat(tradeMoney, 64)
		if err != nil {
			errMsg = fmt.Sprintf("【%v】该金额格式错误", tradeMoney)
			log.Error("err:", err)
			return err
		}
		tradeMonFen := uint32(tradeMon * 100)

		balanceMon, err := strconv.ParseFloat(balanceMoney, 64)
		if err != nil {
			errMsg = fmt.Sprintf("【%v】该金额格式错误", balanceMoney)
			log.Error("err:", err)
			return err
		}
		balanceMonFen := uint32(balanceMon * 100)

		if tradeMonFen+balanceMonFen != totalMonFen {
			errMsg = "金额计算错误"
			return errorx.New(erp.ErrMoneyBalance)
		}

		batchRecord = append(batchRecord, &erp.ModelPlan{
			Type:         uint32(planTyp),
			PlanAt:       uint32(planAt.Unix()),
			Uuid:         uuid,
			Customer:     customer,
			Mark:         mark,
			TotalMoney:   totalMonFen,
			TradeMoney:   tradeMonFen,
			BalanceMoney: balanceMonFen,
			Comment:      comment,
			UserId:       userId,
			Status:       uint32(erp.ModelPlan_StatusWaitConfirm),
		})

		return nil
	}

	isHeader := true
	for _, row := range rows {
		if isHeader {
			isHeader = false
			continue
		}
		if len(row) != 8 {
			utils.RspError(w, "请检查表格列数")
			return
		}

		err = writeRecord(row)
		if err != nil {
			log.Error("err:", err)
			msg := "写入失败"
			if errors.Is(err, errorx.New(erp.ErrExpenseCategoryInvalid)) {
				msg = erp.ErrCodeMap[erp.ErrExpenseCategoryInvalid]
			} else if errors.Is(err, errorx.New(erp.ErrNotPermissionForAccount)) {
				msg = erp.ErrCodeMap[erp.ErrNotPermissionForAccount]
			}
			if errMsg != "" {
				msg = errMsg
			}
			utils.RspError(w, msg)
			return
		}
	}

	if len(batchRecord) > 0 {
		err = dbPlan.newScope().CreateInBatches(batchRecord, 100)
		if err != nil {
			log.Error("err:", err)
			utils.RspError(w, "write record error")
			return
		}
	}

	utils.RspOK(w)
	return
}

func (f *fileService) ExportPlan(w http.ResponseWriter, r *http.Request) {
	var req erp.ListPlanReq

	ctx := ParseCtx(r)

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil && err != io.EOF {
		log.Error("err:", err)
		utils.RspBadError(w, "请求参数错误")
		return
	}

	if !req.IsExport {
		utils.RspBadError(w, "请求参数错误")
		return
	}

	listRsp, err := NewErpService().ListPlan(ctx, &req)
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "导出失败")
		return
	}

	if listRsp.Paginate.Total > 200000 {
		log.Error("err:", err)
		utils.RspError(w, "超出导出上限20万")
		return
	}

	excel := excelize.NewFile()
	defer excel.Close()
	index, err := excel.NewSheet("应付列表")
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "导出失败")
		return
	}
	excel.SetActiveSheet(index)

	excelWriter, err := excel.NewStreamWriter("应付列表")
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "导出失败")
		return
	}
	// 第一行写表头
	err = excelWriter.SetRow("A1", []interface{}{"日期", "付款ID", "客户名称", "摘要", "金额", "已付金额", "剩余应付", "备注", "状态"})
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "导出失败")
		return
	}

	if len(listRsp.List) > 0 {
		// 第二行开始写数据
		rowNum := 2
		for _, v := range listRsp.List {
			planAt := time.Unix(int64(v.PlanAt), 0).Format("2006-01-02 15:04:05")

			total := float64(v.TotalMoney) / 100
			trade := float64(v.TradeMoney) / 100
			balance := float64(v.BalanceMoney) / 100

			status := "待确认"
			if v.Status == 2 {
				status = "已确认"
			}

			record := []interface{}{planAt, v.Uuid, v.Customer, v.Mark, total, trade, balance, v.Comment, status}

			cell, _ := excelize.CoordinatesToCellName(1, rowNum)
			if err = excelWriter.SetRow(cell, record); err != nil {
				log.Error("err:", err)
				utils.RspError(w, "导出失败")
				return
			}
			rowNum++
		}

	}
	if err = excelWriter.Flush(); err != nil {
		log.Error("err:", err)
		utils.RspError(w, "导出失败")
		return
	}

	fileName := "应付列表.xlsx"
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Disposition", "attachment; filename=\""+url.QueryEscape(fileName)+"\"")

	err = excel.Write(w)
	if err != nil {
		log.Error("err:", err)
		utils.RspError(w, "导出失败")
		return
	}

	return
}
