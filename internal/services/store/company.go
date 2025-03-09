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
	period  int
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

func (s *StoreAdapter) GetCompanyInfo() *domain.CompanyInfo {
	companyInfo := &domain.CompanyInfo{
		Total:    s.maxElements,
		Updating: s.companyUpdater.active,
		Period:   s.companyUpdater.period,
	}

	return companyInfo
}

func (s *StoreAdapter) GetCompany(id int64) (*domain.Company, error) {
	if id > int64(s.maxElements) {
		return nil, errors.New("company id is out of the range")
	}

	c := s.company[id-1]

	return &c, nil
}

func (s *StoreAdapter) GetCompanyUpdates(from, to time.Time, status string, page, size int) ([]domain.Company, int) {
	companies := make([]domain.Company, 0, len(s.company))

	for _, c := range s.company {
		if c.Updated.After(from) && c.Updated.Before(to) && c.Status == status {
			companies = append(companies, c)
		}
	}

	total := len(companies)

	startOffset := (page - 1) * size
	endOffset := startOffset + size

	if startOffset > total {
		return make([]domain.Company, 0), total
	}

	if endOffset > total {
		return companies[startOffset:], total
	}

	return companies[startOffset:endOffset], total
}

func (s *StoreAdapter) StartCompanyUpdates(period int64) {
	s.companyUpdater.Stop()
	s.companyUpdater = NewCompanyUpdater(
		time.Duration(period)*time.Second,
		s.company,
	)
}

func (s *StoreAdapter) StopCompanyUpdates() {
	s.companyUpdater.Stop()
}
