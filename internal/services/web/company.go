package web

import (
	"fmt"
	"net/http"
)

func (web *WebService) getCompany(w http.ResponseWriter, r *http.Request) {
	id, err := readPathInt(r, "id")
	if err != nil {
		badRequestResponse(w, r, fmt.Errorf("id : %w", err))
		return
	}

	company, err := web.api.GetCompany(id)
	if err != nil {
		badRequestResponse(w, r, fmt.Errorf("id : %w", err))
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"company": company}, nil)
	if err != nil {
		serverErrorResponse(w, r, err)
	}
}

func (web *WebService) GetCompanyUpdates(w http.ResponseWriter, r *http.Request) {
	from, err := readQueryTime(r, "from")
	if err != nil {
		badRequestResponse(w, r, fmt.Errorf("from : %w", err))
		return
	}

	to, err := readQueryTime(r, "to")
	if err != nil {
		badRequestResponse(w, r, fmt.Errorf("to : %w", err))
		return
	}

	status, err := readQueryStr(r, "status")
	if err != nil {
		badRequestResponse(w, r, fmt.Errorf("status : %w", err))
		return
	}

	companies := web.api.GetCompanyUpdates(from, to, status)

	err = writeJSON(w, http.StatusOK, envelope{"companies": companies}, nil)
	if err != nil {
		serverErrorResponse(w, r, err)
	}
}
