package web

import "net/http"

func (hs *HttpServer) routes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /v1/healthcheck", hs.healthcheck)
	router.HandleFunc("GET /v1/events", hs.events)

	router.HandleFunc("GET /v1/company/{id}", hs.getCompany)
	router.HandleFunc("GET /v1/company/updates", hs.GetCompanyUpdates)
	router.HandleFunc("GET /v1/company/updates/info", hs.GetCompanyInfo)
	router.HandleFunc("PATCH /v1/company/updates/stop", hs.StopCompanyUpdates)
	router.HandleFunc("PATCH /v1/company/updates/start", hs.StartCompanyUpdates)

	return router
}
