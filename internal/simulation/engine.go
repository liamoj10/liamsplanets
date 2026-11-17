package simulation

import (
	"sync"
	"time"
)

// Body represents a single celestial body in the simulator.
// Matches the JSON format expected by the frontend UI.
type Body struct {
	Name   string  `json:"name"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	VX     float64 `json:"vx"`
	VY     float64 `json:"vy"`
	Radius float64 `json:"radius"`
	Color  string  `json:"color"`
	Mass   float64 `json:"mass,omitempty"` // For gravity calculations
}

// Engine manages the simulation state, bodies, and time loop.
type Engine struct {
	mu     sync.RWMutex
	bodies []Body
	ticker *time.Ticker
	stopCh chan struct{}
}

// NewEngine creates an engine initialized with the inner solar system in classical Latin.
// The bodies are placed at scaled distances to fit small laptop viewports by default.
// Velocities are kept at 0 so they remain visible and stationary until the physics integrator is implemented.
func NewEngine() *Engine {
	// Central star (Sol): Mass = 50000, positioned at origin (0, 0)
	// Target circular speeds sqrt(50000 / r) are noted in comments for when physics is activated.
	return &Engine{
		bodies: []Body{
			{
				Name:   "Sol",
				X:      0,
				Y:      0,
				VX:     0,
				VY:     0,
				Radius: 16,
				Color:  "#facc15",
				Mass:   50000,
			},
			{
				Name:   "Mercurius",
				X:      60,
				Y:      0,
				VX:     0, // Target Vy ≈ 28.87 once physics engine is active
				VY:     0,
				Radius: 3.5,
				Color:  "#a3a3a3",
				Mass:   1.6,
			},
			{
				Name:   "Venus",
				X:      100,
				Y:      0,
				VX:     0, // Target Vy ≈ 22.36 once physics engine is active
				VY:     0,
				Radius: 5.5,
				Color:  "#eab308",
				Mass:   8.2,
			},
			{
				Name:   "Terra",
				X:      150,
				Y:      0,
				VX:     0, // Target Vy ≈ 18.26 once physics engine is active
				VY:     0,
				Radius: 6.0,
				Color:  "#38bdf8",
				Mass:   10.0,
			},
			{
				Name:   "Mars",
				X:      210,
				Y:      0,
				VX:     0, // Target Vy ≈ 15.43 once physics engine is active
				VY:     0,
				Radius: 4.5,
				Color:  "#f87171",
				Mass:   5.0,
			},
		},
		stopCh: make(chan struct{}),
	}
}

// Start kicks off the background physics simulation loop.
func (e *Engine) Start(tickRate time.Duration) {
	e.ticker = time.NewTicker(tickRate)
	go func() {
		lastTime := time.Now()
		for {
			select {
			case now := <-e.ticker.C:
				dt := now.Sub(lastTime).Seconds()
				lastTime = now
				e.Step(dt)
			case <-e.stopCh:
				e.ticker.Stop()
				return
			}
		}
	}()
}

// Stop stops the simulation loop cleanly.
func (e *Engine) Stop() {
	close(e.stopCh)
}

// GetState returns a snapshot copy of the current bodies for thread safety.
func (e *Engine) GetState() []Body {
	e.mu.RLock()
	defer e.mu.RUnlock()

	result := make([]Body, len(e.bodies))
	copy(result, e.bodies)
	return result
}

// Step advances the physics simulation by dt seconds.
func (e *Engine) Step(dt float64) {
	e.mu.Lock()
	defer e.mu.Unlock()

	// TODO: Step 1 (Physics Calculation)
	// Calculate gravitational forces (N-body gravity or 2-body Keplerian gravity)
	// Hint: F = G * (m1 * m2) / (r^2 + softening^2)
	// Update e.bodies[i].VX and e.bodies[i].VY based on acceleration (a = F / m)

	// Step 2: Position integration (simple Euler integration placeholder)
	for i := range e.bodies {
		e.bodies[i].X += e.bodies[i].VX * dt
		e.bodies[i].Y += e.bodies[i].VY * dt
	}

	// TODO: Step 3 (Collisions / Merging / Bounds)
	// Check for planet collisions with the Sun or other bodies if desired.
}
