package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func (app *application) parskey(r *http.Request) (int64, error) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid key")
	}
	return id, nil
}

func (app *application) wJson(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}
	payload = append(payload, '\n')
	for k, v := range headers {
		w.Header()[k] = v
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(payload)
	return nil
}
