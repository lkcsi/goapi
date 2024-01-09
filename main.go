package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lkcsi/goapi/controller"
	"github.com/lkcsi/goapi/service"
)

func main() {
	server := gin.Default()

	bookService := service.NewInMemory()
	bookController := controller.New(&bookService)

	server.GET("/books", bookController.FindAll)
	server.GET("/books/:id", bookController.FindById)
	server.DELETE("/books/:id", bookController.DeleteBookById)
	server.POST("/books", bookController.Save)
	server.PATCH("/books/:id/checkout", bookController.CheckoutBook)

	server.Run("localhost:8080")
}
