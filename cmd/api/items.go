package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Vikuuu/parmaan-patr-web/internal/data"
)

func (app *application) showItemHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	item := data.Item{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "user1",
		HsnSac:    123,
		Gst:       18,
		Price:     "123.23",
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"item": item}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createItemHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name   string `json:"name"`
		HsnSac int64  `json:"hsn_sac"`
		Gst    int64  `json:"gst"`
		Price  string `json:"price"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintln(w, "%+v\n", input)
}
