package variant

import (
	"database/sql"
)

// PostgresRepo : Variant repo struct for postgres
type PostgresRepo struct {
	DB *sql.DB
}

func getNullFloat64(value float64) sql.NullFloat64 {
	if value == 0 {
		return sql.NullFloat64{}
	}
	return sql.NullFloat64{
		Float64: value,
		Valid:   true,
	}
}

// CreateVariant : Postgres function to create a variant
func (pg *PostgresRepo) CreateVariant(request *CreateVariantRequest) (*Variant, error) {
	var variant Variant
	var name, size, colour sql.NullString
	var discountedPrice sql.NullFloat64
	query := `
		INSERT INTO
			variant (name, max_retail_price, discounted_price, size, colour, id_product, created_at, updated_at)
		VALUES
			(
				$1, $2, $3, $4, $5, $6, NOW(), NOW()
			)
			RETURNING id_variant, name, max_retail_price, discounted_price, size, colour, id_product, created_at, updated_at
	`
	row := pg.DB.QueryRow(query, request.Name, request.MaxRetailPrice, getNullFloat64(request.DiscountedPrice), request.Size, request.Colour, request.ProductID)
	err := row.Scan(&variant.ID, &name, &variant.MaxRetailPrice, &discountedPrice, &size, &colour, &variant.ProductID,
		&variant.CreatedAt, &variant.UpdatedAt)
	if name.Valid {
		variant.Name = name.String
	}
	if size.Valid {
		variant.Size = size.String
	}
	if colour.Valid {
		variant.Colour = colour.String
	}
	if discountedPrice.Valid {
		variant.DiscountedPrice = discountedPrice.Float64
	}
	return &variant, err
}

// IsValidProductID : Postgres function to validate id_product
func (pg *PostgresRepo) IsValidProductID(productID int) (bool, error) {
	var isValid bool
	query := `
		SELECT
			EXISTS
			(
				SELECT
					1
				FROM
					product
				WHERE
					id_product = $1
					AND deleted_at IS NULL
			)
	`
	err := pg.DB.QueryRow(query, productID).Scan(&isValid)
	return isValid, err
}
