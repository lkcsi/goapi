package custerror

import "fmt"

type CustError interface {
	Code() int
	Error() string
}

type custError struct {
	code    int
	message string
}

func (c *custError) Code() int {
	return c.code
}

func (err *custError) Error() string {
	return err.message
}

func NewOutOfStockError(id string) CustError {
	return &custError{400, fmt.Sprintf("book with id: %s is out of order", id)}
}

func NotFoundError(id string) CustError {
	return &custError{404, fmt.Sprintf("book with id: %s is not found", id)}
}
