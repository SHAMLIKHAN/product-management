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
