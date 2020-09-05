package status

import (
	"database/sql"
	"log"
	"net/http"
	"pm/utils"
)

// HandlerInterface : Status handler
type HandlerInterface interface {
	GetAppStatus(w http.ResponseWriter, r *http.Request)
}

// Handler : Status handler struct
type Handler struct{}

// NewHTTPHandler : Returns statys HTTP handler
func NewHTTPHandler(db *sql.DB) HandlerInterface {
	return &Handler{}
}

// GetAppStatus : to get app status
func (ah *Handler) GetAppStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("App : GET /v1/app/status API hit!")
	status := Status{
		Database: "running...",
		Server:   "running...",
	}
	utils.Send(w, 200, status)
}
