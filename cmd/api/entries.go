// Filename: cmd/api/entries.go

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) CreateEntryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new entry ...")
}

func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request) {
	// Use the params from context function to get the request context as a slice
	params := httprouter.ParamsFromContext(r.Context())
	// Get the value of the id parameter
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	//Display the school id 
	fmt.Fprintf(w, "show the details for entries %d\n", id)
}
