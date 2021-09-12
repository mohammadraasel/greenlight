package main

import "github.com/julienschmidt/httprouter"

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.GET("/v1/healthcheck", app.healthcheck)
	router.POST("/v1/movies", app.createMovie)
	router.GET("/v1/movies/:id", app.getMovie)

	return router
}
