package variant

// CreateVariantRequest :
type CreateVariantRequest struct {
	Name            string  `json:"name"`
	MaxRetailPrice  float64 `json:"max_retail_price" validate:"required"`
	DiscountedPrice float64 `json:"discounted_price"`
	Size            string  `json:"size"`
	Colour          string  `json:"colour"`
	ProductID       int     `json:"id_product"`
}

// GetVariantRequest :
type GetVariantRequest struct {
	ProductID int
	VariantID int
}

// ListVariantRequest :
type ListVariantRequest struct {
	ProductID int
	Limit     int
	Offset    int
}

// RemoveVariantRequest :
type RemoveVariantRequest struct {
	ProductID int
	VariantID int
}

// UpdateVariantRequest :
type UpdateVariantRequest struct {
	ProductID       int
	VariantID       int
	Name            string  `json:"name,omitempty"`
	MaxRetailPrice  float64 `json:"max_retail_price"`
	DiscountedPrice float64 `json:"discounted_price,omitempty"`
	Size            string  `json:"size,omitempty"`
	Colour          string  `json:"colour,omitempty"`
}
