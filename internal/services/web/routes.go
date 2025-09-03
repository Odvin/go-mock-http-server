package web

import "net/http"

func (hs *HttpServer) routes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /v1/healthcheck", hs.healthcheck)
	router.HandleFunc("GET /v1/events", hs.events)

	router.HandleFunc("GET /v1/companies/{id}", hs.getCompany)
	router.HandleFunc("GET /v1/companies/updates", hs.GetCompanyUpdates)
	router.HandleFunc("GET /v1/companies/updates/info", hs.GetCompanyInfo)
	router.HandleFunc("PATCH /v1/companies/updates/stop", hs.StopCompanyUpdates)
	router.HandleFunc("PATCH /v1/companies/updates/start", hs.StartCompanyUpdates)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	return router
}
