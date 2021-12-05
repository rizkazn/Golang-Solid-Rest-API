package routers

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rizkazn/gorestfull/controllers"
	"github.com/rizkazn/gorestfull/repos"
	"github.com/rizkazn/gorestfull/services"
)

func UserRoute(r *mux.Router, db *sql.DB) {

	userrepo := repos.User(db)
	svc := services.UserService(userrepo)
	cr := controllers.User(svc)

	// Prefix and Subrouter
	route := r.PathPrefix("/users").Subrouter()
	// Route Handlers / Endpoints
	route.HandleFunc("", cr.GetAll).Methods(http.MethodGet)
	route.HandleFunc("/", cr.GetAll).Methods(http.MethodGet)
	route.HandleFunc("/", cr.Add).Methods(http.MethodPost)
}
