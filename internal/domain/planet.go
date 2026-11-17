package domain

// BodyClassification categorizes the astrological / physical nature of the body.
type BodyClassification string

const (
	ClassStar        BodyClassification = "STAR"
	ClassTerrestrial BodyClassification = "TERRESTRIAL"
	ClassGasGiant    BodyClassification = "GAS_GIANT"
	ClassMoon        BodyClassification = "MOON"
	ClassAsteroid    BodyClassification = "ASTEROID"
)

// CelestialBody is the rich domain entity representing any celestial object
// in the simulation or eventual game.
type CelestialBody struct {
	ID             string             `json:"id"`
	Name           string             `json:"name"`
	Classification BodyClassification `json:"classification"`

	// Spatial kinematics
	Position     Vector2D `json:"position"`
	Velocity     Vector2D `json:"velocity"`
	Acceleration Vector2D `json:"acceleration"`

	// Physical properties
	Mass   float64 `json:"mass"`
	Radius float64 `json:"radius"`

	// Visual styling
	Color      string `json:"color"`
	TrailColor string `json:"trail_color,omitempty"`

	// Orbital hierarchy (e.g. Moon orbits Earth, Earth orbits Sol)
	ParentBodyID *string `json:"parent_body_id,omitempty"`

	// TODO: Future Game Mechanics (Prototype pattern / decorator pattern)
	// AtmosphereDensity float64
	// Resources         map[string]float64
	// Population        int64
	// HasColony         bool
}

// Speed returns the current scalar speed (magnitude of velocity vector).
func (b *CelestialBody) Speed() float64 {
	// TODO: Implement scalar speed calculation
	return b.Velocity.Magnitude()
}

// KineticEnergy calculates 0.5 * m * v^2.
func (b *CelestialBody) KineticEnergy() float64 {
	// TODO: Implement kinetic energy calculation
	return 0.5 * b.Mass * b.Velocity.MagnitudeSquared()
}

// CollidesWith checks if this body physically overlaps with another.
// dist(p1, p2) <= (r1 + r2)
func (b *CelestialBody) CollidesWith(other *CelestialBody) bool {
	// TODO: Implement collision detection
	return b.Position.Distance(other.Position) <= (b.Radius + other.Radius)
}

// GravitationalForce calculates the gravitational pull vector exerted BY this body ON another body.
// F = G * (m1 * m2 / (r^2 + softening^2)) * direction_unit_vector
func (b *CelestialBody) GravitationalForce(other *CelestialBody, G, softening float64) Vector2D {
	// TODO: Implement gravitational force vector calculation
	// 1. Compute delta vector from other -> b: delta = b.Position.Sub(other.Position)
	// 2. Compute distSq = delta.MagnitudeSquared() + softening*softening
	// 3. Compute magnitude = (G * b.Mass * other.Mass) / distSq
	// 4. Return delta.Normalize().Mul(magnitude)
	return Vector2D{}
}

// OrbitalVelocity calculates the theoretical circular orbit speed around a primary mass M at radius r:
// v = sqrt(G * M / r)
func OrbitalVelocity(primaryMass, orbitalRadius, G float64) float64 {
	// TODO: Implement circular orbital speed formula
	return 0.0
}
