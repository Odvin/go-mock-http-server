package application

import (
	"time"

	"github.com/Odvin/go-mock-http-server/internal/application/domain"
	"github.com/Odvin/go-mock-http-server/internal/ports"
)

type API interface {
	GetCompany(id int64) (*domain.Company, error)
	GetCompanyUpdates(from, to time.Time, status string) []domain.Company
	StopCompanyUpdates()
	StartCompanyUpdates(period int64) error
	GetCompanyInfo() *domain.CompanyInfo
}

type Application struct {
	store ports.Store
}

func Init(store ports.Store) *Application {
	return &Application{
		store: store,
	}
}
