package app

import (
	"github.com/JIAWea/erpServer/config"
	"github.com/JIAWea/erpServer/pkg/auth"
	log "github.com/ml444/glog"
	"os"
)

var trans *auth.RsaNormal

func Init() {
	var err error
	trans, err = auth.NewRsa([]byte(config.DefaultConfig.Trans.PubKey), []byte(config.DefaultConfig.Trans.PriKey))
	if err != nil {
		log.Error("err:", err)
		os.Exit(-1)
	}
}
