package product

import (
	"errors"
	"net/http"
	"strconv"
)

func validateListProductRequest(request *ListProductRequest, r *http.Request) (*ListProductRequest, error) {
	limit := r.URL.Query().Get("limit")
	if limit != "" {
		productLimit, err := strconv.Atoi(limit)
		if err != nil {
			return nil, errors.New("invalid query parameter, limit")
		}
		request.Limit = productLimit
	}
	page := r.URL.Query().Get("page")
	if page != "" {
		productPage, err := strconv.Atoi(page)
		if err != nil {
			return nil, errors.New("invalid query parameter, page")
		}
		request.Offset = (productPage - 1) * request.Limit
	}
	return request, nil
}
