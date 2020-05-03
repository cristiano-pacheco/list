package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cristiano-pacheco/list/pkg/models"
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
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	list, err := app.lists.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.respondwithJSON(w, 200, list)
}

func (app *application) listUpdate(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Fprintf(w, "List Update %s", id)
}

func (app *application) listDestroy(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Fprintf(w, "List Destroy %s", id)
}
