package variant

import "time"

// Variant : variant struct
type Variant struct {
	ID              int       `json:"id"`
	Name            string    `json:"name,omitempty"`
	MaxRetailPrice  float64   `json:"max_retail_price"`
	DiscountedPrice float64   `json:"discounted_price,omitempty"`
	Size            string    `json:"size,omitempty"`
	Colour          string    `json:"colour,omitempty"`
	ProductID       int       `json:"id_product"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
