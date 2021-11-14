package main

import (
	"fmt"
	//"guide/router"

	"guide/transport"

	"log"
	"net/http" // digunakan untuk mengakses objek permintaan dan respons dari api
)

func main() {
	//r := router.Router()
	http.HandleFunc("/admin", transport.GETlistadminEndpoint)
	http.HandleFunc("/admin/{admin_id}", transport.GETadminEndpoint)
	http.HandleFunc("/postadmin", transport.POSTadminEndpoint)
	//http.HandleFunc("/insert", POSTartikel)
	fmt.Println("Server dijalankan pada port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
