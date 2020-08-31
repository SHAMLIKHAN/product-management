package category

import (
	"context"
	"database/sql"
	"errors"
	"pm/utils"
)

// ServiceInterface : Category service
type ServiceInterface interface {
	CreateCategory(context.Context, *CreateCategoryRequest) (*Category, error)
}

// Service : Category service struct
type Service struct {
	cr Repo
}

// NewService : Returns category service
func NewService(db *sql.DB) ServiceInterface {
	return &Service{
		cr: NewRepo(db),
	}
}

// CreateCategory : to create a category
func (cs *Service) CreateCategory(ctx context.Context, request *CreateCategoryRequest) (*Category, error) {
	isUnique, err := cs.cr.IsUniqueCategory(ctx, request.Name)
	if err != nil {
		return nil, err
	}
	if isUnique {
		return nil, errors.New(utils.CategoryNameExistsError)
	}
	category, err := cs.cr.CreateCategory(ctx, request)
	if err != nil {
		return nil, err
	}
	return category, nil
}
