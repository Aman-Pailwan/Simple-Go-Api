package controllers

import (
	"net/http"

	"github.com/Aman-Pailwan/models"
	"github.com/gin-gonic/gin"
)

func GetBookById(c *gin.Context) {
	id := c.Param("id")

	for _, v := range models.Books {
		if v.ID == id {
			c.IndentedJSON(http.StatusFound, v)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "Book with id " + id + " not found",
	})
}

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Books)
}

func PostBooks(c *gin.Context) {

	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, err)
	}
	models.Books = append(models.Books, newBook)

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Successfully Added Book in the list",
	})
}
