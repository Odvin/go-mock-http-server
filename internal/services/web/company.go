package web

import "net/http"

func (web *WebService) getCompany(w http.ResponseWriter, r *http.Request) {
	id, err := readIDParam(r)
	if err != nil {
		notFoundResponse(w, r)
		return
	}

	compnay, err := web.api.GetCompany(id)
	if err != nil {
		serverErrorResponse(w, r, err)
	}

	err = writeJSON(w, http.StatusOK, envelope{"company": compnay}, nil)
	if err != nil {
		serverErrorResponse(w, r, err)
	}
}
