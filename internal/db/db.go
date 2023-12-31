package db

import (
	"time"

	"gorm.io/driver/mysql"
	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	log "github.com/ml444/glog"
)

var db *gorm.DB
var migrateModelList []interface{}

func Db() *gorm.DB {
	return db
}
func InitDB(dsn string, debug bool) error {
	var err error
	//dsn := "user:pass@tcp(127.0.0.1:3306)/erp?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err = gorm.Open(sqlite.Open("erp.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if debug {
        db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
    } else {
        db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    }
	if err != nil {
		log.Error(err)
		return err
	}
	err = db.AutoMigrate(migrateModelList...)
	if err != nil {
		log.Error(err)
		return err
	}
	sqlDb, err := db.DB()
	if err != nil {
		log.Errorf("err: %v", err)
		return err
	}
	sqlDb.SetConnMaxLifetime(time.Hour * 16)
	sqlDb.SetConnMaxIdleTime(time.Hour * 2)
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	return nil
}

func CloseDB() error {
	if sqlDb, err := db.DB(); err == nil {
		return sqlDb.Close()
	} else {
		println(err.Error())
		return err
	}
}

func RegisterModel(mList ...interface{}) {
	migrateModelList = append(migrateModelList, mList...)
}
