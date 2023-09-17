package main

import (
	"github.com/JIAWea/erpServer/internal/db"
	log "github.com/ml444/glog"
)

func Defer() {
	CloseDbConns()
}

func CloseDbConns() {
	err := db.CloseDB()
	if err != nil {
		log.Error(err)
		return
	}
}
