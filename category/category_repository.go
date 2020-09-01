package category

import (
	"context"
	"database/sql"
)

// Repo : Category repository
type Repo interface {
	CreateCategory(context.Context, *CreateCategoryRequest) (*Category, error)
	IsExistCategory(context.Context, *UpdateCategoryRequest) (bool, error)
	IsExistProduct(context.Context, *RemoveCategoryRequest) (bool, error)
	IsExistSubCategories(context.Context, *RemoveCategoryRequest) (bool, error)
	IsUniqueCategory(context.Context, string) (bool, error)
	ListCategory(context.Context, *ListCategoryRequest) ([]ProductCategoryRow, error)
	RemoveCategory(context.Context, *RemoveCategoryRequest) error
	UpdateCategory(context.Context, *UpdateCategoryRequest, map[string]interface{}) error
}

// NewRepo : Returns category repo
func NewRepo(db *sql.DB) Repo {
	return &PostgresRepo{
		DB: db,
	}
}
