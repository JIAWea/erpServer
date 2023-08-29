package utils

import (
	"fmt"
	"net/http"
)

func RspOK(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{}`))
}

func RspError(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
}

func RspBadError(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
}

func Rsp401Error(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
}
