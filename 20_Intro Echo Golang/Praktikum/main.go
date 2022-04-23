package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookWithCode struct {
	Code string `json:"code"`
	Book
}

func main() {
	books := map[string]Book{
		"001": {Title: "Tapak Jejak", Author: "Fiersa Besari"},
		"002": {Title: "The Hobbit", Author: "J.R.R. Tolkien"},
		"003": {Title: "The Da Vinci Code", Author: "Dan Brown"},
	}

	e := echo.New()

	e.GET("/books/:code", func(c echo.Context) error {
		bookCode := c.Param("code")
		book, exist := books[bookCode]
		if !exist {
			return c.String(http.StatusNotFound, "book not found")
		}
		return c.JSON(http.StatusOK, book)
	})

	e.GET("/books/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, books)
	})

	e.POST("books/", func(c echo.Context) error {
		bookWithCode := BookWithCode{}
		if err := c.Bind(&bookWithCode); err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		newBookCode := bookWithCode.Code
		book := Book{Title: bookWithCode.Title, Author: bookWithCode.Author}
		books[newBookCode] = book
		fmt.Println(books)
		return c.JSON(http.StatusOK, book)
	})

	e.PUT("books/:code", func(c echo.Context) error {
		bookCode := c.Param("code")
		book := Book{}
		if err := c.Bind(&book); err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		books[bookCode] = book
		fmt.Println(books)
		return c.JSON(http.StatusOK, book)
	})

	e.DELETE("books/:code", func(c echo.Context) error {
		bookCode := c.Param("code")
		deletedBook := books[bookCode]
		delete(books, bookCode)
		return c.JSON(http.StatusOK, deletedBook)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
