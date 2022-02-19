package app

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jwtsmtp/helper"
)

func InitConnect() *gorm.DB{
	dsn := "root:@tcp(127.0.0.1:3306)/jwtsmtp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.ErrorIfNotNil(err)

	return db
}