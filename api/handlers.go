package api

import (
	"net/http"

	"../db"
	"../model"
	"github.com/labstack/echo"
)

func getBooks(c echo.Context) error {
	var books []model.Books

	if err := db.GetDB().Table("books").Find(&books).Error; err != nil {
		return c.String(http.StatusBadRequest, "cannot connect to database")
	}

	return c.JSON(http.StatusOK, books)
}

func getBook(c echo.Context) error {
	id := c.Param("id")

	var book model.Books

	if err := db.GetDB().Table("books").First(&book, id).Error; err != nil {
		return c.String(http.StatusBadRequest, "cannot connect to database")
	}

	return c.JSON(http.StatusOK, book)
}

func createBook(c echo.Context) error {
	var book model.Books

	err := c.Bind(&book)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	var check model.Books
	if err := db.GetDB().Table("books").Where("bookname = ? AND name = ? AND surname = ?", book.Bookname, book.Authorname, book.Authorsurname).Find(&check).Error; err == nil {
		return c.String(http.StatusConflict, "This book has already been added")
	}

	if err := db.GetDB().Table("books").Save(&book).Error; err != nil {
		return c.String(http.StatusBadRequest, "cannot connect to database")
	}

	return c.String(http.StatusOK, "book added")
}

func deleteBook(c echo.Context) error {
	id := c.Param("id")

	if err := db.GetDB().Table("books").Where("ID = ?", id).Delete(&model.Books{}).Error; err != nil {
		return c.String(http.StatusBadRequest, "cannot connect to database")
	}

	return c.String(http.StatusOK, "book deleted")
}
