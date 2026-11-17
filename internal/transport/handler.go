package transport

import (
	"encoding/json"
	"net/http"

	"github.com/liamoj10/liamsplanets/internal/simulation"
)

// Handler handles HTTP requests for the simulation.
type Handler struct {
	engine *simulation.Engine
}

// NewHandler creates a transport handler wrapping the simulation engine.
func NewHandler(engine *simulation.Engine) *Handler {
	return &Handler{engine: engine}
}

// HandleGetState handles GET /state: returns current planetary coordinates as JSON.
func (h *Handler) HandleGetState(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	state := h.engine.GetState()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(state); err != nil {
		http.Error(w, "Failed to encode state", http.StatusInternalServerError)
		return
	}
}

// TODO: HandleWebSocket (When ready to upgrade from polling to real-time stream)
// func (h *Handler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
// 	// Upgrade HTTP connection to WebSocket (e.g. using gorilla/websocket or nhooyr.io/websocket)
// 	// Stream simulation ticks directly to the client at 60fps
// }
