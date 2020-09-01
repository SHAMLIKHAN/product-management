package product

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"pm/utils"
	"strconv"

	"github.com/go-chi/chi"
	"gopkg.in/go-playground/validator.v9"
)

// HandlerInterface : Product handler
type HandlerInterface interface {
	CreateProduct(w http.ResponseWriter, r *http.Request)
	GetProduct(w http.ResponseWriter, r *http.Request)
	ListProduct(w http.ResponseWriter, r *http.Request)
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

// GetProduct : to get a product
func (ph *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("App : GET /app/product/{id_product} API hit!")
	productID, err := strconv.Atoi(chi.URLParam(r, "id_product"))
	if err != nil {
		log.Println("Error : ", err.Error())
		utils.Fail(w, 400, utils.DecodeErrorCode, errors.New("invalid id_product").Error())
		return
	}
	request := GetProductRequest{
		ProductID: productID,
	}
	product, err := ph.ps.GetProduct(r.Context(), &request)
	if err != nil {
		log.Println("Error : ", err.Error())
		utils.Fail(w, 500, utils.DatabaseErrorCode, err.Error())
		return
	}
	log.Println("App : products fetched! id_product : ", productID)
	utils.Send(w, 200, product)
}

// ListProduct : to list out all products
func (ph *Handler) ListProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("App : GET /app/product API hit!")
	req := ListProductRequest{
		Limit:  DefaultListProductLimit,
		Offset: DefaultOffset,
	}
	request, err := validateListProductRequest(&req, r)
	if err != nil {
		log.Println("Error : ", err.Error())
		utils.Fail(w, 400, utils.ValidationErrorCode, err.Error())
		return
	}
	products, err := ph.ps.ListProduct(r.Context(), request)
	if err != nil {
		log.Println("Error : ", err.Error())
		utils.Fail(w, 500, utils.DatabaseErrorCode, err.Error())
		return
	}
	log.Println("App : products listed!")
	utils.Send(w, 200, products)
}
