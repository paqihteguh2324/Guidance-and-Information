package datastruct

import (
	"context"
	"time"
)

type HelloWorldRequest struct {
	NAME string `json:"name"`
}

type HelloWorldResponse struct {
	MESSAGE string `json:"message"`
}

type userInformation struct {
	User_ID      int
	Username     string
	Email        string
	Firstname    string
	Lastname     string
	Phonenumber  string
	Password     string
	Created_Date time.Time
	Created_By   string
	Updated_Date time.Time
	Updated_By   string
}

type (
	// Guidelines view from sdx_sertifikasi_db
	Guidelines struct {
		ID                    int    `json:"id"`
		GuidelinesName        string `json:"guidelines_name"`
		GuidelinesDescription string `json:"guidelines_description"`
		GuidelinesType        string `json:"guidelines_type"`
		GuidelinesLink        string `json:"guidelines_link"`
	}

	// GuidelinesRepository DB action from Guidelines table
	GuidelinesRepository interface {
		GetGuidelinesDocument(ctx context.Context) ([]Guidelines, error)
	}
)
