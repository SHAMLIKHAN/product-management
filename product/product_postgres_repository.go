package product

import (
	"database/sql"
)

// PostgresRepo : Product repo struct for postgres
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

// CreateProduct : Postgres function to create a product
func (pg *PostgresRepo) CreateProduct(request *CreateProductRequest) (*Product, error) {
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
	row := pg.DB.QueryRow(query, request.Name, request.Description, request.ImageURL, request.CategoryID)
	err := row.Scan(&product.ID, &product.Name, &description, &ImageURL, &product.IDCategory, &product.CreatedAt, &product.UpdatedAt)
	product.Description = description.String
	product.ImageURL = ImageURL.String
	return &product, err
}

// IsUniqueProduct : Postgres function to verify unique product
func (pg *PostgresRepo) IsUniqueProduct(name string) (bool, error) {
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
	err := pg.DB.QueryRow(query, name).Scan(&isUnique)
	return isUnique, err
}