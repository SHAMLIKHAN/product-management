package category

// CreateCategoryRequest :
type CreateCategoryRequest struct {
	Name     string `json:"name" validate:"required"`
	ParentID int    `json:"id_parent" validate:"omitempty,gt=0"`
}
