package store

import (
	"time"

	"github.com/Odvin/go-mock-http-server/internal/app"
	"github.com/Odvin/go-mock-http-server/pkg/mediator"
)

var ps = mediator.GetPubSub()

type Store struct {
	companies      []app.Company
	companyUpdater *CompanyUpdater
}

func Init(n int) *Store {
	companies := make([]app.Company, n)
	seedCompany(companies)

	return &Store{
		companies:      companies,
		companyUpdater: NewCompanyUpdater(10*time.Second, companies),
	}
}
