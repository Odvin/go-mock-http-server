package app

import (
	"errors"
	"time"
)

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

func (a *Application) GetCompany(id int64) (*Company, error) {
	if id <= 0 {
		return nil, errors.New("company id must be positive")
	}

	company, err := a.store.GetCompany(id)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (a *Application) GetCompanyUpdates(from, to time.Time, status string, page, size int) ([]Company, int) {
	return a.store.GetCompanyUpdates(from, to, status, page, size)
}

func (a *Application) StopCompanyUpdates() {
	a.store.StopCompanyUpdates()
}

func (a *Application) StartCompanyUpdates(period int64) error {
	if period < 1 || period > 3600 {
		return errors.New("period to update companies must be in [1; 3600] seconds")
	}

	a.store.StartCompanyUpdates(period)

	return nil
}

func (a *Application) GetCompanyInfo() *CompanyInfo {
	return a.store.GetCompanyInfo()
}
