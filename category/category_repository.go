package category

import "database/sql"

// Repo : Category repository
type Repo interface {
	CreateCategory(*CreateCategoryRequest) (*Category, error)
	IsUniqueCategory(string) (bool, error)
}

// NewRepo : Returns category repo
func NewRepo(db *sql.DB) Repo {
	return &PostgresRepo{
		DB: db,
	}
}
