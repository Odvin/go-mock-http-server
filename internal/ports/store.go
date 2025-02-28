package ports

import (
	"github.com/Odvin/go-mock-http-server/internal/application/domain"
)

type Store interface {
	GetCompany(id int) (*domain.Company, error)
}
