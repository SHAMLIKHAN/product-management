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
	ListCategory(context.Context, *ListCategoryRequest) ([]ProductCategory, error)
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

// ListCategory : to list out all categories and its products
func (cs *Service) ListCategory(ctx context.Context, request *ListCategoryRequest) ([]ProductCategory, error) {
	rows, err := cs.cr.ListCategory(ctx, request)
	if err != nil {
		return nil, err
	}
	categoryList := formatProductCategoryRows(rows)
	return categoryList, nil
}

func formatProductCategoryRows(rows []ProductCategoryRow) []ProductCategory {
	var categories, categoryList []ProductCategory
	var subCategoryIDs []int
	for _, row := range rows {
		var (
			category           ProductCategory
			product            Product
			variant            Variant
			isCategoryAppended bool
		)
		variant.VariantID = row.VariantID
		variant.VariantName = row.VariantName
		variant.MaxRetailPrice = row.MaxRetailPrice
		variant.DiscountedPrice = row.DiscountedPrice
		variant.VariantSize = row.VariantSize
		variant.VaraintColour = row.VaraintColour
		for _, c := range categories {
			var isProductAppended bool
			if row.ParentID == c.CategoryID {
				for ip, p := range c.Products {
					if row.ProductID == p.ProductID {
						if variant.VariantID != 0 {
							p.Variants = append(p.Variants, variant)
							c.Products[ip] = p
						}
						isProductAppended = true
						break
					}
					if isProductAppended {
						continue
					}
					product.ProductID = row.ProductID
					product.Description = row.Description
					product.ImageURL = row.ImageURL
					product.CategoryID = row.ParentID
					if variant.VariantID != 0 {
						product.Variants = append(product.Variants, variant)
					}
					if product.ProductID != 0 {
						c.Products = append(c.Products, product)
					}
				}
				isCategoryAppended = true
				break
			}
		}
		if isCategoryAppended {
			continue
		}
		product.ProductID = row.ProductID
		product.Description = row.Description
		product.ImageURL = row.ImageURL
		product.CategoryID = row.ParentID
		if variant.VariantID != 0 {
			product.Variants = append(product.Variants, variant)
		}
		category.CategoryID = row.ParentID
		category.CategoryName = row.ParentCategoryName
		if row.ChildID != 0 {
			subCategoryIDs = append(subCategoryIDs, row.ChildID)
			var isChildAppended bool
			for _, id := range category.ParentOf {
				if id == row.ChildID {
					isChildAppended = true
					break
				}
			}
			if !isChildAppended {
				category.ParentOf = append(category.ParentOf, row.ChildID)
			}
		}
		if product.ProductID != 0 {
			category.Products = append(category.Products, product)
		}
		categories = append(categories, category)
	}
	for _, category := range categories {
		formatCategories(category, categories)
	}
	for _, category := range categories {
		var isChild bool
		for _, id := range subCategoryIDs {
			if category.CategoryID == id {
				isChild = true
				break
			}
		}
		if !isChild {
			categoryList = append(categoryList, category)
		}
	}
	return categoryList
}

func formatCategories(category ProductCategory, categories []ProductCategory) {
	for index, c := range categories {
		if c.CategoryID == category.CategoryID {
			continue
		}
		if len(c.ParentOf) == 0 {
			continue
		}
		for _, id := range c.ParentOf {
			if id == category.CategoryID {
				category.IsVisited = true
				c.Categories = append(c.Categories, category)
				categories[index] = c
				return
			}
		}
		formatCategories(category, c.Categories)
	}
	return
}
