package web

import (
	"encoding/json"
	"errors"
	"maps"
	"net/http"
	"strconv"
	"time"
)

type envelope map[string]any

func writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	maps.Copy(w.Header(), headers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func readIDParam(r *http.Request) (int, error) {
	paramId := r.PathValue("id")

	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return int(id), nil
}

func readQueryTime(r *http.Request, param string) (time.Time, error) {
	queryTime := r.URL.Query().Get(param)

	t, err := time.Parse(time.RFC3339, queryTime)
	if err != nil {
		return time.Time{}, errors.New("unrecognized time format")
	}

	return t, nil
}
