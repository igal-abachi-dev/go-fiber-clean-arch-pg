package repositories

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"go-fiber-clean-arch-pg/entities"
)

type BookQueries struct {
}

func NewBookRepository() BookRepository {
	var repo = &BookQueries{}
	return repo
}

type BookRepository interface {
	GetBookById(id int32) (entities.Book, error)
	GetBooksByCategory(category int32) ([]entities.Book, error)
}

const getBookById = `-- name: GetBookById :one
SELECT "Name", "Author", "Id", "Category", "commentary" FROM public."Book"
WHERE "Id" = $1 LIMIT 1
`

func (q *BookQueries) GetBookById(id int32) (entities.Book, error) {
	ctx := context.Background()
	db, _ := pgxpool.Connect(ctx, PgConnection)

	var item entities.Book
	err := pgxscan.Get(ctx, db, &item, getBookById, id)

	return item, err
}

const getBooksByCategory = `-- name: GetBooksByCategory :many
SELECT "Name", "Author", "Id", "Category", "commentary" FROM public."Book"
WHERE "Category" = $1
ORDER BY "Id" ASC
`

func (q *BookQueries) GetBooksByCategory(category int32) ([]entities.Book, error) {

	ctx := context.Background()
	db, _ := pgxpool.Connect(ctx, PgConnection)

	var items []entities.Book
	err := pgxscan.Select(ctx, db, &items, getBooksByCategory, category)

	if err != nil {
		return nil, err
	}
	return items, nil
}
