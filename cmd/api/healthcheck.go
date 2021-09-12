package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := envelope{
		"status": "available",
		"systemInfo": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}
