package store

import "github.com/Odvin/go-mock-http-server/internal/application/domain"

type StoreAdapter struct {
	maxElements int
	company     []domain.Company
}

func Init(maxElements int) *StoreAdapter {
	company := make([]domain.Company, maxElements)
	seedCompany(company)

	return &StoreAdapter{
		maxElements: maxElements,
		company:     company,
	}
}
