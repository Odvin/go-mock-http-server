package store

import (
	"time"

	"github.com/Odvin/go-mock-http-server/internal/application/domain"
)

type Store struct {
	maxElements    int
	company        []domain.Company
	companyUpdater *companyUpdater
}

func Init(maxElements int) *Store {
	company := make([]domain.Company, maxElements)
	seedCompany(company)

	return &Store{
		maxElements:    maxElements,
		company:        company,
		companyUpdater: NewCompanyUpdater(10*time.Second, company),
	}
}
