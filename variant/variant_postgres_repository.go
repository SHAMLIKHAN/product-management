package variant

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"pm/utils"
	"strings"
	"time"
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
func (pg *PostgresRepo) CreateVariant(ctx context.Context, request *CreateVariantRequest) (*Variant, error) {
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
	row := pg.DB.QueryRowContext(ctx, query, request.Name, request.MaxRetailPrice, getNullFloat64(request.DiscountedPrice), request.Size, request.Colour, request.ProductID)
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

// GetVariant : Postgres function to get variant of a product
func (pg *PostgresRepo) GetVariant(ctx context.Context, request *GetVariantRequest) (*Variant, error) {
	var variant Variant
	var name, size, colour sql.NullString
	var discountedPrice sql.NullFloat64
	query := `
		SELECT
			name, max_retail_price, discounted_price, size, colour, created_at, updated_at
		FROM
			variant
		WHERE
			id_variant = $1
			AND id_product = $2
			AND deleted_at IS NULL
	`
	row := pg.DB.QueryRowContext(ctx, query, request.VariantID, request.ProductID)
	err := row.Scan(&name, &variant.MaxRetailPrice, &discountedPrice, &size, &colour, &variant.CreatedAt, &variant.UpdatedAt)
	if err != nil {
		return nil, err
	}
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
	variant.ID = request.VariantID
	variant.ProductID = request.ProductID
	return &variant, nil
}

// IsValidProductID : Postgres function to validate id_product
func (pg *PostgresRepo) IsValidProductID(ctx context.Context, productID int) (bool, error) {
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
	err := pg.DB.QueryRowContext(ctx, query, productID).Scan(&isValid)
	return isValid, err
}

// ListVariant : Postgres function to list variants of a product
func (pg *PostgresRepo) ListVariant(ctx context.Context, request *ListVariantRequest) ([]Variant, error) {
	var variants []Variant
	var name, size, colour sql.NullString
	var discountedPrice sql.NullFloat64
	query := `
		SELECT
			id_variant, name, max_retail_price, discounted_price, size, colour, created_at, updated_at
		FROM
			variant
		WHERE
			id_product = $1
			AND deleted_at IS NULL
		LIMIT $2
		OFFSET $3
	`
	rows, err := pg.DB.QueryContext(ctx, query, request.ProductID, request.Limit, request.Offset)
	if err != nil {
		return nil, err
	}
	var variant Variant
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&variant.ID, &name, &variant.MaxRetailPrice, &discountedPrice, &size, &colour, &variant.CreatedAt, &variant.UpdatedAt)
		if err != nil {
			return nil, err
		}
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
		variant.ProductID = request.ProductID
		variants = append(variants, variant)
	}
	return variants, nil
}

// RemoveVariant : Postgres function to remove a variant of a product
func (pg *PostgresRepo) RemoveVariant(ctx context.Context, request *RemoveVariantRequest) error {
	query := `
		UPDATE
			variant
		SET
			deleted_at = NOW()
		WHERE
			id_product = $1
			AND id_variant = $2
			AND deleted_at IS NULL
	`
	result, err := pg.DB.ExecContext(ctx, query, request.ProductID, request.VariantID)
	if err != nil {
		return err
	}
	updateCount, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return errors.New(utils.InvalidVariantIDError)
	}
	return nil
}

// UpdateVariant : Postgres function to update variant of a product
func (pg *PostgresRepo) UpdateVariant(ctx context.Context, request *UpdateVariantRequest, columns map[string]interface{}) error {
	var slice []string
	paramsPosition := 0
	columns["updated_at"] = time.Now()
	var params []interface{}
	for column, value := range columns {
		paramsPosition++
		slice = append(slice, fmt.Sprintf(" %s = $%d ", column, paramsPosition))
		params = append(params, value)
	}
	update := strings.Join(slice, ", ")
	updateQuery := fmt.Sprintf("%s", update)
	paramsPosition++
	mainQuery := `
		UPDATE
			variant
		SET
			%s
		WHERE
			id_product = $%d
			AND id_variant = $%d
			AND deleted_at IS NULL
	`
	query := fmt.Sprintf(mainQuery, updateQuery, paramsPosition, paramsPosition+1)
	params = append(params, request.ProductID, request.VariantID)
	result, err := pg.DB.ExecContext(ctx, query, params...)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(utils.InvalidVariantIDError)
	}
	return err
}
