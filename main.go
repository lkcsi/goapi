package main

import (
	"errors"
	"log"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{Id: 1, Title: "Title_1", Author: "Author_1", Quantity: 1},
	{Id: 2, Title: "Title_2", Author: "Author_2", Quantity: 3},
	{Id: 3, Title: "Title_3", Author: "Author_3", Quantity: 0},
	{Id: 4, Title: "Title_4", Author: "Author_4", Quantity: 4},
}

func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, books)
}

func createBook(context *gin.Context) {
	var newBook book
	if err := context.BindJSON(&newBook); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	books = append(books, newBook)
	context.IndentedJSON(http.StatusCreated, newBook)
}

func deleteById(context *gin.Context) {
	str := context.Param("id")
	id, err := strconv.Atoi(str)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	index, err := findBookIndex(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	books = append(books[:index], books[index+1:]...)
	context.IndentedJSON(http.StatusNoContent, nil)
}
func checkoutBook(context *gin.Context) {
	str := context.Param("id")
	id, err := strconv.Atoi(str)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	book, err := findBookById(id)
	if book.Quantity <= 0 {
		context.IndentedJSON(http.StatusBadGateway, gin.H{"message": "out of stock"})
		return
	}

	book.Quantity -= 1
	context.IndentedJSON(http.StatusOK, gin.H{"quantity": book.Quantity})
}

func getBookById(context *gin.Context) {
	str := context.Param("id")
	id, err := strconv.Atoi(str)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	book, err := findBookById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, book)
}

func findBookIndex(id int) (int, error) {
	for i, book := range books {
		if book.Id == id {
			return i, nil
		}
	}
	return 0, errors.New("book not found")
}

func findBookById(id int) (*book, error) {
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
	server.DELETE("/books/:id", deleteById)
	server.POST("/books", createBook)
	server.PATCH("/checkout/:id", checkoutBook)

	server.Run("localhost:8080")
}
