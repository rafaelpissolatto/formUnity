package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/rafaelpissolatto/formUnity/backend/internal/adapter/handler/http"
	"github.com/rafaelpissolatto/formUnity/backend/pkg/logger"
)

// Init loads env vars
func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Create a logger
	logLevel := logger.GetLogLevelFromEnv("LOG_LEVEL", logger.LevelInfo)
	logger := logger.NewLoggerStd(logLevel)

	// Create a new server
	logger.Info("Starting HTTP server...")
	server, err := http.NewServer(ctx, ":8081", http.NewRouter(ctx, logger), logger)
	if err != nil {
		logger.Error(err)
	}

	// Start the server
	if err := server.Start(ctx); err != nil {
		logger.Error(err)
	}

	// Gracefully shutdown the server
	defer func() {
		logger.Info("Shutting down HTTP server...")
		if err := server.Shutdown(ctx); err != nil {
			logger.Error(err)
		}
	}()
}
