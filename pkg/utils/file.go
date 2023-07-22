package utils

import (
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GetCurrentDir() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return strings.Replace(dir, "\\", "/", -1)
}

func EncodeMD5(value []byte) string {
	m := md5.New()
	m.Write(value)
	return hex.EncodeToString(m.Sum(nil))
}

func IsPathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetFileName(name string) string {
	ext := path.Ext(name)
	return strings.TrimRight(name, ext)
}

func SaveFile(src io.Reader, dst string) error {
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func ReadFromCsv(filename string) ([][]string, error) {
	f, err := os.Open(filename) // 打开文件
	if err != nil {
		return nil, err
	}
	var recordAll [][]string

	r := csv.NewReader(f)
	recordAll, err = r.ReadAll()
	if err != nil {
		return nil, err
	}

	return recordAll, nil
}
