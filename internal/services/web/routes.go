package web

import "net/http"

func (web *WebService) routes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /v1/healthcheck", web.healthcheck)

	router.HandleFunc("GET /v1/company/{id}", web.getCompany)
	router.HandleFunc("GET /v1/company/updates", web.GetCompanyUpdates)

	return router
}
