package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lkcsi/goapi/controller"
	"github.com/lkcsi/goapi/service"
)

func authEnabled() bool {
	v := os.Getenv("AUTH_ENABLED")
	r, ok := strconv.ParseBool(v)
	if ok == nil {
		return r
	}
	return false
}

var authService service.AuthService
var bookService service.BookService

func getAuthService() service.AuthService {
	if authEnabled() {
		return service.NewJwtAuthService()
	}
	return service.NewFakeAuthService()
}

func getBookService() service.BookService {
	if os.Getenv("REPOSITORY") == "SQL" {
		return service.NewSqlBookService()
	}
	return service.NewInMemoryBookService()
}

func main() {
	godotenv.Load()

	server := gin.Default()

	bookService = getBookService()
	bookController := controller.New(&bookService)
	authService = getAuthService()

	books := server.Group("/books")
	books.Use(authService.Auth)
	books.GET("", bookController.FindAll)
	books.GET("/:id", bookController.FindById)
	books.DELETE("/:id", bookController.DeleteBookById)
	books.DELETE("", bookController.DeleteAll)
	books.POST("", bookController.Save)
	books.PATCH("/:id/checkout", bookController.CheckoutBook)

	server.Run("localhost:8080")
}
