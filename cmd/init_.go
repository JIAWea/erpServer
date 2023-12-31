package main

import (
	"github.com/JIAWea/erpServer/api/erp"
	"github.com/JIAWea/erpServer/config"
	"github.com/JIAWea/erpServer/internal/db"
	"github.com/ml444/gkit/auth"
	gkitlog "github.com/ml444/gkit/log"
	log "github.com/ml444/glog"
	logconf "github.com/ml444/glog/config"
	level "github.com/ml444/glog/level"
	"github.com/redis/go-redis/v9"
)

func Init(cfg *config.Config) error {
	var err error

	// setting logger
	err = InitLogger(cfg.Debug)
	if err != nil {
		log.Errorf("err: %v", err)
		return err
	}

	// init db
	err = db.InitDB(cfg.DbDSN, cfg.Debug)
	if err != nil {
		log.Errorf("err: %v", err)
		return err
	}

	// TODO redis
	auth.InitRedis(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Username: cfg.Redis.User,
		Password: cfg.Redis.Passwd,
	})

	return nil
}

func InitLogger(debug bool) error {
	err := log.InitLog(
		logconf.SetLoggerName(erp.ClientName),
		logconf.SetLevel2Logger(level.InfoLevel),
	)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	if debug {
		err = log.InitLog(logconf.SetLevel2Logger(level.DebugLevel))
		if err != nil {
			log.Errorf("err: %v", err)
			return err
		}
	}
	gkitlog.SetLogger(log.GetLogger())
	return err
}
