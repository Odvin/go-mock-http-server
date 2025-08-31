package web

import (
	"net/http"
)

func (hs *HttpServer) healthcheck(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": hs.env,
			"version":     hs.ver,
		},
	}

	err := writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
