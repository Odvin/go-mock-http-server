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
			ID:      i,
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

func (s *StoreAdapter) GetCompany(id int) (*domain.Company, error) {
	if id > s.maxElements || id < 0 {
		return nil, errors.New("invalid index")
	}

	c := s.company[id]

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
