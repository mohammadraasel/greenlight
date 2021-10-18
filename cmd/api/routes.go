package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.GET("/v1/healthcheck", app.healthcheck)

	router.GET("/v1/movies", app.requirePermission("movies:read", app.getMovies))
	router.POST("/v1/movies", app.requirePermission("movies:write", app.createMovie))
	router.GET("/v1/movies/:id", app.requirePermission("movies:read", app.getMovie))
	router.PATCH("/v1/movies/:id", app.requirePermission("movies:write", app.updateMovie))
	router.DELETE("/v1/movies/:id", app.requirePermission("movies:write", app.deleteMovie))

	router.POST("/v1/users", app.registerUser)
	router.PUT("/v1/users/activated", app.activateUser)
	router.POST("/v1/tokens/authentication", app.createAuthenticationToken)

	return app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router))))
}
