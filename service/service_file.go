package app

import (
	"bytes"
	"fmt"
	"github.com/JIAWea/erpServer/config"
	"github.com/JIAWea/erpServer/pkg/utils"
	log "github.com/ml444/glog"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
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

	savePath := filepath.Join(config.DefaultConfig.AssetDir, md5)
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

func (f *fileService) DownloadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fileName := r.FormValue("file_name")
	md5 := r.FormValue("md5")
	savePath := filepath.Join(config.DefaultConfig.AssetDir, utils.GetFileName(md5))

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
