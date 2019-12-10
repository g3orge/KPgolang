package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Router() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/books", getBooks)
	e.GET("/books/:id", getBook)
	e.POST("/books", createBook)
	e.POST("/books/:id", deleteBook)
	e.POST("/register", register)
	e.POST("/login", login)

	return e
}
