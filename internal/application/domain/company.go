package domain

import "time"

type Company struct {
	ID      int       `json:"id"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Active  bool      `json:"active"`
	Company string    `json:"company"`
	Status  string    `json:"status"`
	Phone   string    `json:"phone"`
	Email   string    `json:"email"`
	Staff   int       `json:"staff"`
}

type CompanyInfo struct {
	Total    int  `json:"total"`
	Updating bool `json:"updating"`
	Period   int  `json:"period"`
}
