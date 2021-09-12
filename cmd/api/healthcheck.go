package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) healthcheck(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Status: available\n")
	fmt.Fprintf(w, "Environment: %s\n", app.config.env)
	fmt.Fprintf(w, "Version: %s\n", version)
}
