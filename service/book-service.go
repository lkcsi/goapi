package service

import (
	"errors"

	"github.com/lkcsi/goapi/entity"
)

type BookService interface {
	Save(entity.Book) (*entity.Book, error)
	FindAll() ([]entity.Book, error)
	FindById(string) (*entity.Book, error)
	DeleteById(string) (*entity.Book, error)
	Checkout(string) (*entity.Book, error)
}

type bookService struct {
	books []entity.Book
}

func NewInMemory() BookService {
	books := []entity.Book{
		{Id: "1", Title: "Title_1", Author: "Author_1", Quantity: 1},
		{Id: "2", Title: "Title_2", Author: "Author_2", Quantity: 3},
		{Id: "3", Title: "Title_3", Author: "Author_3", Quantity: 0},
		{Id: "4", Title: "Title_4", Author: "Author_4", Quantity: 4},
	}
	bs := bookService{books: books}
	return &bs
}

func (bs *bookService) Save(book entity.Book) (*entity.Book, error) {
	book.Id = "5"
	bs.books = append(bs.books, book)
	return &book, nil
}

func (bs *bookService) FindAll() ([]entity.Book, error) {
	return bs.books, nil
}

func (bs *bookService) FindById(id string) (*entity.Book, error) {
	return bs.findBookById(id)
}

func (bs *bookService) DeleteById(id string) (*entity.Book, error) {
	index, err := bs.findBookIndex(id)
	if err != nil {
		return nil, err
	}
	book := bs.books[index]
	bs.books = append(bs.books[:index], bs.books[index+1:]...)
	return &book, nil
}

func (bs *bookService) Checkout(id string) (*entity.Book, error) {
	book, err := bs.findBookById(id)
	if err != nil {
		return nil, err
	}
	book.Quantity -= 1
	return book, nil
}

func (bs *bookService) findBookIndex(id string) (int, error) {
	for i, book := range bs.books {
		if book.Id == id {
			return i, nil
		}
	}
	return 0, errors.New("book not found")
}

func (bs *bookService) findBookById(id string) (*entity.Book, error) {
	for i, book := range bs.books {
		if book.Id == id {
			return &bs.books[i], nil
		}
	}
	return nil, errors.New("book not found")
}
