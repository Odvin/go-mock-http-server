package ports

import (
	"time"

	"github.com/Odvin/go-mock-http-server/internal/application/domain"
)

type Store interface {
	GetCompany(id int64) (*domain.Company, error)
	GetCompanyUpdates(from, to time.Time, status string, page, size int) ([]domain.Company, int)
	StopCompanyUpdates()
	StartCompanyUpdates(period int64)
	GetCompanyInfo() *domain.CompanyInfo
}
