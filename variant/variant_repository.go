package variant

import (
	"context"
	"database/sql"
)

// Repo : Variant repository
type Repo interface {
	CreateVariant(context.Context, *CreateVariantRequest) (*Variant, error)
	GetVariant(context.Context, *GetVariantRequest) (*Variant, error)
	IsValidProductID(context.Context, int) (bool, error)
	ListVariant(context.Context, *ListVariantRequest) ([]Variant, error)
	RemoveVariant(context.Context, *RemoveVariantRequest) error
	UpdateVariant(context.Context, *UpdateVariantRequest, map[string]interface{}) error
}

// NewRepo : Returns variant repo
func NewRepo(db *sql.DB) Repo {
	return &PostgresRepo{
		DB: db,
	}
}
