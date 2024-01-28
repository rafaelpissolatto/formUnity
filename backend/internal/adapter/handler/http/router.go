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
	r.Mount("/api/v1", apiV1Router(ctx, logger))
	return r
}

// Func to be router for api/v1
func apiV1Router(ctx context.Context, logger logger.Logger) http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Router is working! (GET /api/v1)"))
		logger.Info("Router is working! (GET /api/v1)")
	})
	return r
}
