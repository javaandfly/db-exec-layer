package db

import (
	"db-exec-layer/config"
	"fmt"

	"github.com/javaandfly/gwebz/db"
	"gorm.io/gorm"
)

var _sqlPool *gorm.DB

func GetDB() *gorm.DB {

	return _sqlPool
}

func SetDB(sqlPool *gorm.DB) (err error) {
	if sqlPool == nil {
		err = fmt.Errorf("sql pool not found")
		return
	}
	_sqlPool = sqlPool
	return
}

func NewDBEngine(config *config.DatabaseSource) (engine *gorm.DB, err error) {

	return db.NewDB(
		config.DriverName,
		config.Host,
		config.Port,
		config.Database,
		config.Username,
		config.Password,
		config.Charset,
		config.Debug)

}
