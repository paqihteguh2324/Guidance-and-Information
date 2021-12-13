package main

import (
	"aph-go-service/router"
	//"aph-go-service/transport"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Server dijalankan pada port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

	/*
		http.HandleFunc("/admin", transport.GETlistadminEndpoint)
		http.HandleFunc("/getadmin/{admin_id}", transport.GETadminEndpoint)
		http.HandleFunc("/postadmin", transport.POSTadminEndpoint)

		http.HandleFunc("/artikel", transport.GETlistartikelEndpoint)
		//http.HandleFunc("/artikel/{Artikel_id}", transport.GETartikelEndpoint)
		http.HandleFunc("/postartikel", transport.POSTartikelEndpoint)

		http.HandleFunc("/video", transport.GETlistvideoEndpoint)
		//http.HandleFunc("/video/{Video_id}", transport.GETvideoEndpoint)
		http.HandleFunc("/postvideo", transport.POSTvideoEndpoint)
		//http.HandleFunc("/insert", POSTartikel)
		fmt.Println("Server dijalankan pada port 8080...")
		log.Fatal(http.ListenAndServe(":8080", nil))
	*/
}
