package transport

import (
	"encoding/json"
	"strconv"

	"guide/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type respons struct {
	Admin_id string `json:"admin_id,omitempty"`
	Message  string `json:"message,omitempty"`
}

type Respons struct {
	Status  int                `json:"status"`
	Message string             `jsong:"message"`
	Data    []server.Tbl_Admin `json:"data"`
}

func POSTadminEndpoint(W http.ResponseWriter, r *http.Request) {
	var admin server.Tbl_Admin

	err := json.NewDecoder(r.Body).Decode(&admin)

	if err != nil {
		log.Fatalf("tidak bisa mendecode dari request body. %v", err)
	}

	insertAdmin_id := server.POSTadmin(admin)

	res := respons{
		Admin_id: insertAdmin_id,
		Message:  "Data buku telah ditambahkan",
	}

	json.NewEncoder(W).Encode(res)
}

func GETadminEndpoint(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Context-Type", "application/x-WWW-form-urlencoded")
	W.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	Admin_id, err := strconv.Atoi(params["Admin_id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int. %v", err)
	}

	admin, err := server.GETadmin(string(Admin_id))

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data admin. %v", err)
	}

	json.NewEncoder(W).Encode(admin)
}

func GETlistadminEndpoint(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Context-Type", "application/x-WWW-form-urlencoded")
	W.Header().Set("Access-Control-Allow-Origin", "*")

	admins, err := server.GETlistadmin()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	var respons Respons
	respons.Status = 1
	respons.Message = "Success"
	respons.Data = admins

	json.NewEncoder(W).Encode(respons)
}
