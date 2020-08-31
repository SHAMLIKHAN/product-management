package variant

import (
	"database/sql"
	"errors"
	"pm/utils"
)

// ServiceInterface : Variant service
type ServiceInterface interface {
	CreateVariant(*CreateVariantRequest) (*Variant, error)
}

// Service : Variant service struct
type Service struct {
	vr Repo
}

// NewService : Returns variant service
func NewService(db *sql.DB) ServiceInterface {
	return &Service{
		vr: NewRepo(db),
	}
}

// CreateVariant : to create a variant
func (vs *Service) CreateVariant(request *CreateVariantRequest) (*Variant, error) {
	isValid, err := vs.vr.IsValidProductID(request.ProductID)
	if err != nil {
		return nil, err
	}
	if !isValid {
		return nil, errors.New(utils.IDProductDoesNotExistError)
	}
	return vs.vr.CreateVariant(request)
}
