package application

import "github.com/Odvin/go-mock-http-server/internal/application/domain"

func (app *Application) GetCompany(id int) (*domain.Company, error) {
	company, err := app.store.GetCompany(id)
	if err != nil {
		return nil, err
	}

	return company, nil
}
