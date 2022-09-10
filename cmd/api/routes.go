package main 

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)


func (app *application) routes() *httprouter.Router {
	// Create
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/entries", app.CreateEntryHandler)
	router.HandlerFunc(http.MethodGet, "/v1/entries/:id", app.showEntryHandler)

	return router
}