package variant

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

// HandlerInterface : Variant handler
type HandlerInterface interface {
	CreateVariant(w http.ResponseWriter, r *http.Request)
}

// Handler : Variant handler struct
type Handler struct {
	vs ServiceInterface
}

// NewHTTPHandler : Returns variant HTTP handler
func NewHTTPHandler(db *sql.DB) HandlerInterface {
	return &Handler{
		vs: NewService(db),
	}
}

// CreateVariant : to create a variant
func (vh *Handler) CreateVariant(w http.ResponseWriter, r *http.Request) {
	log.Println("App : POST /app/product/{id_product}/variant API hit!")
	productID, err := strconv.Atoi(chi.URLParam(r, "id_product"))
	if err != nil {
		log.Println("Error : ", err.Error())
		utils.Fail(w, 400, utils.DecodeErrorCode, errors.New("invalid id_product").Error())
		return
	}
	var request CreateVariantRequest
	body := json.NewDecoder(r.Body)
	err = body.Decode(&request)
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
	request.ProductID = productID
	variant, err := vh.vs.CreateVariant(&request)
	if err != nil {
		log.Println("Error : ", err.Error())
		if err.Error() == utils.IDProductDoesNotExistError {
			utils.Fail(w, 400, utils.IDProductDoesNotExistErrorCode, err.Error())
			return
		}
		utils.Fail(w, 500, utils.DatabaseErrorCode, err.Error())
		return
	}
	log.Println("App : variant created! id_variant : ", variant.ID)
	utils.Send(w, 200, variant)
}
