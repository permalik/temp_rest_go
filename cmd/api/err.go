package main

import (
	"fmt"
	"net/http"
)

// TODO: determine purpose of request arg
// error
func (app *application) err(_ *http.Request, err error) {
	app.logger.Println(err)
}

// error response
func (app *application) err_res(w http.ResponseWriter, r *http.Request, status int, msg interface{}) {
	data := wrap_json{"err": msg}
	err := app.w_json(w, status, data, nil, true)
	if err != nil {
		app.err(r, err)
		w.WriteHeader(500)
	}
}

// error response: internal server
func (app *application) err_res_int_srv(w http.ResponseWriter, r *http.Request, err error) {
	app.err(r, err)
	msg := "internal server error"
	app.err_res(w, r, http.StatusInternalServerError, msg)
}

// error response: not found
func (app *application) err_res_notf(w http.ResponseWriter, r *http.Request) {
	msg := "resource not found"
	app.err_res(w, r, http.StatusNotFound, msg)
}

// error response: method not allowed
func (app *application) err_res_method_nota(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("unsupported method: %s", r.Method)
	app.err_res(w, r, http.StatusMethodNotAllowed, msg)
}

// error response: bad request
func (app *application) err_res_bad_req(w http.ResponseWriter, r *http.Request, err error) {
	app.err_res(w, r, http.StatusBadRequest, err.Error())
}
