package http

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NewRouter returns a new HTTP router
func NewRouter(ctx context.Context) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Router is working! (GET /)"))
		log.Println("Router is working! (GET /)")
	})
	return r
}
