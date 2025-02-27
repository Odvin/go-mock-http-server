package store

import "github.com/Odvin/go-mock-http-server/internal/application/domain"

type StoreAdapter struct {
	maxElements int
	cases       []domain.Case
}

func Init(maxElements int) *StoreAdapter {
	cases := make([]domain.Case, maxElements)
	seedCases(cases)

	return &StoreAdapter{
		maxElements: maxElements,
		cases:       cases,
	}
}
