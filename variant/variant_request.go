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

// ListVariantRequest :
type ListVariantRequest struct {
	ProductID int
	Limit     int
	Offset    int
}
