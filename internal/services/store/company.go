package store

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Odvin/go-mock-http-server/internal/app"
	"github.com/brianvoe/gofakeit/v7"
)

type CompanyUpdater struct {
	done    chan struct{}
	stopped chan struct{}
	active  bool
	period  int
}

func (u CompanyUpdater) Stop() {
	if u.active {
		close(u.done)
		log.Println("store: company updater is stopping")
		<-u.stopped
		log.Println("store: company updater is stopped")
	}
}

func NewCompanyUpdater(period time.Duration, company []app.Company) *CompanyUpdater {
	companyUpdater := &CompanyUpdater{
		done:    make(chan struct{}),
		stopped: make(chan struct{}),
		active:  true,
		period:  int(period.Seconds()),
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

func updateCompany(companies []app.Company) {
	updates := gofakeit.Number(1, len(companies))

	for range updates {
		id := gofakeit.Number(0, len(companies)-1)

		companies[id].Updated = time.Now().UTC()
		companies[id].Active = gofakeit.Bool()
		companies[id].Phone = gofakeit.Phone()
		companies[id].Email = gofakeit.Email()
		companies[id].Staff = gofakeit.Number(15, 350)

		log.Printf("store: company %d updated", id)
	}

	ps.Publish("UpdateCompany", fmt.Sprintf("store: updated %d companies", updates))
}

func seedCompany(companies []app.Company) {
	statuses := []string{"public", "private"}

	var created time.Time
	for i := range len(companies) {
		created = gofakeit.Date()

		companies[i] = app.Company{
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

func (s *Store) GetCompanyInfo() *app.CompanyInfo {
	companyInfo := &app.CompanyInfo{
		Total:    len(s.companies),
		Updating: s.companyUpdater.active,
		Period:   s.companyUpdater.period,
	}

	return companyInfo
}

func (s *Store) GetCompany(id int64) (*app.Company, error) {
	if id > int64(len(s.companies)) {
		return nil, errors.New("company id is out of the range")
	}

	c := s.companies[id-1]

	return &c, nil
}

func (s *Store) GetCompanyUpdates(from, to time.Time, status string, page, size int) ([]app.Company, int) {
	updatedCompanies := make([]app.Company, 0, len(s.companies))

	for _, c := range s.companies {
		if c.Updated.After(from) && c.Updated.Before(to) && c.Status == status {
			updatedCompanies = append(updatedCompanies, c)
		}
	}

	total := len(updatedCompanies)

	startOffset := (page - 1) * size
	endOffset := startOffset + size

	if startOffset > total {
		return make([]app.Company, 0), total
	}

	if endOffset > total {
		return updatedCompanies[startOffset:], total
	}

	return updatedCompanies[startOffset:endOffset], total
}

func (s *Store) StartCompanyUpdates(period int64) {
	s.companyUpdater.Stop()
	s.companyUpdater = NewCompanyUpdater(
		time.Duration(period)*time.Second,
		s.companies,
	)
}

func (s *Store) StopCompanyUpdates() {
	s.companyUpdater.Stop()
}
