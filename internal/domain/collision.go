package domain

// CollisionResult encapsulates what happens when two bodies collide.
type CollisionResult struct {
	DestroyedBodyIDs []string
	Debris           []*CelestialBody
}

// CollisionHandler defines the strategy for resolving celestial impacts.
type CollisionHandler interface {
	Resolve(b1, b2 *CelestialBody) CollisionResult
}

// ShatterCollisionHandler implements the destruction & debris shattering mechanic.
// When two bodies collide, they break into smaller asteroid debris fragments.
type ShatterCollisionHandler struct {
	DebrisCount int     // Number of fragments generated per collision
	MinMass     float64 // Minimum mass threshold to generate debris
}

func NewShatterCollisionHandler(debrisCount int) *ShatterCollisionHandler {
	return &ShatterCollisionHandler{
		DebrisCount: debrisCount,
		MinMass:     1.0,
	}
}

// Resolve destroys the colliding bodies and disperses debris fragments.
func (h *ShatterCollisionHandler) Resolve(b1, b2 *CelestialBody) CollisionResult {
	// TODO: Implement shattering collision physics:
	// 1. Total mass M = b1.Mass + b2.Mass
	// 2. Center of mass position: P_cm = (b1.Pos*m1 + b2.Pos*m2) / M
	// 3. Center of mass velocity: V_cm = (b1.Vel*m1 + b2.Vel*m2) / M
	// 4. Generate h.DebrisCount smaller asteroid fragments scattered radially from P_cm
	//    with random dispersion velocities added to V_cm.
	// 5. Return CollisionResult with DestroyedBodyIDs: [b1.ID, b2.ID] and Debris list.
	return CollisionResult{
		DestroyedBodyIDs: []string{b1.ID, b2.ID},
		Debris:           []*CelestialBody{},
	}
}
