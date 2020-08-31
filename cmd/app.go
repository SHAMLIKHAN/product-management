package cmd

import (
	"database/sql"
	"log"
	"net/http"
	"pm/router"
)

// App : Struct to represent this app
type App struct {
	router.Router
}

// NewApp : to get App Struct
func NewApp(db *sql.DB) *App {
	return &App{
		Router: router.NewRouter(db),
	}
}

// Serve : to Run API Server
func (a *App) Serve(addr string) {
	router := a.Router.Setup()
	log.Println("App : Server is listening...")
	http.ListenAndServe(addr, router)
}
