package router

import (
	"database/sql"
	"pm/category"
	"pm/status"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// Router : Basic router
type Router interface {
	Setup() *chi.Mux
}

// ChiRouter : Router that holds database connection
type ChiRouter struct {
	DB *sql.DB
}

// NewRouter : Returns basic router
func NewRouter(db *sql.DB) Router {
	return &ChiRouter{
		DB: db,
	}
}

// Setup : chi router
func (r *ChiRouter) Setup() *chi.Mux {
	cr := chi.NewRouter()
	cr.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type"},
	}))

	sh := status.NewHTTPHandler(r.DB)
	ch := category.NewHTTPHandler(r.DB)

	cr.Route("/v1", func(cr chi.Router) {
		cr.Get("/app/status", sh.GetAppStatus)
		cr.Post("/app/category", ch.CreateCategory)
	})
	return cr
}
