package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lkcsi/goapi/entity"
)

type sqlBookRepository struct {
	connectionString string
}

func NewSqlBookRepository() BookRepository {
	return &sqlBookRepository{"root:asdfgh@tcp(localhost:3308)/book_db"}
}

func (repo *sqlBookRepository) FindAll() ([]entity.Book, error) {
	db, err := sql.Open("mysql", repo.connectionString)
	defer db.Close()
	if err != nil {
		return nil, err
	}

	res, err := db.Query("SELECT * FROM books")
	defer res.Close()
	if err != nil {
		return nil, err
	}

	result := make([]entity.Book, 0)
	for res.Next() {
		var book entity.Book
		err := res.Scan(&book.Id, &book.Title, &book.Author, &book.Quantity)
		if err != nil {
			return nil, err
		}
		result = append(result, book)
	}
	return result, nil
}

func (*sqlBookRepository) DeleteById(string) error {
	panic("unimplemented")
}

func (repo *sqlBookRepository) FindById(id string) (*entity.Book, error) {
	db, err := sql.Open("mysql", repo.connectionString)
	defer db.Close()
	if err != nil {
		return nil, err
	}

	var book entity.Book
	row := db.QueryRow("SELECT id, title, author, quantity FROM books WHERE id=?", id)
	err = row.Scan(&book.Id, &book.Title, &book.Author, &book.Quantity)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (*sqlBookRepository) Save(*entity.Book) error {
	panic("unimplemented")
}
