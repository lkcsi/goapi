package main

import (
	"errors"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity string `json:"quantity"`
}

var books = []book{
	{Id: "1", Title: "Title_1", Author: "Author_1", Quantity: "Quantity_1"},
	{Id: "2", Title: "Title_2", Author: "Author_2", Quantity: "Quantity_2"},
	{Id: "3", Title: "Title_3", Author: "Author_3", Quantity: "Quantity_3"},
	{Id: "4", Title: "Title_4", Author: "Author_4", Quantity: "Quantity_4"},
}

func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, books)
}

func createBook(context *gin.Context) {
	var newBook book
	if err := context.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	context.IndentedJSON(http.StatusCreated, newBook)
}

func getBookById(context *gin.Context) {
	id := context.Param("id")
	book, err := findBookById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, err)
	}
	context.IndentedJSON(http.StatusOK, book)
}

func findBookById(id string) (*book, error) {
	for i, book := range books {
		if book.Id == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func main() {
	log.Println("Application started")

	server := gin.Default()
	server.GET("/books", getBooks)
	server.GET("/books/:id", getBookById)
	server.POST("/books", createBook)
	server.Run("localhost:8080")
}
