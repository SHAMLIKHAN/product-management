package product

import (
	"context"
	"database/sql"
	"errors"
	"pm/utils"
)

// ServiceInterface : Product service
type ServiceInterface interface {
	CreateProduct(context.Context, *CreateProductRequest) (*Product, error)
	GetProduct(context.Context, *GetProductRequest) (*VariantProduct, error)
	ListProduct(context.Context, *ListProductRequest) ([]VariantProduct, error)
	RemoveProduct(context.Context, *RemoveProductRequest) error
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
func (ps *Service) CreateProduct(ctx context.Context, request *CreateProductRequest) (*Product, error) {
	isUnique, err := ps.pr.IsUniqueProduct(ctx, request.Name)
	if err != nil {
		return nil, err
	}
	if isUnique {
		return nil, errors.New(utils.ProductNameExistsError)
	}
	product, err := ps.pr.CreateProduct(ctx, request)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// GetProduct : to get a product
func (ps *Service) GetProduct(ctx context.Context, request *GetProductRequest) (*VariantProduct, error) {
	rows, err := ps.pr.GetProduct(ctx, request)
	if err != nil {
		return nil, err
	}
	var (
		product  VariantProduct
		variant  Variant
		variants []Variant
	)
	for _, row := range rows {
		if row.IDVariant != 0 {
			variant.ID = row.IDVariant
			variant.Name = row.VariantName
			variant.MaxRetailPrice = row.MaxRetailPrice
			variant.DiscountedPrice = row.DiscountedPrice
			variant.Size = row.VariantSize
			variant.Colour = row.VariantColour
			variants = append(variants, variant)
		}
	}
	product.ID = rows[0].IDProduct
	product.Name = rows[0].ProductName
	product.Description = rows[0].Description
	product.ImageURL = rows[0].ImageURL
	product.IDCategory = rows[0].IDCategory
	product.Variants = variants
	return &product, nil
}

// ListProduct : to list out products
func (ps *Service) ListProduct(ctx context.Context, request *ListProductRequest) ([]VariantProduct, error) {
	rows, err := ps.pr.ListProduct(ctx, request)
	if err != nil {
		return nil, err
	}
	var productList []VariantProduct
	if len(rows) > 0 {
		for _, row := range rows {
			var (
				product    VariantProduct
				variant    Variant
				variants   []Variant
				isAppended bool
			)
			for index, product := range productList {
				if product.ID == row.IDProduct {
					variant.ID = row.IDVariant
					variant.Name = row.VariantName
					variant.MaxRetailPrice = row.MaxRetailPrice
					variant.DiscountedPrice = row.DiscountedPrice
					variant.Size = row.VariantSize
					variant.Colour = row.VariantColour
					product.Variants = append(product.Variants, variant)
					productList[index] = product
					isAppended = true
					break
				}
			}
			if isAppended {
				continue
			}
			if row.IDVariant != 0 {
				variant.ID = row.IDVariant
				variant.Name = row.VariantName
				variant.MaxRetailPrice = row.MaxRetailPrice
				variant.DiscountedPrice = row.DiscountedPrice
				variant.Size = row.VariantSize
				variant.Colour = row.VariantColour
				variants = append(variants, variant)
			}
			product.ID = row.IDProduct
			product.Name = row.ProductName
			product.Description = row.Description
			product.ImageURL = row.ImageURL
			product.IDCategory = row.IDCategory
			product.Variants = variants
			productList = append(productList, product)
		}
	}
	return productList, nil
}

// RemoveProduct : to remove a product and its variants
func (ps *Service) RemoveProduct(ctx context.Context, request *RemoveProductRequest) error {
	return ps.pr.RemoveProduct(ctx, request)
}
