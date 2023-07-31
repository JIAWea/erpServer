package app

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/JIAWea/erpServer/api/erp"
	"github.com/ml444/gkit/errorx"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"

	"github.com/JIAWea/erpServer/config"
	"github.com/JIAWea/erpServer/pkg/utils"
	log "github.com/ml444/glog"
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

	data, err := io.ReadAll(file)
	if err != nil {
		log.Error("读取文件失败", err)
		msg := fmt.Sprintf("读取文件失败，错误：%v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(msg))
		return
	}

	if !strings.HasSuffix(fileName, ".xlsx") {
		fileName += ".xlsx"
	}
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Disposition", "attachment; filename=\""+url.QueryEscape(fileName)+"\"")
	_, err = w.Write(data)
	if err != nil {
		log.Error("下载文件失败", err)
		msg := fmt.Sprintf("文件打开失败，错误：%v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(msg))
		return
	}

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
