package category

import (
	"database/sql"
	"errors"
	"pm/utils"
)

// ServiceInterface : Category service
type ServiceInterface interface {
	CreateCategory(*CreateCategoryRequest) (*Category, error)
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
func (cs *Service) CreateCategory(request *CreateCategoryRequest) (*Category, error) {
	isUnique, err := cs.cr.IsUniqueCategory(request.Name)
	if err != nil {
		return nil, err
	}
	if isUnique {
		return nil, errors.New(utils.CategoryNameExistsError)
	}
	category, err := cs.cr.CreateCategory(request)
	if err != nil {
		return nil, err
	}
	return category, nil
}
