package datastruct

type HelloWorldRequest struct {
	NAME string `json:"name"`
}

type HelloWorldResponse struct {
	MESSAGE string `json:"message"`
}

//penulisan huruf awal harus besar agar menjadi publik (exported)
//penulisan huruf kecil berarti private (unexported)
type Tbl_Admin struct {
	Admin_id       int64  `json:"admin_id"`
	Admin_name     string `json:"admin_name"`
	Admin_username string `json:"admin_username"`
	Admin_password string `json:"admin_password"`
	Admin_email    string `json:"admin_email"`
	Admin_phone    string `json:"admin_phone"`
}

type Tbl_User struct {
	Username          string `json:"username"`
	Name              string `json:"name"`
	Phonenumber       string `json:"phonenumber"`
	Password          string `json:"password"`
	Email_verified    string `json:"email_verified"`
	Image_file        string `json:"image_file"`
	Identity_type     string `json:"identity_type"`
	Identity_no       string `json:"identity_no"`
	Address_ktp       string `json:"address_ktp"`
	Domisili          string `json:"domisili"`
	Video_create_date string `json:"create_date"`
	Update_date       string `json:"update_date"`
	Email             string `json:"email"`
	User_id           int64  `json:"user_id"`
}

type Tbl_Artikel struct {
	Artikel_id2          int64  `json:"artikel_id2"`
	Video_id             int64  `json:"video_id"`
	Artikel_headings     string `json:"artikel_headings"`
	Artikel_desc         string `json:"artikel_desc"`
	Artikel_image        string `json:"artikel_image"`
	Artikel_created_date string `json:"artikel_created_date"`
	Artikel_created_by   string `json:"artikel_created_by"`
	Artikel_update_date  string `json:"artikel_update_date"`
	Artikel_update_by    string `json:"artikel_update_by"`
	Artikel_subheadings  string `json:"artikel_subheadings"`
	Keyword              string `json:"keyword"`
	Artikel_author       string `json:"Artikel_author"`
}

type Tbl_Video struct {
	Video_id           int64  `json:"video_id"`
	Artikel_id2        int64  `json:"artikel_id"`
	Video_headings     string `json:"video_headings"`
	Video_desc         string `json:"video_desc"`
	Video_link         string `json:"video_link"`
	Video_created_date string `json:"video_created_date"`
	Video_craeted_by   string `json:"video_craeted_by"`
	Video_update_date  string `json:"video_update_date"`
	Video_update_by    string `json:"video_update_by"`
}

type Response1 struct {
	Status   string `json:"status"`
	Message  string `json:"message,omitempty"`
	Admin_id int64  `json:"admin_id,omitempty"`
}

type Response2 struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    []Tbl_Admin `json:"data"`
}

type Response3 struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []Tbl_Artikel `json:"data"`
}

type Response4 struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    []Tbl_Video `json:"data"`
}

type Response5 struct {
	Status      string `json:"status"`
	Message     string `json:"message,omitempty"`
	Artikel_id2 int64  `json:"artikel2_id,omitempty"`
}

type Response6 struct {
	Status   string `json:"status"`
	Message  string `json:"message,omitempty"`
	Video_id int64  `json:"video_id,omitempty"`
}
