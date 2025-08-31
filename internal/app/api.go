package app

import (
	"time"
)

type API interface {
	GetCompany(id int64) (*Company, error)
	GetCompanyUpdates(from, to time.Time, status string, page, size int) ([]Company, int)
	StopCompanyUpdates()
	StartCompanyUpdates(period int64) error
	GetCompanyInfo() *CompanyInfo
}
