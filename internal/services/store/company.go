package store

import (
	"errors"
	"time"

	"github.com/Odvin/go-mock-http-server/internal/application/domain"
	"github.com/brianvoe/gofakeit/v7"
)

func seedCompany(company []domain.Company) {
	statuses := []string{"public", "private"}

	var created time.Time
	for i := range len(company) {
		created = gofakeit.Date()

		company[i] = domain.Company{
			ID:      i + 1,
			Created: created,
			Updated: created,
			Active:  gofakeit.Bool(),
			Company: gofakeit.Company(),
			Status:  gofakeit.RandomString(statuses),
			Phone:   gofakeit.Phone(),
			Email:   gofakeit.Email(),
			Staff:   gofakeit.Number(15, 350),
		}

	}
}

func (s *StoreAdapter) GetCompany(id int64) (*domain.Company, error) {
	if id > int64(s.maxElements) {
		return nil, errors.New("company id is out of the range")
	}

	c := s.company[id-1]

	return &c, nil
}

func (s *StoreAdapter) GetCompanyUpdates(from, to time.Time, status string) []domain.Company {
	var companies []domain.Company

	for _, c := range s.company {
		if c.Updated.After(from) && c.Updated.Before(to) && c.Status == status {
			companies = append(companies, c)
		}
	}

	return companies
}
