package handler

import (
	"second_project/internal/middleware"

	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddleware.StripSlashes)

	r.Route("/account", func(r chi.Router) {
		r.Use(middleware.Authorization)

		r.Get("/coins", GetCoinBalance)
	})
}
