package datastruct

// type mengolah struct {
// 	admin_id   string `json:"admin_id"`
// 	video_id string `json:"video_id"`
// 	status_video string `json:"status_video"`
// }

// type mengelola struct {
// 	admin_id   string `json:"admin_id"`
// 	artikel_id string `json:"artikel_id"`
// 	status_artikel string `json:"status_artikel"`
// }

type member struct {
	member_id   string `json:"member_id"`
	member_name string `json:"member_name"`
}

type admin struct {
	admin_id       string `json:"admin_id"`
	admin_name     string `json:"admin_name"`
	admin_username string `json:"admin_username"`
	admin_password string `json:"admin_password"`
	admin_email    string `json:"admin_email"`
	admin_phone    string `json:"admin_phone"`
}

type artikel struct {
	artikel_id    string `json:"artikel_data"`
	video_id      string `json:"video_id"`
	artikel_name  string `json:"artikel_name"`
	artikel_desc  string `json:"artikel_desc"`
	artikel_image string `json:"artikel_image"`
	created_date  string `json:"created_date"`
	created_by    string `json:"created_by"`
	update_date   string `json:"update_date"`
	update_by     string `json:"update_by"`
}

type video struct {
	video_id     string `json:"video_id"`
	artikel_id   string `json:"artikel_id"`
	video_name   string `json:"video_name"`
	video_desc   string `json:"video_desc"`
	video_link   string `json:"video_link"`
	created_date string `json:"created_date"`
	created_by   string `json:"created_by"`
	update_date  string `json:"update_date"`
	update_by    string `json:"update_by"`
}

type service interface {
}
