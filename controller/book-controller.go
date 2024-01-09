package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lkcsi/goapi/entity"
	"github.com/lkcsi/goapi/service"
)

type BookController interface {
	FindAll(context *gin.Context)
	FindById(context *gin.Context)
	DeleteById(context *gin.Context)
	Checkout(context *gin.Context)
	Save(context *gin.Context)
}

type bookController struct {
	bookService service.BookService
}

func New(s *service.BookService) *bookController {
	return &bookController{bookService: *s}
}

func (c *bookController) FindAll(context *gin.Context) {
	books, err := c.bookService.FindAll()
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, books)
}

func (c *bookController) Save(context *gin.Context) {
	var requestedBook entity.Book
	if err := context.BindJSON(&requestedBook); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	newBook, err := c.bookService.Save(requestedBook)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusCreated, newBook)
}

func (c *bookController) DeleteBookById(context *gin.Context) {
	id := context.Param("id")
	book, err := c.bookService.DeleteById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusNoContent, book)
}
func (c *bookController) CheckoutBook(context *gin.Context) {
	id := context.Param("id")
	book, err := c.bookService.Checkout(id)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusAccepted, book)
}

func (c *bookController) FindById(context *gin.Context) {
	id := context.Param("id")
	book, err := c.bookService.FindById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, book)
}
