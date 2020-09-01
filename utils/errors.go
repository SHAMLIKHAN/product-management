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
	// ProductNameExistsErrorCode :
	ProductNameExistsErrorCode = 121
	// IDProductDoesNotExistErrorCode :
	IDProductDoesNotExistErrorCode = 122
	// InvalidVariantIDErrorCode :
	InvalidVariantIDErrorCode = 131

	// CategoryNameExistsError :
	CategoryNameExistsError = "category name exists"
	// ProductNameExistsError :
	ProductNameExistsError = "product name exists"
	// IDProductDoesNotExistError :
	IDProductDoesNotExistError = "id_product doesn't exist"
	// InvalidVariantIDError :
	InvalidVariantIDError = "invalid id_variant"
)
