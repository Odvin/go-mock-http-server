package store

import (
	"time"

	"github.com/Odvin/go-mock-http-server/internal/application/domain"
)

type StoreAdapter struct {
	maxElements    int
	company        []domain.Company
	companyUpdater *companyUpdater
}

func Init(maxElements int) *StoreAdapter {
	company := make([]domain.Company, maxElements)
	seedCompany(company)

	return &StoreAdapter{
		maxElements:    maxElements,
		company:        company,
		companyUpdater: NewCompanyUpdater(30*time.Second, company),
	}
}
