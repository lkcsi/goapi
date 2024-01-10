package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/lkcsi/goapi/controller"
	"github.com/lkcsi/goapi/entity"
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

func authService() service.AuthService {
	if authEnabled() {
		return service.NewJwtAuthService()
	}
	return service.NewFakeAuthService()
}

func main() {
	godotenv.Load()

	test_db()

	server := gin.Default()

	bookService := service.NewInMemory()
	bookController := controller.New(&bookService)
	authService := authService()

	books := server.Group("/books")
	books.Use(authService.Auth)
	books.GET("", bookController.FindAll)
	books.GET("/:id", bookController.FindById)
	books.DELETE("/:id", bookController.DeleteBookById)
	books.POST("", bookController.Save)
	books.PATCH("/:id/checkout", bookController.CheckoutBook)

	server.Run("localhost:8080")
}

func test_db() {
	db, err := sql.Open("mysql", "root:asdfgh@tcp(localhost:3308)/book_db")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var version string
	err = db.QueryRow("SELECT VERSION()").Scan(&version)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(version)

	res, err := db.Query("SELECT * FROM books")
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var book entity.Book
		err := res.Scan(&book.Id, &book.Title, &book.Author, &book.Quantity)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(book)
	}

}
