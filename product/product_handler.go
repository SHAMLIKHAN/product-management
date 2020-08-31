package product

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"pm/utils"

	"gopkg.in/go-playground/validator.v9"
)

// HandlerInterface : Product handler
type HandlerInterface interface {
	CreateProduct(w http.ResponseWriter, r *http.Request)
}

// Handler : Product handler struct
type Handler struct {
	ps ServiceInterface
}

// NewHTTPHandler : Returns product HTTP handler
func NewHTTPHandler(db *sql.DB) HandlerInterface {
	return &Handler{
		ps: NewService(db),
	}
}

// CreateProduct : to create a product
func (ph *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("App : POST /app/product API hit!")
	var request CreateProductRequest
	body := json.NewDecoder(r.Body)
	err := body.Decode(&request)
	if err != nil {
		log.Println("Error : ", err.Error())
		utils.Fail(w, 400, utils.DecodeErrorCode, err.Error())
		return
	}
	validator := validator.New()
	err = validator.Struct(request)
	if err != nil {
		log.Println("Error : ", err.Error())
		utils.Fail(w, 400, utils.ValidationErrorCode, err.Error())
		return
	}
	product, err := ph.ps.CreateProduct(r.Context(), &request)
	if err != nil {
		log.Println("Error : ", err.Error())
		if err.Error() == utils.ProductNameExistsError {
			utils.Fail(w, 200, utils.ProductNameExistsErrorCode, err.Error())
			return
		}
		utils.Fail(w, 500, utils.DatabaseErrorCode, err.Error())
		return
	}
	log.Println("App : product created! id_product : ", product.ID)
	utils.Send(w, 200, product)
}
