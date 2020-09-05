package category

import (
	"errors"
	"net/http"
	"strconv"
)

func validateListCategoryRequest(request *ListCategoryRequest, r *http.Request) (*ListCategoryRequest, error) {
	limit := r.URL.Query().Get("limit")
	if limit != "" {
		categoryLimit, err := strconv.Atoi(limit)
		if err != nil {
			return nil, errors.New("invalid query parameter, limit")
		}
		request.Limit = categoryLimit
	}
	page := r.URL.Query().Get("page")
	if page != "" {
		categoryPage, err := strconv.Atoi(page)
		if err != nil {
			return nil, errors.New("invalid query parameter, page")
		}
		request.Offset = (categoryPage - 1) * request.Limit
	}
	return request, nil
}
