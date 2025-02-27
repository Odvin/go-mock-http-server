package store

import (
	"errors"

	"github.com/Odvin/go-mock-http-server/internal/application/domain"
	"github.com/brianvoe/gofakeit/v7"
)

func seedCases(cases []domain.Case) {
	statuses := []string{"public", "private"}

	for i := range len(cases) {
		created := gofakeit.PastDate()
		cases[i] = domain.Case{
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

func (s *StoreAdapter) GetCase(id int) (*domain.Case, error) {
	if id > s.maxElements || id < 0 {
		return nil, errors.New("invalid index")
	}

	c := s.cases[id]

	return &c, nil
}
