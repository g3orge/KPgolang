package main

import (
	"log"

	"../Kproject/api"
	"../Kproject/db"
	"../Kproject/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	DB, err := gorm.Open("mysql", "root:root@/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Print(err)
	}

	DB.LogMode(true)

	DB.Table("user").AutoMigrate(&model.User{})
	DB.Table("books").AutoMigrate(&model.Books{})
	DB.Table("token").AutoMigrate(&model.Token{})

	db.Db(DB)
}

func main() {
	r := api.Router()

	r.Logger.Fatal(r.Start(":8080"))
}
