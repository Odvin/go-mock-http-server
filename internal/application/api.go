package application

import (
	"github.com/Odvin/go-mock-http-server/internal/application/domain"
	"github.com/Odvin/go-mock-http-server/internal/ports"
)

type API interface {
	getCase(id int) (*domain.Case, error)
}

type Application struct {
	store ports.Store
}

func Init(store ports.Store) *Application {
	return &Application{
		store: store,
	}
}
