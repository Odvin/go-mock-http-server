package web

import (
	"encoding/json"
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

	var page, size int64

	page, err = readQueryInt(r, "page")
	if err != nil || page < 1 {
		page = 1
	}

	size, err = readQueryInt(r, "size")
	if err != nil || size < 1 {
		size = 20
	}

	companies, total := web.api.GetCompanyUpdates(from, to, status, int(page), int(size))

	subset := map[string]int64{
		"total": int64(total),
		"page":  page,
		"size":  size,
	}

	err = writeJSON(w, http.StatusOK, envelope{"companies": companies, "subset": subset}, nil)
	if err != nil {
		serverErrorResponse(w, r, err)
	}
}

func (web *WebService) StopCompanyUpdates(w http.ResponseWriter, r *http.Request) {
	web.api.StopCompanyUpdates()

	err := writeJSON(w, http.StatusOK, envelope{"updates": "company", "stopped": true}, nil)
	if err != nil {
		serverErrorResponse(w, r, err)
	}
}

func (web *WebService) StartCompanyUpdates(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Period int64 `json:"period"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		serverErrorResponse(w, r, err)
	}

	err = web.api.StartCompanyUpdates(input.Period)
	if err != nil {
		badRequestResponse(w, r, fmt.Errorf("body : %w", err))
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"updates": "company", "stopped": false}, nil)
	if err != nil {
		serverErrorResponse(w, r, err)
	}
}

func (web *WebService) GetCompanyInfo(w http.ResponseWriter, r *http.Request) {
	companyInfo := web.api.GetCompanyInfo()

	err := writeJSON(w, http.StatusOK, envelope{"company_info": companyInfo}, nil)
	if err != nil {
		serverErrorResponse(w, r, err)
	}
}
