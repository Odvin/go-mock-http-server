package application

import (
	"github.com/Odvin/go-mock-http-server/internal/application/domain"
	"github.com/Odvin/go-mock-http-server/internal/ports"
)

type API interface {
	GetCompany(id int) (*domain.Company, error)
}

type Application struct {
	store ports.Store
}

func Init(store ports.Store) *Application {
	return &Application{
		store: store,
	}
}
