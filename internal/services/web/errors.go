package web

import (
	"log"
	"net/http"
)

func logError(r *http.Request, err error) {
	log.Printf("%s %s %v", r.Method, r.URL.String(), err)
}

func errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}

	err := writeJSON(w, status, env, nil)
	if err != nil {
		logError(r, err)
		w.WriteHeader(500)
	}
}

func serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	logError(r, err)

	message := "the server encountered a problem and could not process your request"
	errorResponse(w, r, http.StatusInternalServerError, message)
}

func notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	errorResponse(w, r, http.StatusNotFound, message)
}

func badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	errorResponse(w, r, http.StatusBadRequest, err.Error())
}
