package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() *gorm.DB {
	dsn := GetEnv().Db.DSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
