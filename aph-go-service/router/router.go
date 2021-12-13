package router

import (
	"aph-go-service/transport"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	//admin
	router.HandleFunc("/api/admin", transport.GETlistadminEndpoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/admin/{Admin_id}", transport.GETadminEndpoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/admin", transport.POSTadminEndpoint).Methods("POST", "OPTIONS")

	//artikel
	router.HandleFunc("/api/artikel", transport.GETlistartikelEndpoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/artikel/{Artikel_id}", transport.GETartikelEndpoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/artikel", transport.POSTartikelEndpoint).Methods("POST", "OPTIONS")

	//video
	router.HandleFunc("/api/video", transport.GETlistvideoEndpoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/video/{Video_id}", transport.GETvideoEndpoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/video", transport.POSTvideoEndpoint).Methods("POST", "OPTIONS")

	return router
}
