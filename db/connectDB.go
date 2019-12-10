package db

import "github.com/jinzhu/gorm"

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func Db(inDb *gorm.DB) {
	db = inDb
}
