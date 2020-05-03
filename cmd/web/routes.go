package main

import (
	"github.com/go-chi/chi"
)

func (app *application) routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", app.home)
	return router
}
