package category

import (
	"database/sql"
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
func (pg *PostgresRepo) CreateCategory(request *CreateCategoryRequest) (*Category, error) {
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
	row := pg.DB.QueryRow(query, request.Name, getNullInt32(request.ParentID))
	err := row.Scan(&category.ID, &category.Name, &parentID, &category.CreatedAt, &category.UpdatedAt)
	category.IDParent = int(parentID.Int32)
	return &category, err
}

// IsUniqueCategory : Postgres function to verify unique category
func (pg *PostgresRepo) IsUniqueCategory(name string) (bool, error) {
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
	err := pg.DB.QueryRow(query, name).Scan(&isUnique)
	return isUnique, err
}
