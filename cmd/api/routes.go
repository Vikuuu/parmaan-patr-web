package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// Initialize a new httprouter router instance.
	router := httprouter.New()

	// Register the relevant methods.
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/invoices", app.createInvoiceHandler)
	router.HandlerFunc(http.MethodGet, "/v1/invoices/:id", app.showInvoiceHandler)

	return router
}
