package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Vikuuu/parmaan-patr-web/internal/data"
)

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

func (app *application) createInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID         int64     `json:"id"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		To         string    `json:"to"`
		From       string    `json:"from"`
		Items      []string  `json:"items"`
		TotalPrice uint32    `json:"total_price"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintln(w, "%+v\n", input)
}
