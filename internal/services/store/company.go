package store

import (
	"errors"
	"log"
	"time"

	"github.com/Odvin/go-mock-http-server/internal/application/domain"
	"github.com/brianvoe/gofakeit/v7"
)

type companyUpdater struct {
	done    chan struct{}
	stopped chan struct{}
	active  bool
}

func (u companyUpdater) Stop() {
	if u.active {
		close(u.done)
		log.Println("store: company updater is stopping")
		<-u.stopped
		log.Println("store: company updater is stopped")
	}
}

func NewCompanyUpdater(period time.Duration, company []domain.Company) *companyUpdater {
	companyUpdater := &companyUpdater{
		done:    make(chan struct{}),
		stopped: make(chan struct{}),
		active:  true,
	}

	log.Printf("store: company updater stats with period in %f seconds", period.Seconds())

	go func() {
		ticker := time.NewTicker(period)

		defer func() {
			ticker.Stop()
			close(companyUpdater.stopped)
			companyUpdater.active = false
		}()

		for {
			select {
			case <-companyUpdater.done:
				return
			case <-ticker.C:
				updateCompany(company)
			}
		}
	}()

	return companyUpdater
}

func updateCompany(company []domain.Company) {
	updates := gofakeit.Number(1, len(company))

	for range updates {
		update := gofakeit.Number(0, len(company)-1)

		company[update].Updated = time.Now().UTC()
		company[update].Active = gofakeit.Bool()
		company[update].Phone = gofakeit.Phone()
		company[update].Email = gofakeit.Email()
		company[update].Staff = gofakeit.Number(15, 350)

		log.Printf("store: company %d updated", update)
	}

	log.Printf("store: updated %d companies", updates)
}

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

func (s *StoreAdapter) StartCompanyUpdates(period int64) {
	s.companyUpdater.Stop()
	s.companyUpdater = NewCompanyUpdater(time.Duration(period*int64(time.Second)), s.company)
}

func (s *StoreAdapter) StopCompanyUpdates() {
	s.companyUpdater.Stop()
}
