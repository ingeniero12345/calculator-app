package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/linktic/calculator-app/backend/internal/api"
	"github.com/linktic/calculator-app/backend/internal/service"
)

const (
	defaultPort          = "8080"
	readHeaderTimeout    = 5 * time.Second
	gracefulShutdownWait = 10 * time.Second
)

func main() {
	address := ":" + envOr("PORT", defaultPort)

	server := &http.Server{
		Addr:              address,
		Handler:           api.NewRouter(service.NewCalculator()),
		ReadHeaderTimeout: readHeaderTimeout,
	}

	go func() {
		log.Printf("calculator API listening on %s", address)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdownWait)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("graceful shutdown failed: %v", err)
	}
	log.Println("server stopped")
}

func envOr(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
