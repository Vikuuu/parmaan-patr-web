package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// Initialize a new httprouter router instance.
	router := httprouter.New()

	// Custom error handler for 404 not found and 405 method not
	// allowed error responses.
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// Register the relevant methods.
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	// Invoice related paths
	router.HandlerFunc(http.MethodPost, "/v1/invoices", app.createInvoiceHandler)
	router.HandlerFunc(http.MethodGet, "/v1/invoices/:id", app.showInvoiceHandler)

	// Items related paths
	router.HandlerFunc(http.MethodPost, "/v1/items", app.createItemHandler)
	router.HandlerFunc(http.MethodGet, "/v1/items/:id", app.showItemHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/items/:id", app.updateItemHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/items/:id", app.deleteItemHandler)
	return router
}
