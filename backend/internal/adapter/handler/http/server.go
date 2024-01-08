package http

import (
	"context"
	"net/http"
)

// NewServer returns a new HTTP server
func NewServer(ctx context.Context, addr string, handler http.Handler) *http.Server {
	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return srv
}
