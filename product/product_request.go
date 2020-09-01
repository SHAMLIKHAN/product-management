package product

// CreateProductRequest :
type CreateProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	CategoryID  int    `json:"id_category" validate:"required,gt=0"`
}

// GetProductRequest :
type GetProductRequest struct {
	ProductID int
}

// ListProductRequest :
type ListProductRequest struct {
	Limit  int
	Offset int
}
