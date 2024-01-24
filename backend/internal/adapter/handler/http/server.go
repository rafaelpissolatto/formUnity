package http

import (
	"context"
	"net/http"

	"github.com/rafaelpissolatto/formUnity/backend/pkg/logger"
)

// Server is a HTTP server.
type Server struct {
	addr    string
	handler http.Handler
	logger  logger.Logger
}

// Option is a function that configures a server.
type Option func(*Server)

// NewServer returns a new HTTP server.
func NewServer(ctx context.Context, addr string, handler http.Handler, logger logger.Logger, opts ...Option) (*Server, error) {
	srv := &Server{
		addr:    addr,
		handler: handler,
		logger:  logger,
	}
	for _, opt := range opts {
		opt(srv)
	}
	return srv, nil
}

// Start starts the HTTP server.
func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("Starting HTTP server...")
	return http.ListenAndServe(s.addr, s.handler)
}

// Shutdown gracefully shuts down the HTTP server.
func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("Shutting down HTTP server...")
	// if SIGTERM is received, the context is canceled
	return nil
}
