package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string `json :"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

var books = []book{
	{ID: "1", Name: "Atomic Habits", Author: "James Clear"},
}

func main() {
	r := gin.Default()
	r.GET("/books", getBooks)
	r.GET("/books/:id", getBookById)
	r.POST("/books", postBooks)
	r.Run()
}

func getBookById(c *gin.Context) {
	id := c.Param("id")

	for _, v := range books {
		if v.ID == id {
			c.IndentedJSON(http.StatusFound, v)
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "Book with id " + id + " not found",
	})
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func postBooks(c *gin.Context) {

	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, err)
	}
	books = append(books, newBook)

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Successfully Added Book in the list",
	})
}
