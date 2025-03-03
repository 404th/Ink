package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/404th/Ink/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Run starts the HTTP server with graceful shutdown
func Run(cfg *config.Config, sugar *zap.SugaredLogger, r *gin.Engine) {
	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.ProjectPort // fallback to config
	}

	// Configure the server
	srv := &http.Server{
		Addr:    cfg.ProjectHost + ":" + port,
		Handler: r,
	}

	// Channel to listen for OS signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		sugar.Infof("Starting server on %s:%s", cfg.ProjectHost, port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			sugar.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for termination signal
	<-sigChan
	sugar.Infoln("Shutdown signal received, shutting down server...")

	// Create a context for graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		sugar.Fatalf("HTTP shutdown error: %v", err)
	}

	sugar.Infoln("Server gracefully stopped.")
}
