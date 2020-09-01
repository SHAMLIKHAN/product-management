package category

// CreateCategoryRequest :
type CreateCategoryRequest struct {
	Name     string `json:"name" validate:"required"`
	ParentID int    `json:"id_parent" validate:"omitempty,gt=0"`
}

// ListCategoryRequest :
type ListCategoryRequest struct {
	Limit  int
	Offset int
}

// RemoveCategoryRequest :
type RemoveCategoryRequest struct {
	CategoryID int
}

// UpdateCategoryRequest :
type UpdateCategoryRequest struct {
	CategoryID int
	Name       string `json:"name"`
	ParentID   int    `json:"id_parent" validate:"omitempty,gt=0"`
}
