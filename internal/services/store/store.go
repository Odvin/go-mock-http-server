package store

import (
	"time"

	"github.com/Odvin/go-mock-http-server/internal/app"
)

type Store struct {
	maxElements    int
	company        []app.Company
	companyUpdater *CompanyUpdater
}

func Init(maxElements int) *Store {
	company := make([]app.Company, maxElements)
	seedCompany(company)

	return &Store{
		maxElements:    maxElements,
		company:        company,
		companyUpdater: NewCompanyUpdater(10*time.Second, company),
	}
}
