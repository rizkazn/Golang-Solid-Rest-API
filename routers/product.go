package routers

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rizkazn/gorestfull/controllers"
	"github.com/rizkazn/gorestfull/middleware"
	"github.com/rizkazn/gorestfull/repos"
	"github.com/rizkazn/gorestfull/services"
)

func ProductRoute(r *mux.Router, db *sql.DB) {

	prodrepo := repos.Products(db)
	svc := services.ProductService(prodrepo)
	cr := controllers.Products(svc)

	// Prefix and Subrouter
	route := r.PathPrefix("/products").Subrouter()
	// Route Handlers / Endpoints
	route.HandleFunc("", cr.GetAll).Methods(http.MethodGet)
	route.HandleFunc("/", cr.GetAll).Methods(http.MethodGet)
	route.HandleFunc("/", middleware.Validate(cr.Add)).Methods(http.MethodPost)
	route.HandleFunc("/{id}", middleware.Validate(cr.Update)).Methods(http.MethodPut)
	route.HandleFunc("/{id}", middleware.Validate(cr.Delete)).Methods(http.MethodDelete)
	route.HandleFunc("/order/name", cr.SortByName).Methods(http.MethodGet)
	route.HandleFunc("/order/catId", cr.SortByCat).Methods(http.MethodGet)
	route.HandleFunc("/order/time", cr.SortByNewest).Methods(http.MethodGet)
	route.HandleFunc("/order/price", cr.SortByPrice).Methods(http.MethodGet)
	route.HandleFunc("/search/name", cr.SearchByName).Methods(http.MethodGet)
}
