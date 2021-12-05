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

func CategoryRoute(r *mux.Router, db *sql.DB) {

	catrepo := repos.Category(db)
	svc := services.CategoryService(catrepo)
	cr := controllers.Category(svc)

	// Prefix and Subrouter
	route := r.PathPrefix("/categories").Subrouter()
	// Route Handlers / Endpoints
	route.HandleFunc("", cr.GetAll).Methods(http.MethodGet)
	route.HandleFunc("/", cr.GetAll).Methods(http.MethodGet)
	route.HandleFunc("/", middleware.Validate(cr.Add)).Methods(http.MethodPost)
	route.HandleFunc("/{id}", middleware.Validate(cr.Update)).Methods(http.MethodPut)
	route.HandleFunc("/{id}", middleware.Validate(cr.Delete)).Methods(http.MethodDelete)
}
