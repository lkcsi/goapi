package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		setError(context, err)
		return
	}
	context.IndentedJSON(http.StatusOK, books)
}

func (c *bookController) Save(context *gin.Context) {
	context.Writer.Header().Set("content-type", "application/json")
	var requestedBook entity.Book
	if err := context.BindJSON(&requestedBook); err != nil {
		setError(context, err)
		return
	}

	newBook, err := c.bookService.Save(requestedBook)
	if err != nil {
		setError(context, err)
		return
	}

	context.IndentedJSON(http.StatusCreated, newBook)
}

func (c *bookController) DeleteBookById(context *gin.Context) {
	id := context.Param("id")
	book, err := c.bookService.DeleteById(id)
	if err != nil {
		setError(context, err)
		return
	}
	context.IndentedJSON(http.StatusNoContent, book)
}
func (c *bookController) CheckoutBook(context *gin.Context) {
	id := context.Param("id")
	book, err := c.bookService.Checkout(id)
	if err != nil {
		setError(context, err)
		return
	}
	context.IndentedJSON(http.StatusAccepted, book)
}

func (c *bookController) FindById(context *gin.Context) {
	id := context.Param("id")
	book, err := c.bookService.FindById(id)
	if err != nil {
		setError(context, err)
		return
	}
	context.IndentedJSON(http.StatusOK, book)
}

type ErrorMsg struct {
	Message string `json:"message"`
}

func setError(context *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": getErrorMsg(ve[0])})
	} else {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " field is mandatory"
	case "gt":
		return fe.Field() + " should be greater than " + fe.Param()
	}
	return "unkown error"
}
