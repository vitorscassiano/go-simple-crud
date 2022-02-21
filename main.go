package main

import (
	"github.com/gin-gonic/gin"
	"shortener.com/v1/controllers"
	"shortener.com/v1/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
