package main

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// go "errors"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Hello", Author: "Nguyen", Quantity: 1},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)

	router.POST("/books", func(c *gin.Context) {
		var newBook book

		if err := c.BindJSON(&newBook); err != nil {
			return
		}

		books = append(books, newBook)
		c.IndentedJSON(http.StatusCreated, newBook)
	})

	router.Run("localhost:3000")
}
