package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/permalik/temp_rest_go/internal/data"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "home...")
}

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	data := wrap_json{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.Env,
			"version":     "0.1.0",
		},
	}
	err := app.w_json(w, http.StatusOK, data, nil, true)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (app *application) create_item(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create item...")
}

func (app *application) read_items(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "read items...")
}

func (app *application) read_item(w http.ResponseWriter, r *http.Request) {
	id, err := app.parse_key(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	data := data.Item{
		ID:        id,
		Name:      "test_item",
		Quantity:  12,
		Pounds:    15,
		Types:     []string{"primary", "secondary", "tertiary"},
		CreatedAt: time.Now(),
	}

	err = app.w_json(w, http.StatusOK, wrap_json{"item": data}, nil, true)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
