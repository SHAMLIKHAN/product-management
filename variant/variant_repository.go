package variant

import "database/sql"

// Repo : Variant repository
type Repo interface {
	CreateVariant(*CreateVariantRequest) (*Variant, error)
	IsValidProductID(int) (bool, error)
}

// NewRepo : Returns variant repo
func NewRepo(db *sql.DB) Repo {
	return &PostgresRepo{
		DB: db,
	}
}
