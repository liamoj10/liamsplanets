package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 1. Initialize routing using standard library ServeMux
	mux := http.NewServeMux()

	// Serve the frontend static assets (HTML/JS/Canvas) from ./web
	fileServer := http.FileServer(http.Dir("./web"))
	mux.Handle("/", fileServer)

	// Placeholder endpoint for WebSocket connection (Target: SOL-02)
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("WebSocket simulation stream endpoint (Coming in SOL-02)"))
	})

	// 2. Configure the HTTP server
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 3. Start server in a background goroutine so it doesn't block shutdown handling
	go func() {
		log.Printf("Starting server on http://localhost:8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// 4. Set up graceful shutdown channel listening for SIGINT (Ctrl+C) or SIGTERM
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server gracefully...")

	// Give active connections/simulation loops 5 seconds to wrap up
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped cleanly.")
}
