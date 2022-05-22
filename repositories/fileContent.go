package repositories

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"go-fiber-clean-arch-pg/entities"
)

type FileContentQueries struct {
}

func NewFileContentRepository() FileContentRepository {
	var repo = &FileContentQueries{}
	return repo
}

type FileContentRepository interface {
	GetBookTextContent(bookID int32) (entities.GetBookTextContentRow, error)
	GetBookUrl(bookID int32) (entities.GetBookUrlRow, error)
}

const getBookTextContent = `-- name: GetBookTextContent :one
SELECT content FROM public."FileContent"
WHERE "book_id" = $1 LIMIT 1
`

func (q *FileContentQueries) GetBookTextContent(bookID int32) (entities.GetBookTextContentRow, error) {
	ctx := context.Background()
	db, _ := pgxpool.Connect(ctx, PgConnection)

	var item entities.GetBookTextContentRow
	err := pgxscan.Get(ctx, db, &item, getBookTextContent, bookID)

	return item, err
}

const getBookUrl = `-- name: GetBookUrl :one
SELECT url, password FROM public."FileContent"
WHERE "book_id" = $1 LIMIT 1
`

func (q *FileContentQueries) GetBookUrl(bookID int32) (entities.GetBookUrlRow, error) {
	ctx := context.Background()
	db, _ := pgxpool.Connect(ctx, PgConnection)

	var item entities.GetBookUrlRow
	err := pgxscan.Get(ctx, db, &item, getBookUrl, bookID)

	return item, err
}
