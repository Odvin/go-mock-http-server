package web

import (
	"encoding/json"
	"errors"
	"fmt"
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

	w.Header().Set("Content-Type", "app/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func readPathInt(r *http.Request, path string) (int64, error) {
	pathValue := r.PathValue(path)

	value, err := strconv.ParseInt(pathValue, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%s invalid path value", pathValue)
	}

	return value, nil
}

func readQueryTime(r *http.Request, param string) (time.Time, error) {
	queryTime := r.URL.Query().Get(param)

	t, err := time.Parse(time.RFC3339, queryTime)
	if err != nil {
		return time.Time{}, errors.New("unrecognized time format")
	}

	return t, nil
}

func readQueryStr(r *http.Request, param string) (string, error) {
	s := r.URL.Query().Get(param)

	if s == "" {
		return "", errors.New("value was not provided")
	}

	return s, nil
}

func readQueryInt(r *http.Request, param string) (int64, error) {
	queryValue := r.URL.Query().Get(param)

	value, err := strconv.ParseInt(queryValue, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%s invalid path value", queryValue)
	}

	return value, nil
}
