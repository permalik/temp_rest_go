package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "home...")
}

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"environment": app.config.env,
		"status":      "available",
		"version":     version,
	}
	err := app.wJson(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (app *application) createItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create item...")
}

func (app *application) readItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "read items...")
}

func (app *application) readItem(w http.ResponseWriter, r *http.Request) {
	id, err := app.parskey(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "read item %d...", id)
}
