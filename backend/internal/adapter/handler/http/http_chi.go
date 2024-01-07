package http

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NewServer returns a new HTTP server
func NewServer(ctx context.Context, addr string, handler http.Handler) *http.Server {
	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return srv
}

// NewRouter returns a new HTTP router
func NewRouter(ctx context.Context) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Router is working! (GET /)"))
		log.Println("Router is working! (GET /)")
	})
	return r
}
