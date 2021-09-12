package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) createMovie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) getMovie(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "show the details of movie %s\n", params.ByName("id"))
}
