package product

import "database/sql"

// Repo : Product repository
type Repo interface {
	CreateProduct(*CreateProductRequest) (*Product, error)
	IsUniqueProduct(string) (bool, error)
}

// NewRepo : Returns product repo
func NewRepo(db *sql.DB) Repo {
	return &PostgresRepo{
		DB: db,
	}
}
