package product

import (
	"context"
	"database/sql"
)

// PostgresRepo : Product repo struct for postgres
type PostgresRepo struct {
	DB *sql.DB
}

// CreateProduct : Postgres function to create a product
func (pg *PostgresRepo) CreateProduct(ctx context.Context, request *CreateProductRequest) (*Product, error) {
	var product Product
	var description, ImageURL sql.NullString
	query := `
		INSERT INTO
			product (name, description, image_url, id_category, created_at, updated_at)
		VALUES
			(
				$1, $2, $3, $4, NOW(), NOW()
			)
			RETURNING id_product, name, description, image_url, id_category, created_at, updated_at
	`
	row := pg.DB.QueryRowContext(ctx, query, request.Name, request.Description, request.ImageURL, request.CategoryID)
	err := row.Scan(&product.ID, &product.Name, &description, &ImageURL, &product.IDCategory, &product.CreatedAt, &product.UpdatedAt)
	product.Description = description.String
	product.ImageURL = ImageURL.String
	return &product, err
}

// IsUniqueProduct : Postgres function to verify unique product
func (pg *PostgresRepo) IsUniqueProduct(ctx context.Context, name string) (bool, error) {
	var isUnique bool
	query := `
		SELECT
			EXISTS
			(
				SELECT
					1
				FROM
					product
				WHERE
					name = $1
					AND deleted_at IS NULL
			)
	`
	err := pg.DB.QueryRowContext(ctx, query, name).Scan(&isUnique)
	return isUnique, err
}

// ListProduct : Postgres function to list out products
func (pg *PostgresRepo) ListProduct(ctx context.Context, request *ListProductRequest) ([]VariantProductRow, error) {
	var vpList []VariantProductRow
	var description, imageURL, variantName, size, colour sql.NullString
	var maxRetailPrice, discountedPrice sql.NullFloat64
	var variantID sql.NullInt32
	query := `
		WITH recent_product AS
		(
			SELECT
				*
			FROM
				product
			WHERE
				deleted_at IS NULL
			LIMIT $1
			OFFSET $2
		)
		SELECT
			p.id_product,
			p.name AS product_name,
			p.description,
			p.image_url,
			p.id_category,
			v.id_variant,
			v.name AS variant_name,
			v.max_retail_price,
			v.discounted_price,
			v.size,
			v.colour
		FROM
			recent_product p
			LEFT JOIN
				variant v
				ON p.id_product = v.id_product
		WHERE
			v.deleted_at IS NULL
	`
	rows, err := pg.DB.QueryContext(ctx, query, request.Limit, request.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var vp VariantProductRow
	for rows.Next() {
		err := rows.Scan(&vp.IDProduct, &vp.ProductName, &description, &imageURL, &vp.IDCategory,
			&variantID, &variantName, &maxRetailPrice, &discountedPrice, &size, &colour)
		if err != nil {
			return nil, err
		}
		if description.Valid {
			vp.Description = description.String
		}
		if imageURL.Valid {
			vp.ImageURL = imageURL.String
		}
		if variantID.Valid {
			vp.IDVariant = int(variantID.Int32)
			if variantName.Valid {
				vp.VariantName = variantName.String
			}
			if size.Valid {
				vp.VariantSize = size.String
			}
			if colour.Valid {
				vp.VariantColour = colour.String
			}
			if maxRetailPrice.Valid {
				vp.MaxRetailPrice = maxRetailPrice.Float64
			}
			if discountedPrice.Valid {
				vp.DiscountedPrice = discountedPrice.Float64
			}
		}
		vpList = append(vpList, vp)
	}
	return vpList, nil
}
