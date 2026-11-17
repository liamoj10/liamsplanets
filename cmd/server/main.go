package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/liamoj10/liamsplanets/internal/simulation"
	"github.com/liamoj10/liamsplanets/internal/transport"
)

func main() {
	// 1. Initialize Simulation Engine
	engine := simulation.NewEngine()

	// Start simulation loop at 60 ticks per second (~16ms)
	engine.Start(16 * time.Millisecond)

	// 2. Initialize Transport Handler
	handler := transport.NewHandler(engine)

	// 3. Routing
	mux := http.NewServeMux()

	// Serve the frontend static assets (HTML/JS/CSS) from ./web
	fileServer := http.FileServer(http.Dir("./web"))
	mux.Handle("/", fileServer)

	// Endpoint polled by the UI: GET /state
	mux.HandleFunc("/state", handler.HandleGetState)

	// Placeholder endpoint for future WebSocket connection
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("WebSocket simulation stream endpoint (Coming soon)"))
	})

	// 4. Configure HTTP Server
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 5. Start server in a background goroutine
	go func() {
		log.Printf("Starting server on http://localhost:8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// 6. Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server gracefully...")
	engine.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped cleanly.")
}
