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
	return r
}
