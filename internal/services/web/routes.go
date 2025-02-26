package web

import "net/http"

func (web *WebService) routes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /v1/healthcheck", web.healthcheck)

	return router
}
