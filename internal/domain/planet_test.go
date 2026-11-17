package domain

import (
	"math"
	"testing"
)

func TestVector2D_Add(t *testing.T) {
	v1 := Vector2D{X: 1, Y: 2}
	v2 := Vector2D{X: 3, Y: 4}
	res := v1.Add(v2)

	if res.X != 4 || res.Y != 6 {
		t.Errorf("expected (4, 6), got (%v, %v)", res.X, res.Y)
	}
}

func TestVector2D_Magnitude(t *testing.T) {
	v := Vector2D{X: 3, Y: 4}
	mag := v.Magnitude()

	if math.Abs(mag-5.0) > 1e-6 {
		t.Errorf("expected magnitude 5, got %v", mag)
	}
}

func TestCelestialBody_CollidesWith(t *testing.T) {
	b1 := &CelestialBody{
		Position: Vector2D{X: 0, Y: 0},
		Radius:   10,
	}
	b2 := &CelestialBody{
		Position: Vector2D{X: 15, Y: 0},
		Radius:   6,
	}

	if !b1.CollidesWith(b2) {
		t.Errorf("expected bodies with distance 15 and combined radii 16 to collide")
	}

	b2.Position = Vector2D{X: 20, Y: 0}
	if b1.CollidesWith(b2) {
		t.Errorf("expected bodies with distance 20 and combined radii 16 NOT to collide")
	}
}
