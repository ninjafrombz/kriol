// Filename: cmd/api/helpers.go

package main

import (
	"errors"
	"net/http"
	"strconv"
	"encoding/json"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	// Use the params from context function to get the request context as a slice
	params := httprouter.ParamsFromContext(r.Context())
	// Get the value of the id parameter
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("Invalid ID Parameter")
	}
	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	//Convert Map into a JSON object
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	//Add a new line to make viewing on the terminal easier
	js = append(js, '\n')
	//Add the headers that have been provided to us 
	for key, value := range headers{
		w.Header()[key] = value
	}
	// SPecify that we will serve our responses using JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	// Write the byte slice containing the JSON response body
	w.Write(js)
	return nil
}