package application

import (
	"errors"
	"time"

	"github.com/Odvin/go-mock-http-server/internal/application/domain"
)

func (app *Application) GetCompany(id int64) (*domain.Company, error) {
	if id <= 0 {
		return nil, errors.New("company id must be positive")
	}

	company, err := app.store.GetCompany(id)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (app *Application) GetCompanyUpdates(from, to time.Time, status string, page, size int) ([]domain.Company, int) {
	return app.store.GetCompanyUpdates(from, to, status, page, size)
}

func (app *Application) StopCompanyUpdates() {
	app.store.StopCompanyUpdates()
}

func (app *Application) StartCompanyUpdates(period int64) error {
	if period < 1 || period > 3600 {
		return errors.New("period to update companies must be in [1; 3600] seconds")
	}

	app.store.StartCompanyUpdates(period)

	return nil
}

func (app *Application) GetCompanyInfo() *domain.CompanyInfo {
	return app.store.GetCompanyInfo()
}
