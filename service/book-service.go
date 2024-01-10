package service

import (
	"github.com/google/uuid"
	"github.com/lkcsi/goapi/custerror"
	"github.com/lkcsi/goapi/entity"
	"github.com/lkcsi/goapi/repository"
)

type BookService interface {
	Save(entity.Book) (*entity.Book, error)
	FindAll() ([]entity.Book, error)
	FindById(string) (*entity.Book, error)
	DeleteById(string) error
	Checkout(string) (*entity.Book, error)
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewInMemoryBookService() BookService {
	repo := repository.NewImBookRepository()
	bs := bookService{repo}
	return &bs
}

func NewSqlBookService() BookService {
	repo := repository.NewSqlBookRepository()
	return &bookService{repo}
}

func (bs *bookService) Save(book entity.Book) (*entity.Book, error) {
	book.Id = uuid.NewString()
	if err := bs.bookRepository.Save(&book); err != nil {
		return nil, err
	}
	return &book, nil
}

func (bs *bookService) FindAll() ([]entity.Book, error) {
	return bs.bookRepository.FindAll()
}

func (bs *bookService) FindById(id string) (*entity.Book, error) {
	book, err := bs.bookRepository.FindById(id)
	if err != nil {
		return nil, custerror.NotFoundError(id)
	}
	return book, nil
}

func (bs *bookService) DeleteById(id string) error {
	if err := bs.bookRepository.DeleteById(id); err != nil {
		return err
	}
	return nil
}

func (bs *bookService) Checkout(id string) (*entity.Book, error) {
	book, err := bs.bookRepository.FindById(id)
	if err != nil {
		return nil, custerror.NotFoundError(id)
	}
	if book.Quantity == 0 {
		return nil, custerror.NewOutOfStockError(id)
	}
	book.Quantity -= 1
	if err := bs.bookRepository.Save(book); err != nil {
		return nil, err
	}
	return book, nil
}
