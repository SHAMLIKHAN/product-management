package product

// CreateProductRequest :
type CreateProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	CategoryID  int    `json:"id_category" validate:"required,gt=0"`
}
