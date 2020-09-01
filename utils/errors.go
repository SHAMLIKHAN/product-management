package utils

const (
	// DecodeErrorCode : error decoding request
	DecodeErrorCode = 101
	// ValidationErrorCode : error validating request
	ValidationErrorCode = 102
	// DatabaseErrorCode : error processing database
	DatabaseErrorCode = 103
	// CategoryNameExistsErrorCode :
	CategoryNameExistsErrorCode = 111
	// SomeProductsAreBelongsToCategoryErrorCode :
	SomeProductsAreBelongsToCategoryErrorCode = 112
	// SomeSubCategoriesAreBelongsToCategoryErrorCode :
	SomeSubCategoriesAreBelongsToCategoryErrorCode = 113
	// InvalidCategoryIDErrorCode :
	InvalidCategoryIDErrorCode = 114
	// ProductNameExistsErrorCode :
	ProductNameExistsErrorCode = 121
	// IDProductDoesNotExistErrorCode :
	IDProductDoesNotExistErrorCode = 122
	// InvalidProductIDErrorCode :
	InvalidProductIDErrorCode = 123
	// InvalidVariantIDErrorCode :
	InvalidVariantIDErrorCode = 131
	// NothingToUpdateVariantErrorCode :
	NothingToUpdateVariantErrorCode = 132

	// CategoryNameExistsError :
	CategoryNameExistsError = "category name exists"
	// ProductNameExistsError :
	ProductNameExistsError = "product name exists"
	// IDProductDoesNotExistError :
	IDProductDoesNotExistError = "id_product doesn't exist"
	// InvalidVariantIDError :
	InvalidVariantIDError = "invalid id_variant"
	// InvalidProductIDError :
	InvalidProductIDError = "invalid id_product"
	// SomeProductsAreBelongsToCategoryError :
	SomeProductsAreBelongsToCategoryError = "some products are belongs to this category, category can't be removed"
	// SomeSubCategoriesAreBelongsToCategoryError :
	SomeSubCategoriesAreBelongsToCategoryError = "some sub categories are belongs to this category, category can't be removed"
	// InvalidCategoryIDError :
	InvalidCategoryIDError = "invalid id_category"
	// NothingToUpdateVariantError :
	NothingToUpdateVariantError = "nothing to update variant"
)
