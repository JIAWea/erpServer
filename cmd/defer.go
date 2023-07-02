package main

import (
	log "github.com/ml444/glog"
	"github.com/JIAWea/erpServer/internal/db"
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
