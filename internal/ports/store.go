package ports

import (
	"time"

	"github.com/Odvin/go-mock-http-server/internal/application/domain"
)

type Store interface {
	GetCompany(id int64) (*domain.Company, error)
	GetCompanyUpdates(from, to time.Time, status string) []domain.Company
	StopCompanyUpdates()
	StartCompanyUpdates(period int64)
	GetCompanyInfo() *domain.CompanyInfo
}
