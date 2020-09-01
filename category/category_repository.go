package category

import (
	"context"
	"database/sql"
)

// Repo : Category repository
type Repo interface {
	CreateCategory(context.Context, *CreateCategoryRequest) (*Category, error)
	IsExistProduct(context.Context, *RemoveCategoryRequest) (bool, error)
	IsExistSubCategories(context.Context, *RemoveCategoryRequest) (bool, error)
	IsUniqueCategory(context.Context, string) (bool, error)
	ListCategory(context.Context, *ListCategoryRequest) ([]ProductCategoryRow, error)
	RemoveCategory(context.Context, *RemoveCategoryRequest) error
}

// NewRepo : Returns category repo
func NewRepo(db *sql.DB) Repo {
	return &PostgresRepo{
		DB: db,
	}
}
