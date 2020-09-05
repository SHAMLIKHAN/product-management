package utils

import (
	"encoding/json"
	"net/http"
)

const (
	// StatusOk : All good
	StatusOk = "ok"
	// StatusNOk : Bad response
	StatusNOk = "nok"
)

// CustomError : The error format in api response
type CustomError struct {
	Code    int    `json:"code"`
	Details string `json:"details"`
}

// Response : The api response format
type Response struct {
	Status string       `json:"status"`
	Error  *CustomError `json:"error,omitempty"`
	Result *interface{} `json:"result,omitempty"`
}

// Send : General function to send api response
func Send(w http.ResponseWriter, status int, payload interface{}) {
	response := Response{
		Status: StatusOk,
		Result: &payload,
	}
	result, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(result)
}

// Fail : General function to send api response
func Fail(w http.ResponseWriter, status, code int, details string) {
	response := &Response{
		Status: StatusNOk,
		Error: &CustomError{
			Code:    code,
			Details: details,
		},
	}
	result, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(result)
}
