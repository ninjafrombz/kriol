// Filename: cmd/api/entries.go

package main

import (
	"fmt"
	"net/http"
)

func (app *application) CreateEntryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new entry ...")
}

func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	//Display the school id
	fmt.Fprintf(w, "show the details for entries %d\n", id)
}
