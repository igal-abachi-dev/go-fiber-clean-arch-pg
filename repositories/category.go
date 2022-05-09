package repositories

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"go-fiber-clean-arch-pg/entities"
)

const getAllCategories = `-- name: GetAllCategories :many
SELECT Name, Level, id FROM public."Category"
ORDER BY id ASC
`

func (q *Queries) GetAllCategories() ([]*entities.Category, error) {
	ctx := context.Background()
	db, _ := pgxpool.Connect(ctx, PgConnection)

	var items []*entities.Category
	err := pgxscan.Select(ctx, db, &items, getAllCategories)

	if err != nil {
		return nil, err
	}
	return items, nil
}
