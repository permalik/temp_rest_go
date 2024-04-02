package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type wrap_json map[string]interface{}

func (app *application) parse_key(r *http.Request) (int64, error) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid key")
	}
	return id, nil
}

func (app *application) r_json(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	max_bytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(max_bytes))
	data := json.NewDecoder(r.Body)
	data.DisallowUnknownFields()

	err := data.Decode(dst)
	if err != nil {
		var syntax_error *json.SyntaxError
		var unmarshal_type_error *json.UnmarshalTypeError
		var invalid_unmarshal_error *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntax_error):
			return fmt.Errorf("malformed json. char %d", syntax_error.Offset)
		// TODO: remove once resolved: https://github.com/golang/go/issues/25956
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("malformed json.")
		case errors.As(err, &unmarshal_type_error):
			if unmarshal_type_error.Field != "" {
				return fmt.Errorf("invalid json type. field %q", unmarshal_type_error.Field)
			}
			return fmt.Errorf("invalid json type. char %d", unmarshal_type_error.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("invalid. empty body.")
		case strings.HasPrefix(err.Error(), "json: unknown field"):
			field_name := strings.TrimPrefix(err.Error(), "json: unknown field")
			return fmt.Errorf("unknown field. %s", field_name)
		// TODO: refactor if resolved and distinct error type is created: https://github.com/golang/go/issues/30715
		case err.Error() == "http: request body too large":
			return fmt.Errorf("body <= %d bytes", max_bytes)
		case errors.As(err, &invalid_unmarshal_error):
			panic(err)
		default:
			return err
		}
	}

	err = data.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body may only contain single json value")
	}

	return nil
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
