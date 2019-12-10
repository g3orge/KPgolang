package api

import (
	"net/http"
	"time"

	"../db"
	"../model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func register(c echo.Context) error {
	var dbUser model.User

	err := c.Bind(&dbUser)
	if err != nil {
		return c.String(http.StatusBadRequest, "cannot bind")
	}

	var test model.User
	if err := db.GetDB().Table("user").Where("email = ?", dbUser.Email).Find(&test).Error; err == nil {
		return c.String(http.StatusConflict, "User with this email already exist")
	}

	if err := db.GetDB().Table("user").Save(&dbUser).Error; err != nil {
		return c.String(http.StatusConflict, "cannot save user")
	}

	return c.String(http.StatusOK, "user added")
}

func login(c echo.Context) error {
	var test model.User
	var jwtKey = []byte("my_secret_key")

	err := c.Bind(&test)
	if err != nil {
		return c.String(http.StatusBadRequest, "cannot bind")
	}

	var test2 model.User
	if err := db.GetDB().Table("user").Where("email = ? and password = ?", test.Email, test.Password).Find(&test2).Error; err != nil {
		return c.String(http.StatusLocked, "You are not registered")
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &model.Claims{
		Email: test.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.String(http.StatusInternalServerError, "error in jwt")
	}
	var tokenjwt = model.Token{
		UserId: test.ID,
		Token:  tokenString,
	}
	db.GetDB().Table("token").Save(&tokenjwt)

	return c.String(http.StatusOK, "you are registered")
}
