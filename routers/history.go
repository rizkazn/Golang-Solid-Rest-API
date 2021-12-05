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

func HistoryRoute(r *mux.Router, db *sql.DB) {

	histrepo := repos.History(db)
	svc := services.HistoryService(histrepo)
	cr := controllers.History(svc)

	// Prefix and Subrouter
	route := r.PathPrefix("/histories").Subrouter()
	// Route Handlers / Endpoints
	route.HandleFunc("", cr.GetAll).Methods(http.MethodGet)
	route.HandleFunc("/", cr.GetAll).Methods(http.MethodGet)
	route.HandleFunc("/", middleware.Validate(cr.Add)).Methods(http.MethodPost)
	route.HandleFunc("/{id}", middleware.Validate(cr.Update)).Methods(http.MethodPut)
	route.HandleFunc("/{id}", middleware.Validate(cr.Delete)).Methods(http.MethodDelete)
}
