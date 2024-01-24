package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rafaelpissolatto/formUnity/backend/pkg/logger"
)

// NewRouter returns a new HTTP router
func NewRouter(ctx context.Context, logger logger.Logger) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Router is working! (GET /)"))
		logger.Info("Router is working! (GET /)")
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("API v1 is working! (GET /api/v1)"))
			logger.Info("API v1 is working! (GET /api/v1)")
		})

		// TODO: add routes for api/v1
		r.Route("/forms", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("API v1 forms is working! (GET /api/v1/forms)"))
				logger.Info("API v1 forms is working! (GET /api/v1/forms)")

			})
		})

		r.Route("/forms/{id}", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("API v1 forms is working! (GET /api/v1/forms/{id})"))
				logger.Info("API v1 forms is working! (GET /api/v1/forms/{id})")
			})
		})

		r.Route("/forms", func(r chi.Router) {
			r.Post("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("API v1 forms is working! (POST /api/v1/forms)"))
				logger.Info("API v1 forms is working! (POST /api/v1/forms)")
			})
		})

		r.Route("/forms/{id}", func(r chi.Router) {
			r.Put("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("API v1 forms is working! (PUT /api/v1/forms/{id})"))
				logger.Info("API v1 forms is working! (PUT /api/v1/forms/{id})")
			})
		})

		r.Route("/forms/{id}", func(r chi.Router) {
			r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("API v1 forms is working! (DELETE /api/v1/forms/{id})"))
				logger.Info("API v1 forms is working! (DELETE /api/v1/forms/{id})")
			})
		})
	})

	return r
}

// GET /api/v1/forms
// GET /api/v1/forms/{id}
// POST /api/v1/forms
// PUT /api/v1/forms/{id}
// DELETE /api/v1/forms/{id}

// v1 router
