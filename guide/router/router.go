package router

import (
	"guide/transport"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/admin", transport.GETlistadminEndpoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/admin/{Admin_id}", transport.GETadminEndpoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/admin", transport.POSTadminEndpoint).Methods("POST", "OPTIONS")

	return router
}
