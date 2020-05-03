package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	w.Write([]byte("Oh Yeah!"))
}

// lists
func (app *application) listIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List Index"))
}

func (app *application) listStore(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List Store"))
}

func (app *application) listShow(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Fprintf(w, "List Show %s", id)
}

func (app *application) listUpdate(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Fprintf(w, "List Update %s", id)
}

func (app *application) listDestroy(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Fprintf(w, "List Destroy %s", id)
}
