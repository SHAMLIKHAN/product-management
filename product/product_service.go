package product

import (
	"database/sql"
	"errors"
	"pm/utils"
)

// ServiceInterface : Product service
type ServiceInterface interface {
	CreateProduct(*CreateProductRequest) (*Product, error)
}

// Service : Product service struct
type Service struct {
	pr Repo
}

// NewService : Returns product service
func NewService(db *sql.DB) ServiceInterface {
	return &Service{
		pr: NewRepo(db),
	}
}

// CreateProduct : to create a product
func (ps *Service) CreateProduct(request *CreateProductRequest) (*Product, error) {
	isUnique, err := ps.pr.IsUniqueProduct(request.Name)
	if err != nil {
		return nil, err
	}
	if isUnique {
		return nil, errors.New(utils.ProductNameExistsError)
	}
	product, err := ps.pr.CreateProduct(request)
	if err != nil {
		return nil, err
	}
	return product, nil
}
