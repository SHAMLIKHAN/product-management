package category

import (
	"context"
	"database/sql"
	"errors"
	"pm/utils"
)

// PostgresRepo : Category repo struct for postgres
type PostgresRepo struct {
	DB *sql.DB
}

func getNullInt32(value int) sql.NullInt32 {
	if value == 0 {
		return sql.NullInt32{}
	}
	return sql.NullInt32{
		Int32: int32(value),
		Valid: true,
	}
}

// CreateCategory : Postgres function to create a category
func (pg *PostgresRepo) CreateCategory(ctx context.Context, request *CreateCategoryRequest) (*Category, error) {
	var category Category
	var parentID sql.NullInt32
	query := `
		INSERT INTO
			category (name, id_parent, created_at, updated_at)
		VALUES
			(
				$1, $2, NOW(), NOW()
			)
			RETURNING id_category, name, id_parent, created_at, updated_at
	`
	row := pg.DB.QueryRowContext(ctx, query, request.Name, getNullInt32(request.ParentID))
	err := row.Scan(&category.ID, &category.Name, &parentID, &category.CreatedAt, &category.UpdatedAt)
	category.IDParent = int(parentID.Int32)
	return &category, err
}

// IsExistProduct : Postgres function to check is there any undeleted product under the category
func (pg *PostgresRepo) IsExistProduct(ctx context.Context, request *RemoveCategoryRequest) (bool, error) {
	var isExist bool
	query := `
		SELECT
			EXISTS
			(
				SELECT
					1
				FROM
					product
				WHERE
					id_category = $1
					AND deleted_at IS NULL
			)
	`
	err := pg.DB.QueryRowContext(ctx, query, request.CategoryID).Scan(&isExist)
	return isExist, err
}

// IsExistSubCategories : Postgres function to check is there any undeleted sub categories under the category
func (pg *PostgresRepo) IsExistSubCategories(ctx context.Context, request *RemoveCategoryRequest) (bool, error) {
	var isExist bool
	query := `
		SELECT
			EXISTS
			(
				SELECT
					1
				FROM
					category
				WHERE
					id_parent = $1
					AND deleted_at IS NULL
			)
	`
	err := pg.DB.QueryRowContext(ctx, query, request.CategoryID).Scan(&isExist)
	return isExist, err
}

// IsUniqueCategory : Postgres function to verify unique category
func (pg *PostgresRepo) IsUniqueCategory(ctx context.Context, name string) (bool, error) {
	var isUnique bool
	query := `
		SELECT
			EXISTS
			(
				SELECT
					1
				FROM
					category
				WHERE
					name = $1
					AND deleted_at IS NULL
			)
	`
	err := pg.DB.QueryRowContext(ctx, query, name).Scan(&isUnique)
	return isUnique, err
}

// ListCategory : Postgres function to list out all categories and its products
func (pg *PostgresRepo) ListCategory(ctx context.Context, request *ListCategoryRequest) ([]ProductCategoryRow, error) {
	var pcList []ProductCategoryRow
	query := `
		WITH recent_category AS
		(
			SELECT
				*
			FROM
				category
			WHERE
				deleted_at IS NULL
		)
		SELECT
			c1.id_category AS id_parent,
			c1.name AS parent_category_name,
			c2.id_category id_child,
			c2.name AS child_category_name,
			p.id_product,
			p.name AS product_name,
			p.description,
			p.image_url,
			v.id_variant,
			v.name AS variant_name,
			v.max_retail_price,
			v.discounted_price,
			v.size,
			v.colour
		FROM
			recent_category c1
			LEFT JOIN
				recent_category c2
				ON c1.id_category = c2.id_parent
			LEFT JOIN
				product p
				ON p.id_category = c1.id_category
			LEFT JOIN
				variant v
				ON p.id_product = v.id_product
		WHERE
			p.deleted_at IS NULL
			AND v.deleted_at IS NULL
	`
	rows, err := pg.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var pc ProductCategoryRow
		var childID, productID, variantID sql.NullInt32
		var childCategoryName, productName, description, imageURL, variantName, size, colour sql.NullString
		var maxRetailPrice, discountedPrice sql.NullFloat64
		err := rows.Scan(&pc.ParentID, &pc.ParentCategoryName, &childID, &childCategoryName,
			&productID, &productName, &description, &imageURL,
			&variantID, &variantName, &maxRetailPrice, &discountedPrice, &size, &colour)
		if err != nil {
			return nil, err
		}
		if childID.Valid {
			pc.ChildID = int(childID.Int32)
			if childCategoryName.Valid {
				pc.ChildCategoryName = childCategoryName.String
			}
		}
		if productID.Valid {
			pc.ProductID = int(productID.Int32)
			if productName.Valid {
				pc.ProductName = productName.String
			}
			if description.Valid {
				pc.Description = description.String
			}
			if imageURL.Valid {
				pc.ImageURL = imageURL.String
			}
			if variantID.Valid {
				pc.VariantID = int(variantID.Int32)
				if variantName.Valid {
					pc.VariantName = variantName.String
				}
				if maxRetailPrice.Valid {
					pc.MaxRetailPrice = maxRetailPrice.Float64
				}
				if discountedPrice.Valid {
					pc.DiscountedPrice = discountedPrice.Float64
				}
				if size.Valid {
					pc.VariantSize = size.String
				}
				if colour.Valid {
					pc.VaraintColour = colour.String
				}
			}
		}
		pcList = append(pcList, pc)
	}
	return pcList, nil
}

// RemoveCategory : Postgres function to remove a category
func (pg *PostgresRepo) RemoveCategory(ctx context.Context, request *RemoveCategoryRequest) error {
	query := `
		UPDATE
			category
		SET
			deleted_at = NOW()
		WHERE
			id_category = $1
			AND deleted_at IS NULL
	`
	result, err := pg.DB.ExecContext(ctx, query, request.CategoryID)
	if err != nil {
		return err
	}
	updateCount, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return errors.New(utils.InvalidCategoryIDError)
	}
	return nil
}
