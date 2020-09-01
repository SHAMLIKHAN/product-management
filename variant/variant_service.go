package variant

import (
	"context"
	"database/sql"
	"errors"
	"pm/utils"
)

// ServiceInterface : Variant service
type ServiceInterface interface {
	CreateVariant(context.Context, *CreateVariantRequest) (*Variant, error)
	GetVariant(context.Context, *GetVariantRequest) (*Variant, error)
	ListVariant(context.Context, *ListVariantRequest) ([]Variant, error)
	RemoveVariant(context.Context, *RemoveVariantRequest) error
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
func (vs *Service) CreateVariant(ctx context.Context, request *CreateVariantRequest) (*Variant, error) {
	isValid, err := vs.vr.IsValidProductID(ctx, request.ProductID)
	if err != nil {
		return nil, err
	}
	if !isValid {
		return nil, errors.New(utils.IDProductDoesNotExistError)
	}
	return vs.vr.CreateVariant(ctx, request)
}

// GetVariant : to get a variant of a product
func (vs *Service) GetVariant(ctx context.Context, request *GetVariantRequest) (*Variant, error) {
	isValid, err := vs.vr.IsValidProductID(ctx, request.ProductID)
	if err != nil {
		return nil, err
	}
	if !isValid {
		return nil, errors.New(utils.IDProductDoesNotExistError)
	}
	return vs.vr.GetVariant(ctx, request)
}

// ListVariant : to list out all variants of a product
func (vs *Service) ListVariant(ctx context.Context, request *ListVariantRequest) ([]Variant, error) {
	isValid, err := vs.vr.IsValidProductID(ctx, request.ProductID)
	if err != nil {
		return nil, err
	}
	if !isValid {
		return nil, errors.New(utils.IDProductDoesNotExistError)
	}
	return vs.vr.ListVariant(ctx, request)
}

// RemoveVariant : to remove variant of a product
func (vs *Service) RemoveVariant(ctx context.Context, request *RemoveVariantRequest) error {
	isValid, err := vs.vr.IsValidProductID(ctx, request.ProductID)
	if err != nil {
		return err
	}
	if !isValid {
		return errors.New(utils.IDProductDoesNotExistError)
	}
	return vs.vr.RemoveVariant(ctx, request)
}
