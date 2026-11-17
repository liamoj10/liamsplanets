package physics

import "github.com/liamoj10/liamsplanets/internal/simulation"

// Integrator is the Strategy interface for physics stepping algorithms.
// Refactoring Guru: https://refactoring.guru/design-patterns/strategy
//
// The Strategy pattern lets you define a family of algorithms, put each of them
// into a separate struct, and make their objects interchangeable at runtime.
type Integrator interface {
	// Step calculates gravitational forces and integrates velocity & position
	// over time delta dt (in seconds).
	Step(bodies []simulation.Body, dt float64) []simulation.Body
}

// EulerIntegrator implements basic Euler numerical integration:
// x(t + dt) = x(t) + v(t) * dt
// v(t + dt) = v(t) + a(t) * dt
type EulerIntegrator struct {
	G float64 // Gravitational constant
}

// Step performs the Euler integration.
func (e *EulerIntegrator) Step(bodies []simulation.Body, dt float64) []simulation.Body {
	// TODO: Implement Euler integration:
	// 1. For each body i, compute net gravitational force from every other body j:
	//    F = G * (m1 * m2) / (dist^2 + softening^2)
	// 2. Acceleration a = F / m1
	// 3. Update velocity: v += a * dt
	// 4. Update position: x += v * dt
	return bodies
}

// VerletIntegrator implements Velocity-Verlet (symplectic) integration.
// Much higher energy conservation for orbital mechanics than Euler.
type VerletIntegrator struct {
	G float64
}

// Step performs Velocity-Verlet integration.
func (v *VerletIntegrator) Step(bodies []simulation.Body, dt float64) []simulation.Body {
	// TODO: Implement Velocity Verlet integration:
	// 1. x(t + dt) = x(t) + v(t)*dt + 0.5*a(t)*dt^2
	// 2. Compute a(t + dt) using new positions
	// 3. v(t + dt) = v(t) + 0.5*(a(t) + a(t + dt))*dt
	return bodies
}

// RungeKutta4Integrator implements 4th-order Runge-Kutta (RK4).
// Highly accurate classic integrator for celestial mechanics.
type RungeKutta4Integrator struct {
	G float64
}

// Step performs RK4 integration.
func (rk *RungeKutta4Integrator) Step(bodies []simulation.Body, dt float64) []simulation.Body {
	// TODO: Implement RK4:
	// Sample slopes k1, k2, k3, k4 at intervals 0, dt/2, dt/2, dt
	// Weighted average: (k1 + 2*k2 + 2*k3 + k4) / 6
	return bodies
}
