package main

import (
	"github.com/go-chi/chi"
)

func (app *application) routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", app.home)

	router.Get("/lists", app.listIndex)
	router.Post("/lists", app.listStore)
	router.Get("/lists/{id}", app.listShow)
	router.Put("/lists/{id}", app.listUpdate)
	router.Delete("/lists/{id}", app.listDestroy)

	return router
}
