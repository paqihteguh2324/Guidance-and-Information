package service

import (
	"aph-go-service/config"
	"aph-go-service/datastruct"
	"database/sql"
	fmt "fmt"
	"log"
)

func HelloWorld(name string) string {
	var helloOutput string
	helloOutput = fmt.Sprintf("Hello, %s ", name)
	return helloOutput
}

//admin parameternya ya
func POSTadmin(admin datastruct.Tbl_Admin) int64 {
	db := config.CreateConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO tbl_admin (admin_id, admin_name, admin_username, admin_password, admin_email, admin_phone) VALUES ($1, $2, $3, $4, $5,$6)`

	var Admin_id int64

	err := db.QueryRow(sqlStatement, admin.Admin_id, admin.Admin_name, admin.Admin_username, admin.Admin_password, admin.Admin_email, admin.Admin_phone).Scan(&Admin_id)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	fmt.Printf("insert data single record %v", Admin_id)

	return Admin_id
}

//mengambil semua data list admin
func GETlistadmin() ([]datastruct.Tbl_Admin, error) {
	db := config.CreateConnection()

	defer db.Close()

	var admins []datastruct.Tbl_Admin

	sqlStatement := `SELECT * FROM tbl_admin`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi Query %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var admin datastruct.Tbl_Admin
		err = rows.Scan(&admin.Admin_id, &admin.Admin_name, &admin.Admin_username, &admin.Admin_password, &admin.Admin_email, &admin.Admin_phone)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data %v", err)
		}

		admins = append(admins, admin)
	}

	return admins, err
}

//mengambil data admin persatuan
func GETadmin(Admin_id int64) (datastruct.Tbl_Admin, error) {
	db := config.CreateConnection()

	defer db.Close()

	var admin datastruct.Tbl_Admin

	sqlStatement := `SELECT * FROM tbl_admin WHERE admin_id=$1`

	row := db.QueryRow(sqlStatement, admin.Admin_id)

	err := row.Scan(&admin.Admin_id, &admin.Admin_name, &admin.Admin_username, &admin.Admin_password, &admin.Admin_email, &admin.Admin_phone)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return admin, nil
	case nil:
		return admin, nil
	default:
		log.Fatalf("tidak bisa mengambil data %v", err)
	}

	return admin, err
}

//Post Artikel
func POSTartikel(artikel datastruct.Tbl_Artikel) int64 {
	db := config.CreateConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO tbl_artikel (artikel_id2, video_id, artikel_headings, artikel_desc, artikel_image, artikel_created_date, artikel_created_by, artikel_subheadings, keyword, artikel_author) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	var artikel_id2 int64

	err := db.QueryRow(sqlStatement, artikel.Artikel_id2, artikel.Video_id, artikel.Artikel_headings, artikel.Artikel_desc, artikel.Artikel_image, artikel.Artikel_created_date, artikel.Artikel_created_by, artikel.Artikel_update_date, artikel.Artikel_update_by, artikel.Artikel_subheadings, artikel.Keyword, artikel.Artikel_author).Scan(&artikel_id2)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	fmt.Printf("insert data single record %v", artikel_id2)

	return artikel_id2
}

//mengambil semua data list admin
func GETlistartikel() ([]datastruct.Tbl_Artikel, error) {
	db := config.CreateConnection()

	defer db.Close()

	var list_artikel []datastruct.Tbl_Artikel

	sqlStatement := `SELECT * FROM tbl_artikel`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi Query %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var artikel datastruct.Tbl_Artikel
		err = rows.Scan(&artikel.Artikel_id2, &artikel.Video_id, &artikel.Artikel_headings, &artikel.Artikel_desc, &artikel.Artikel_image, &artikel.Artikel_created_date, &artikel.Artikel_created_by, &artikel.Artikel_update_date, &artikel.Artikel_update_by, &artikel.Artikel_subheadings, &artikel.Keyword, &artikel.Artikel_author)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data %v", err)
		}

		list_artikel = append(list_artikel, artikel)
	}

	return list_artikel, err
}

//mengambil data artikel per-id
func GETartikel(Artikel_id2 int64) (datastruct.Tbl_Artikel, error) {
	db := config.CreateConnection()

	defer db.Close()

	var artikel datastruct.Tbl_Artikel

	sqlStatement := `SELECT * FROM tbl_admin WHERE artikel_id2=$1`

	row := db.QueryRow(sqlStatement, artikel.Artikel_id2)

	err := row.Scan(&artikel.Artikel_id2, &artikel.Video_id, &artikel.Artikel_headings, &artikel.Artikel_desc, &artikel.Artikel_image, &artikel.Artikel_created_date, &artikel.Artikel_created_by, &artikel.Artikel_update_date, &artikel.Artikel_update_by, &artikel.Artikel_subheadings, &artikel.Keyword, &artikel.Artikel_author)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return artikel, nil
	case nil:
		return artikel, nil
	default:
		log.Fatalf("tidak bisa mengambil data %v", err)
	}

	return artikel, err
}

//Post Artikel
func POSTvideo(video datastruct.Tbl_Video) int64 {
	db := config.CreateConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO tbl_artikel (video_id, artikel_id2, video_headings, video_desc, video_link, video_created_date, video_created_by, video_update_by, video_update_date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	var video_id int64

	err := db.QueryRow(sqlStatement, video.Video_id, video.Artikel_id2, video.Video_headings, video.Video_desc, video.Video_link, video.Video_created_date, video.Video_craeted_by, video.Video_update_by, video.Video_update_date).Scan(&video_id)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	fmt.Printf("insert data single record %v", video_id)

	return video_id
}

//mengambil semua data list admin
func GETlistvideo() ([]datastruct.Tbl_Video, error) {
	db := config.CreateConnection()

	defer db.Close()

	var list_video []datastruct.Tbl_Video

	sqlStatement := `SELECT * FROM tbl_video`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi Query %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var video datastruct.Tbl_Video
		err = rows.Scan(&video.Video_id, &video.Artikel_id2, &video.Video_headings, &video.Video_desc, &video.Video_link, &video.Video_created_date, &video.Video_craeted_by, &video.Video_update_by, &video.Video_update_date)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data %v", err)
		}

		list_video = append(list_video, video)
	}

	return list_video, err
}

//mengambil data artikel per-id
func GETvideo(Video_id int64) (datastruct.Tbl_Video, error) {
	db := config.CreateConnection()

	defer db.Close()

	var video datastruct.Tbl_Video

	sqlStatement := `SELECT * FROM tbl_admin WHERE video_id=$1`

	row := db.QueryRow(sqlStatement, video.Video_id)

	err := row.Scan(&video.Video_id, &video.Artikel_id2, &video.Video_headings, &video.Video_desc, &video.Video_link, &video.Video_created_date, &video.Video_craeted_by, &video.Video_update_by, &video.Video_update_date)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return video, nil
	case nil:
		return video, nil
	default:
		log.Fatalf("tidak bisa mengambil data %v", err)
	}

	return video, err
}
