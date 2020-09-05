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

	cr.Route("/v1/app", func(cr chi.Router) {
		cr.Get("/status", sh.GetAppStatus)
		cr.Post("/category", ch.CreateCategory)
		cr.Get("/category", ch.ListCategory)
		cr.Patch("/category/{id_category}", ch.UpdateCategory)
		cr.Delete("/category/{id_category}", ch.RemoveCategory)
		cr.Post("/product", ph.CreateProduct)
		cr.Get("/product", ph.ListProduct)
		cr.Get("/product/{id_product}", ph.GetProduct)
		cr.Patch("/product/{id_product}", ph.UpdateProduct)
		cr.Delete("/product/{id_product}", ph.RemoveProduct)
		cr.Post("/product/{id_product}/variant", vh.CreateVariant)
		cr.Get("/product/{id_product}/variant", vh.ListVariant)
		cr.Get("/product/{id_product}/variant/{id_variant}", vh.GetVariant)
		cr.Patch("/product/{id_product}/variant/{id_variant}", vh.UpdateVariant)
		cr.Delete("/product/{id_product}/variant/{id_variant}", vh.RemoveVariant)
	})
	return cr
}
