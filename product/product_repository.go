package product

import (
	"context"
	"database/sql"
)

// Repo : Product repository
type Repo interface {
	CreateProduct(context.Context, *CreateProductRequest) (*Product, error)
	GetProduct(context.Context, *GetProductRequest) ([]VariantProductRow, error)
	IsUniqueProduct(context.Context, string) (bool, error)
	ListProduct(context.Context, *ListProductRequest) ([]VariantProductRow, error)
}

// NewRepo : Returns product repo
func NewRepo(db *sql.DB) Repo {
	return &PostgresRepo{
		DB: db,
	}
}
