package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type wrap_json map[string]interface{}

func (app *application) parse_key(r *http.Request) (int64, error) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid key")
	}
	return id, nil
}

func (app *application) w_json(w http.ResponseWriter, status int, data wrap_json, headers http.Header, indent bool) error {
	var payload []byte
	var err error
	if !indent {
		payload, err = json.Marshal(data)
		if err != nil {
			return err
		}
	} else {
		payload, err = json.MarshalIndent(data, "", "\t")
		if err != nil {
			return err
		}
	}

	payload = append(payload, '\n')

	for k, v := range headers {
		w.Header()[k] = v
	}

	w.Header().Set("Content-Type", "application/json")
	// TODO: write defensive logic against flawed or failed status codes
	w.WriteHeader(status)
	w.Write(payload)
	return nil
}
