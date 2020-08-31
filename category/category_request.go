package category

// CreateCategoryRequest :
type CreateCategoryRequest struct {
	Name     string `json:"name" validate:"required"`
	ParentID int    `json:"id_parent" validate:"omitempty,gt=0"`
}

// UpdateCategoryRequest :
type UpdateCategoryRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	UpdatedBy   int    `json:"updated_by"`
}
