package product

import (
	"context"
	"database/sql"
)

// Repo : Product repository
type Repo interface {
	CreateProduct(context.Context, *CreateProductRequest) (*Product, error)
	IsUniqueProduct(context.Context, string) (bool, error)
}

// NewRepo : Returns product repo
func NewRepo(db *sql.DB) Repo {
	return &PostgresRepo{
		DB: db,
	}
}
