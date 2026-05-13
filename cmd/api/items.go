package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Vikuuu/parmaan-patr-web/internal/data"
	"github.com/Vikuuu/parmaan-patr-web/internal/validator"
)

func (app *application) showItemHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	item, err := app.models.Items.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
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
		Price  int64  `json:"price"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	item := &data.Item{
		Name:   input.Name,
		HsnSac: input.HsnSac,
		Gst:    input.Gst,
		Price:  input.Price,
	}

	// Initialize a new Validator instance.
	v := validator.New()

	if data.ValidateItem(v, item); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Pass the item struct to the Insert() method on our items model.
	err = app.models.Items.Insert(item)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// When sending the HTTP response, we want to include a *Location* header to let
	// the client know which URL they can find the newly-created resource at.
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/items/%d", item.ID))

	// Write a JSON response with 201 Created status code, the item data in the
	// response body, and the Location header.
	err = app.writeJSON(w, http.StatusCreated, envelope{"item": item}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateItemHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// Fetch the existing item record from the database, sending a 404 Not Found
	// response to the client if we couldn't find a matching record.
	item, err := app.models.Items.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Name   *string `json:"name"`
		HsnSac *int64  `json:"hsn_sac"`
		Gst    *int64  `json:"gst"`
		Price  *int64  `json:"price"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Copy the values from the request body to the appropriate fields of the item
	// record. If the input value in nil then we know that no corresponding
	// key/value pair was provided in the JSON request body. So we move on
	// leave the item record unchanged. Otherwise, we update the item
	// record with the new value provided.
	if input.Name != nil {
		item.Name = *input.Name
	}
	if input.HsnSac != nil {
		item.HsnSac = *input.HsnSac
	}
	if input.Gst != nil {
		item.Gst = *input.Gst
	}
	if input.Price != nil {
		item.Price = *input.Price
	}

	v := validator.New()
	if data.ValidateItem(v, item); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Items.Update(item)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"item": item}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.models.Items.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "item successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
