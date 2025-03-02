package ports

import (
	"time"

	"github.com/Odvin/go-mock-http-server/internal/application/domain"
)

type Store interface {
	GetCompany(id int) (*domain.Company, error)
	GetCompanyUpdates(from, to time.Time, status string) []domain.Company
}
