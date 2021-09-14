package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.GET("/v1/healthcheck", app.healthcheck)
	router.POST("/v1/movies", app.createMovie)
	router.GET("/v1/movies/:id", app.getMovie)
	router.PUT("/v1/movies/:id", app.updateMovie)
	router.DELETE("/v1/movies/:id", app.deleteMovie)

	return router
}
