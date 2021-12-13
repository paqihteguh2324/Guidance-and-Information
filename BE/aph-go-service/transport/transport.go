package transport

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"aph-go-service/datastruct"
	"aph-go-service/logging"
	"aph-go-service/service"

	"github.com/gorilla/mux"
)

//admin
func POSTadminEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var admin datastruct.Tbl_Admin

	err := json.NewDecoder(r.Body).Decode(&admin)

	if err != nil {
		log.Fatalf("tidak bisa mendecode dari request body. %v", err)
	}

	insertAdmin_id := service.POSTadmin(admin)

	logging.Log(fmt.Sprintf("menambahkan admin %d", admin.Admin_id))

	res := datastruct.Response1{
		Status:   "Berhasil",
		Admin_id: insertAdmin_id,
		Message:  "Data admin telah ditambahkan",
	}

	json.NewEncoder(w).Encode(res)
}

func GETlistadminEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-WWW-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	admins, err := service.GETlistadmin()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	logging.Log(fmt.Sprintf("menampilkan seluruh admin"))

	res := datastruct.Response2{
		Status:  "Berhasil",
		Message: "Data keseluruhan admin",
		Data:    admins,
	}

	json.NewEncoder(w).Encode(res)
}

func GETadminEndpoint(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Context-Type", "application/x-WWW-form-urlencoded")
	W.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	Admin_id, err := strconv.Atoi(params["Admin_id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int. %v", err)
	}

	admin, err := service.GETadmin(int64(Admin_id))

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data admin. %v", err)
	}

	logging.Log(fmt.Sprintf("menampilkan admin %d", admin.Admin_id))

	res := datastruct.Response1{
		Status:   "Berhasil",
		Admin_id: admin.Admin_id,
		Message:  "Data admin",
	}

	json.NewEncoder(W).Encode(res)
}

//artikel
func POSTartikelEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var artikel datastruct.Tbl_Artikel

	err := json.NewDecoder(r.Body).Decode(&artikel)

	if err != nil {
		log.Fatalf("tidak bisa mendecode dari request body. %v", err)
	}

	insertArtikel_id := service.POSTartikel(artikel)

	logging.Log(fmt.Sprintf("menambahkan admin %d", artikel.Artikel_id2))

	res := datastruct.Response5{
		Status:      "Berhasil",
		Artikel_id2: insertArtikel_id,
		Message:     "Data admin telah ditambahkan",
	}

	json.NewEncoder(w).Encode(res)
}

func GETlistartikelEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-WWW-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	list_artikel, err := service.GETlistartikel()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	logging.Log(fmt.Sprintf("menampilkan seluruh admin"))

	res := datastruct.Response3{
		Status:  "Berhasil",
		Message: "Data keseluruhan artikel",
		Data:    list_artikel,
	}

	json.NewEncoder(w).Encode(res)
}

func GETartikelEndpoint(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Context-Type", "application/x-WWW-form-urlencoded")
	W.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	Artikel_id, err := strconv.Atoi(params["Artikel_id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int. %v", err)
	}

	artikel, err := service.GETartikel(int64(Artikel_id))

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data admin. %v", err)
	}

	logging.Log(fmt.Sprintf("menampilkan admin %d", artikel.Artikel_id2))

	res := datastruct.Response5{
		Status:      "Berhasil",
		Message:     "Data artikel telah ditemukan",
		Artikel_id2: artikel.Artikel_id2,
	}

	json.NewEncoder(W).Encode(res)
}

//Video
func POSTvideoEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var video datastruct.Tbl_Video

	err := json.NewDecoder(r.Body).Decode(&video)

	if err != nil {
		log.Fatalf("tidak bisa mendecode dari request body. %v", err)
	}

	insertVideo_id := service.POSTvideo(video)

	logging.Log(fmt.Sprintf("menambahkan admin %d", video.Video_id))

	res := datastruct.Response6{
		Status:   "Berhasil",
		Message:  "Data video telah ditambahkan",
		Video_id: insertVideo_id,
	}

	json.NewEncoder(w).Encode(res)
}

func GETlistvideoEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-WWW-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	list_video, err := service.GETlistvideo()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	logging.Log(fmt.Sprintf("menampilkan seluruh admin"))

	res := datastruct.Response4{
		Status:  "Berhasil",
		Message: "Data keseluruhan video",
		Data:    list_video,
	}

	json.NewEncoder(w).Encode(res)
}

func GETvideoEndpoint(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Context-Type", "application/x-WWW-form-urlencoded")
	W.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	Video_id, err := strconv.Atoi(params["Video_id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int. %v", err)
	}

	video, err := service.GETvideo(int64(Video_id))

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data admin. %v", err)
	}

	logging.Log(fmt.Sprintf("menampilkan admin %d", video.Video_id))

	res := datastruct.Response6{
		Status:   "Berhasil",
		Video_id: video.Video_id,
		Message:  "Data admin",
	}

	json.NewEncoder(W).Encode(res)
}
