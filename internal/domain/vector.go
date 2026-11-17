package domain

import "math"

// Vector2D represents a 2D spatial vector (position, velocity, acceleration).
type Vector2D struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Add returns the vector sum (v + other).
func (v Vector2D) Add(other Vector2D) Vector2D {
	// TODO: Implement vector addition
	return Vector2D{X: v.X + other.X, Y: v.Y + other.Y}
}

// Sub returns the vector subtraction (v - other).
func (v Vector2D) Sub(other Vector2D) Vector2D {
	// TODO: Implement vector subtraction
	return Vector2D{X: v.X - other.X, Y: v.Y - other.Y}
}

// Mul returns the vector scaled by scalar s.
func (v Vector2D) Mul(s float64) Vector2D {
	// TODO: Implement scalar multiplication
	return Vector2D{X: v.X * s, Y: v.Y * s}
}

// Magnitude returns the Euclidean length of the vector: sqrt(x^2 + y^2).
func (v Vector2D) Magnitude() float64 {
	// TODO: Implement magnitude calculation
	return math.Hypot(v.X, v.Y)
}

// MagnitudeSquared returns x^2 + y^2 (cheaper than Magnitude, useful for gravitational denominators).
func (v Vector2D) MagnitudeSquared() float64 {
	// TODO: Implement magnitude squared calculation
	return v.X*v.X + v.Y*v.Y
}

// Normalize returns a unit vector in the same direction, or zero vector if magnitude is 0.
func (v Vector2D) Normalize() Vector2D {
	// TODO: Implement normalization
	mag := v.Magnitude()
	if mag == 0 {
		return Vector2D{}
	}
	return Vector2D{X: v.X / mag, Y: v.Y / mag}
}

// Distance returns Euclidean distance between this vector and another.
func (v Vector2D) Distance(other Vector2D) float64 {
	// TODO: Implement Euclidean distance
	return v.Sub(other).Magnitude()
}
