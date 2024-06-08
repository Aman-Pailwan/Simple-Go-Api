package routers

import (
	"github.com/Aman-Pailwan/controllers"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.GET("/books", controllers.GetBooks)
	r.GET("/books/:id", controllers.GetBookById)
	r.POST("/books", controllers.PostBooks)
	r.Run()
}
