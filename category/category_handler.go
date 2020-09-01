package category

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

// HandlerInterface : Category handler
type HandlerInterface interface {
	CreateCategory(w http.ResponseWriter, r *http.Request)
	ListCategory(w http.ResponseWriter, r *http.Request)
	RemoveCategory(w http.ResponseWriter, r *http.Request)
}

// Handler : Category handler struct
type Handler struct {
	cs ServiceInterface
}

// NewHTTPHandler : Returns category HTTP handler
func NewHTTPHandler(db *sql.DB) HandlerInterface {
	return &Handler{
		cs: NewService(db),
	}
}

// CreateCategory : to create a category
func (ch *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	log.Println("App : POST /app/category API hit!")
	var request CreateCategoryRequest
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
	category, err := ch.cs.CreateCategory(r.Context(), &request)
	if err != nil {
		log.Println("Error : ", err.Error())
		if err.Error() == utils.CategoryNameExistsError {
			utils.Fail(w, 200, utils.CategoryNameExistsErrorCode, err.Error())
			return
		}
		utils.Fail(w, 500, utils.DatabaseErrorCode, err.Error())
		return
	}
	log.Println("App : category created! id_category : ", category.ID)
	utils.Send(w, 200, category)
}

// ListCategory : to list out all categories and products
func (ch *Handler) ListCategory(w http.ResponseWriter, r *http.Request) {
	log.Println("App : GET /app/category API hit!")
	req := ListCategoryRequest{
		Limit:  DefaultListCategoryLimit,
		Offset: DefaultOffset,
	}
	request, err := validateListCategoryRequest(&req, r)
	if err != nil {
		log.Println("Error : ", err.Error())
		utils.Fail(w, 400, utils.ValidationErrorCode, err.Error())
		return
	}
	categories, err := ch.cs.ListCategory(r.Context(), request)
	if err != nil {
		log.Println("Error : ", err.Error())
		utils.Fail(w, 500, utils.DatabaseErrorCode, err.Error())
		return
	}
	log.Println("App : categories fetched!")
	utils.Send(w, 200, categories)
}

// RemoveCategory : to remove a category
func (ch *Handler) RemoveCategory(w http.ResponseWriter, r *http.Request) {
	log.Println("App : Delete /app/category/{id_category} API hit!")
	categoryID, err := strconv.Atoi(chi.URLParam(r, "id_category"))
	if err != nil {
		log.Println("Error : ", err.Error())
		utils.Fail(w, 400, utils.DecodeErrorCode, errors.New("invalid id_category").Error())
		return
	}
	request := RemoveCategoryRequest{
		CategoryID: categoryID,
	}
	err = ch.cs.RemoveCategory(r.Context(), &request)
	if err != nil {
		log.Println("Error : ", err.Error())
		if err.Error() == utils.SomeSubCategoriesAreBelongsToCategoryError {
			utils.Fail(w, 500, utils.SomeSubCategoriesAreBelongsToCategoryErrorCode, err.Error())
			return
		} else if err.Error() == utils.SomeProductsAreBelongsToCategoryError {
			utils.Fail(w, 500, utils.SomeProductsAreBelongsToCategoryErrorCode, err.Error())
			return
		} else if err.Error() == utils.InvalidCategoryIDError {
			utils.Fail(w, 500, utils.InvalidCategoryIDErrorCode, err.Error())
			return
		}
		utils.Fail(w, 500, utils.DatabaseErrorCode, err.Error())
		return
	}
	log.Println("App : category removed! id_category : ", categoryID)
	utils.Send(w, 200, nil)
}
