package app

import "time"

type Store interface {
	GetCompany(id int64) (*Company, error)
	GetCompanyUpdates(from, to time.Time, status string, page, size int) ([]Company, int)
	StopCompanyUpdates()
	StartCompanyUpdates(period int64)
	GetCompanyInfo() *CompanyInfo
}
