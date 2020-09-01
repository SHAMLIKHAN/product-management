package router

import (
	"database/sql"
	"pm/category"
	"pm/product"
	"pm/status"
	"pm/variant"

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
	ph := product.NewHTTPHandler(r.DB)
	vh := variant.NewHTTPHandler(r.DB)

	cr.Route("/v1", func(cr chi.Router) {
		cr.Get("/app/status", sh.GetAppStatus)
		cr.Post("/app/category", ch.CreateCategory)
		cr.Get("/app/category", ch.ListCategory)
		cr.Delete("/app/category/{id_category}", ch.RemoveCategory)
		cr.Post("/app/product", ph.CreateProduct)
		cr.Get("/app/product", ph.ListProduct)
		cr.Get("/app/product/{id_product}", ph.GetProduct)
		cr.Patch("/app/product/{id_product}", ph.UpdateProduct)
		cr.Delete("/app/product/{id_product}", ph.RemoveProduct)
		cr.Post("/app/product/{id_product}/variant", vh.CreateVariant)
		cr.Get("/app/product/{id_product}/variant", vh.ListVariant)
		cr.Get("/app/product/{id_product}/variant/{id_variant}", vh.GetVariant)
		cr.Patch("/app/product/{id_product}/variant/{id_variant}", vh.UpdateVariant)
		cr.Delete("/app/product/{id_product}/variant/{id_variant}", vh.RemoveVariant)
	})
	return cr
}
