package main

import (
	"fmt"
	"net/http"
)

func (app *application) createInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new invoice")
}

func (app *application) showInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of invoice %d\n", id)
}
