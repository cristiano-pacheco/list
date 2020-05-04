package main

import (
	"encoding/json"
	"errors"
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
	lists, err := app.lists.GetAll()

	if err != nil {
		app.serverError(w, err)
		return
	}

	app.respondwithJSON(w, 200, lists)
}

func (app *application) listStore(w http.ResponseWriter, r *http.Request) {
	list := &models.List{}
	json.NewDecoder(r.Body).Decode(&list)

	listID, err := app.lists.Insert(list.Name)

	if err != nil {
		app.serverError(w, err)
		return
	}

	if listID <= 0 {
		app.serverError(w, models.IsNotPossibleToCreate)
		return
	}

	list, err = app.lists.Get(listID)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.respondwithJSON(w, http.StatusCreated, list)
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

	app.respondwithJSON(w, http.StatusOK, list)
}

func (app *application) listUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	list := &models.List{}
	json.NewDecoder(r.Body).Decode(&list)

	result, err := app.lists.Update(id, list.Name)

	if err != nil {
		app.serverError(w, err)
		return
	}

	if result != true {
		app.serverError(w, models.IsNotPossibleToUpdate)
		return
	}

	list, err = app.lists.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.respondwithJSON(w, http.StatusOK, list)
}

func (app *application) listDestroy(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	result, err := app.lists.Delete(id)

	if err != nil {
		app.serverError(w, err)
		return
	}

	if result != true {
		app.serverError(w, models.IsNotPossibleToDelete)
		return
	}

	app.respondwithJSON(w, http.StatusNoContent, nil)
}
