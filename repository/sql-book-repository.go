package repository

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lkcsi/goapi/entity"
)

type sqlBookRepository struct {
	connectionString string
	db               *sql.DB
}

func NewSqlBookRepository() BookRepository {
	pwd := os.Getenv("MYSQL_PASSWORD")
	port := os.Getenv("MYSQL_PORT")
	conn := fmt.Sprintf("root:%s@tcp(localhost:%s)/book_db", pwd, port)
	return &sqlBookRepository{conn, nil}
}

func (repo *sqlBookRepository) openConnection() error {
	if repo.db == nil {
		db, err := sql.Open("mysql", repo.connectionString)
		if err != nil {
			return err
		}
		repo.db = db
	}
	return nil
}

func (repo *sqlBookRepository) FindAll() ([]entity.Book, error) {
	if err := repo.openConnection(); err != nil {
		return nil, err
	}

	res, err := repo.db.Query("SELECT * FROM books")
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
	if err := repo.openConnection(); err != nil {
		return nil, err
	}

	var book entity.Book
	row := repo.db.QueryRow("SELECT id, title, author, quantity FROM books WHERE id=?", id)
	err := row.Scan(&book.Id, &book.Title, &book.Author, &book.Quantity)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (*sqlBookRepository) Save(*entity.Book) error {
	panic("unimplemented")
}
