package variant

import (
	"context"
	"database/sql"
)

// Repo : Variant repository
type Repo interface {
	CreateVariant(context.Context, *CreateVariantRequest) (*Variant, error)
	IsValidProductID(context.Context, int) (bool, error)
}

// NewRepo : Returns variant repo
func NewRepo(db *sql.DB) Repo {
	return &PostgresRepo{
		DB: db,
	}
}
