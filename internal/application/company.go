package application

import (
	"time"

	"github.com/Odvin/go-mock-http-server/internal/application/domain"
)

func (app *Application) GetCompany(id int) (*domain.Company, error) {
	company, err := app.store.GetCompany(id)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (app *Application) GetCompanyUpdates(from, to time.Time, status string) []domain.Company {
	return app.store.GetCompanyUpdates(from, to, status)
}
