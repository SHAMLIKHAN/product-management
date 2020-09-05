package variant

import (
	"errors"
	"net/http"
	"strconv"
)

func validateListVariantRequest(request *ListVariantRequest, r *http.Request) (*ListVariantRequest, error) {
	limit := r.URL.Query().Get("limit")
	if limit != "" {
		variantLimit, err := strconv.Atoi(limit)
		if err != nil {
			return nil, errors.New("invalid query parameter, limit")
		}
		request.Limit = variantLimit
	}
	page := r.URL.Query().Get("page")
	if page != "" {
		variantPage, err := strconv.Atoi(page)
		if err != nil {
			return nil, errors.New("invalid query parameter, page")
		}
		request.Offset = (variantPage - 1) * request.Limit
	}
	return request, nil
}
