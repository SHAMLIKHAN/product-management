package category

import (
	"context"
	"database/sql"
)

// Repo : Category repository
type Repo interface {
	CreateCategory(context.Context, *CreateCategoryRequest) (*Category, error)
	IsUniqueCategory(context.Context, string) (bool, error)
	ListCategory(context.Context, *ListCategoryRequest) ([]ProductCategoryRow, error)
}

// NewRepo : Returns category repo
func NewRepo(db *sql.DB) Repo {
	return &PostgresRepo{
		DB: db,
	}
}
