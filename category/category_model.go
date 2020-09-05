package category

import "time"

// Category : category struct
type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	IDParent  int       `json:"id_parent"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ProductCategory : product category struct
type ProductCategory struct {
	CategoryID   int               `json:"category_id"`
	CategoryName string            `json:"category_name"`
	Categories   []ProductCategory `json:"categories"`
	ParentOf     []int             `json:"-"`
	ParentID     int               `json:"-"`
	Products     []Product         `json:"products"`
	IsVisited    bool              `json:"-"`
}

// Product : product struct
type Product struct {
	ProductID   int       `json:"product_id"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	Variants    []Variant `json:"variants"`
	CategoryID  int       `json:"-"`
}

// Variant : variant struct
type Variant struct {
	VariantID       int     `json:"varaint_id"`
	VariantName     string  `json:"variant_name"`
	MaxRetailPrice  float64 `json:"max_retail_price"`
	DiscountedPrice float64 `json:"discounted_price"`
	VariantSize     string  `json:"size"`
	VaraintColour   string  `json:"colour"`
}

// ProductCategoryRow : product category row struct
type ProductCategoryRow struct {
	ParentID           int
	ParentCategoryName string
	ChildID            int
	ChildCategoryName  string
	ProductName        string
	ProductID          int
	Description        string
	ImageURL           string
	VariantID          int
	VariantName        string
	MaxRetailPrice     float64
	DiscountedPrice    float64
	VariantSize        string
	VaraintColour      string
}
