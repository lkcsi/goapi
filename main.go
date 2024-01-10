package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lkcsi/goapi/controller"
	"github.com/lkcsi/goapi/middleware"
	"github.com/lkcsi/goapi/service"
)

func main() {
	server := gin.Default()

	bookService := service.NewInMemory()
	bookController := controller.New(&bookService)

	books := server.Group("/books")
	books.Use(middleware.Test())

	books.GET("", bookController.FindAll)
	books.GET("/:id", bookController.FindById)
	books.DELETE("/:id", bookController.DeleteBookById)
	books.POST("", bookController.Save)
	books.PATCH("/:id/checkout", bookController.CheckoutBook)

	server.Run("localhost:8080")
}
