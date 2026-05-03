package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Vikuuu/parmaan-patr-web/internal/data"
)

func (app *application) createInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new invoice")
}

func (app *application) showInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	invoice := data.Invoice{
		ID:         id,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		To:         "user1",
		From:       "user2",
		Items:      []string{"shirt", "pant"},
		TotalPrice: uint32(1234),
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"invoice": invoice}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
