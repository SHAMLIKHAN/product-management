package product

import "time"

// Product : product struct
type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	IDCategory  int       `json:"id_category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Variant : variant struct
type Variant struct {
	ID              int     `json:"id"`
	Name            string  `json:"name,omitempty"`
	MaxRetailPrice  float64 `json:"max_retail_price"`
	DiscountedPrice float64 `json:"discounted_price,omitempty"`
	Size            string  `json:"size,omitempty"`
	Colour          string  `json:"colour,omitempty"`
}

// VariantProduct : product struct with variants
type VariantProduct struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	IDCategory  int       `json:"id_category"`
	Variants    []Variant `json:"variants"`
}

// VariantProductRow : product variant row struct
type VariantProductRow struct {
	IDProduct       int
	ProductName     string
	Description     string
	ImageURL        string
	IDCategory      int
	IDVariant       int
	VariantName     string
	MaxRetailPrice  float64
	DiscountedPrice float64
	VariantSize     string
	VariantColour   string
}
