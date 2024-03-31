package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "home...")
}

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
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
