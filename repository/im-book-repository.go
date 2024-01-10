package repository

import (
	"github.com/lkcsi/goapi/custerror"
	"github.com/lkcsi/goapi/entity"
)

type imBookRepository struct {
	books []entity.Book
}

func NewImBookRepository() BookRepository {
	books := []entity.Book{
		{Id: "1", Title: "Title_1", Author: "Author_1", Quantity: 5},
		{Id: "2", Title: "Title_2", Author: "Author_2", Quantity: 0},
		{Id: "3", Title: "Title_3", Author: "Author_3", Quantity: 6},
		{Id: "4", Title: "Title_4", Author: "Author_4", Quantity: 5},
	}
	return &imBookRepository{books: books}
}

func (bs *imBookRepository) Save(book *entity.Book) error {
	bs.books = append(bs.books, *book)
	return nil
}

func (bs *imBookRepository) FindAll() ([]entity.Book, error) {
	return bs.books, nil
}

func (bs *imBookRepository) FindById(id string) (*entity.Book, error) {
	return bs.findBookById(id)
}

func (bs *imBookRepository) DeleteById(id string) error {
	index, err := bs.findBookIndex(id)
	if err != nil {
		return err
	}
	bs.books = append(bs.books[:index], bs.books[index+1:]...)
	return nil
}

func (bs *imBookRepository) findBookIndex(id string) (int, error) {
	for i, book := range bs.books {
		if book.Id == id {
			return i, nil
		}
	}
	return 0, custerror.NewNotFoundError(id)

}

func (bs *imBookRepository) findBookById(id string) (*entity.Book, error) {
	for i, book := range bs.books {
		if book.Id == id {
			return &bs.books[i], nil
		}
	}
	return nil, custerror.NewNotFoundError(id)
}
