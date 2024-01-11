package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lkcsi/goapi/custerror"
	"github.com/lkcsi/goapi/entity"
	"github.com/lkcsi/goapi/service"
)

type BookController interface {
	FindAll(context *gin.Context)
	FindById(context *gin.Context)
	DeleteById(context *gin.Context)
	Checkout(context *gin.Context)
	Save(context *gin.Context)
	DeleteAll(context *gin.Context)
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

func (c *bookController) DeleteAll(context *gin.Context) {
	if err := c.bookService.DeleteAll(); err != nil {
		setError(context, err)
		return
	}
	context.IndentedJSON(http.StatusNoContent, "")
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
	if err := c.bookService.DeleteById(id); err != nil {
		setError(context, err)
		return
	}
	context.IndentedJSON(http.StatusNoContent, nil)
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
	switch err.(type) {
	case validator.ValidationErrors:
		s := err.(validator.ValidationErrors)
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": getErrorMsg(s[0])})
	case custerror.CustError:
		s := err.(custerror.CustError)
		context.IndentedJSON(s.Code(), gin.H{"error": s.Error()})
	default:
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.ActualTag() {
	case "required":
		return fe.Field() + " field is mandatory"
	case "gte":
		return fe.Field() + " must be greater than or equals " + fe.Param()
	}
	return "unkown error"
}
