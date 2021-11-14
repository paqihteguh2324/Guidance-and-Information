package server

import (
	"database/sql"
	"fmt"
	"guide/config"
	"log"

	_ "github.com/lib/pq"
)

//penulisan huruf awal harus besar agar menjadi publik (exported)
//penulisan huruf kecil berarti private (unexported)
type Tbl_Admin struct {
	Admin_id       string `json:"admin_id"`
	Admin_name     string `json:"admin_name"`
	Admin_username string `json:"admin_username"`
	Admin_password string `json:"admin_password"`
	Admin_email    string `json:"admin_email"`
	Admin_phone    string `json:"admin_phone"`
}

//admin parameternya ya
func POSTadmin(admin Tbl_Admin) string {
	db := config.CreateConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO tbl_admin (admin_id, admin_name, admin_username, admin_password, admin_email, admin_phone) VALUES ($1, $2, $3, $4, $5,$6)`

	var Admin_id string

	err := db.QueryRow(sqlStatement, admin.Admin_id, admin.Admin_name, admin.Admin_username, admin.Admin_password, admin.Admin_email, admin.Admin_phone).Scan(&Admin_id)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	fmt.Printf("insert data single record %v", Admin_id)

	return Admin_id
}

func GETlistadmin() ([]Tbl_Admin, error) {
	db := config.CreateConnection()

	defer db.Close()

	var admins []Tbl_Admin

	sqlStatement := `SELECT * FROM tbl_admin`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi Query %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var admin Tbl_Admin
		err = rows.Scan(&admin.Admin_id, &admin.Admin_name, &admin.Admin_username, &admin.Admin_password, &admin.Admin_email, &admin.Admin_phone)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data %v", err)
		}

		admins = append(admins, admin)
	}

	return admins, err
}

func GETadmin(Admin_id string) (Tbl_Admin, error) {
	db := config.CreateConnection()

	defer db.Close()

	var admin Tbl_Admin

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
