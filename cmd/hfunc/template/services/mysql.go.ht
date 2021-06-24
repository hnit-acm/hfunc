package services

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	_MysqlDB *gorm.DB
)

func Mysql() *gorm.DB {
	if _MysqlDB == nil {
		orm, err := gorm.Open(mysql.New(mysql.Config{
			DSN: "root:12345@tcp(127.0.0.1:3306)/auth_service?charset=utf8&parseTime=True&loc=Local",
		}))
		if err != nil {
			log.Panicln(err)
		}
		_MysqlDB = orm
	}
	return _MysqlDB
}
