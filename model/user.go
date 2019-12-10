package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type (
	User struct {
		gorm.Model
		Email    string `gorm:"column:email" json:"email"`
		Password string `gorm:"column:password" json:"password"`
	}

	Books struct {
		gorm.Model
		Bookname      string `gorm:"column:bookname" json:"bookname"`
		Authorname    string `gorm:"column:name" json:"name"`
		Authorsurname string `gorm:"column:surname" json:"surname"`
	}

	Claims struct {
		Email string `json:"email"`
		jwt.StandardClaims
	}

	Token struct {
		UserId uint   `gorm:"column:userid"`
		Token  string `gorm:"column:token"`
	}
)
