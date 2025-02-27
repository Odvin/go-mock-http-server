package ports

import (
	"time"

	"github.com/Odvin/go-mock-http-server/internal/application/domain"
)

type Store interface {
	GetCase(id int) (*domain.Case, error)
	GetCaseUpdates(from time.Time, to time.Time, status string) ([]domain.Case, error)
}
