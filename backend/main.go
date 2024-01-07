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
	logger.Debug("HTTP server stopped")
	srv := http.NewServer(ctx, ":8080", http.NewRouter(ctx))
	if err := srv.ListenAndServe(); err != nil {
		logger.Error(err)
	}

}
