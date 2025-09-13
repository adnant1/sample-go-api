package handlers

import (
	"github.com/adnant1/sample-go-api/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

// Capital letter H in Handler makes it public
func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		// Middleware for /account route
		router.Use(middleware.Authenticate)

		// GET /account/coins
		router.Get("/coins", GetCoinBalance)
	})
}