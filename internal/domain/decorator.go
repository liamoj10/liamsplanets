package domain

// BodyComponent defines the Decorator interface for celestial objects.
// Refactoring Guru: https://refactoring.guru/design-patterns/decorator
//
// The Decorator pattern lets you attach new behaviors and data to objects
// by placing these objects inside special wrapper objects that contain the behaviors.
type BodyComponent interface {
	GetID() string
	GetName() string
	GetClassification() BodyClassification
	GetPosition() Vector2D
	GetVelocity() Vector2D
	GetMass() float64
	GetRadius() float64
	GetColor() string
	GetBaseBody() *CelestialBody
}

// Ensure CelestialBody satisfies BodyComponent.
var _ BodyComponent = (*CelestialBody)(nil)

func (b *CelestialBody) GetID() string                         { return b.ID }
func (b *CelestialBody) GetName() string                       { return b.Name }
func (b *CelestialBody) GetClassification() BodyClassification { return b.Classification }
func (b *CelestialBody) GetPosition() Vector2D                 { return b.Position }
func (b *CelestialBody) GetVelocity() Vector2D                 { return b.Velocity }
func (b *CelestialBody) GetMass() float64                      { return b.Mass }
func (b *CelestialBody) GetRadius() float64                    { return b.Radius }
func (b *CelestialBody) GetColor() string                      { return b.Color }
func (b *CelestialBody) GetBaseBody() *CelestialBody           { return b }

// BaseBodyDecorator is the base struct forwarding all calls to the wrapped component.
type BaseBodyDecorator struct {
	Wrapped BodyComponent
}

func (d *BaseBodyDecorator) GetID() string                         { return d.Wrapped.GetID() }
func (d *BaseBodyDecorator) GetName() string                       { return d.Wrapped.GetName() }
func (d *BaseBodyDecorator) GetClassification() BodyClassification { return d.Wrapped.GetClassification() }
func (d *BaseBodyDecorator) GetPosition() Vector2D                 { return d.Wrapped.GetPosition() }
func (d *BaseBodyDecorator) GetVelocity() Vector2D                 { return d.Wrapped.GetVelocity() }
func (d *BaseBodyDecorator) GetMass() float64                      { return d.Wrapped.GetMass() }
func (d *BaseBodyDecorator) GetRadius() float64                    { return d.Wrapped.GetRadius() }
func (d *BaseBodyDecorator) GetColor() string                      { return d.Wrapped.GetColor() }
func (d *BaseBodyDecorator) GetBaseBody() *CelestialBody           { return d.Wrapped.GetBaseBody() }

// --- Concrete Decorators ---

// AtmosphericDecorator adds atmosphere and climate telemetry.
type AtmosphericDecorator struct {
	BaseBodyDecorator
	SurfacePressure    float64 // in atmospheres (atm)
	SurfaceTemperature float64 // in Kelvin
	AtmosphereGas      string  // e.g. "N2-O2", "CO2", "H2-He"
	IsHabitable        bool
}

// ComputeTemperature estimates surface temperature based on stellar luminosity and distance.
func (a *AtmosphericDecorator) ComputeTemperature(starDistance, starLuminosity float64) float64 {
	// TODO: Calculate equilibrium temperature: T = (L_star / (16 * pi * sigma * d^2))^(1/4)
	return a.SurfaceTemperature
}

// ResourceBearingDecorator adds mineable game resources.
type ResourceBearingDecorator struct {
	BaseBodyDecorator
	// Resources maps resource type (Iron, WaterIce, RareMinerals, Hydrogen) to abundance units
	Resources map[string]float64
}

// Mine extracts an amount of a specified resource from the body.
func (r *ResourceBearingDecorator) Mine(resource string, amount float64) float64 {
	// TODO: Decrement resource abundance and return amount extracted
	return 0.0
}

// ColonizedDecorator adds human/alien settlements and infrastructure.
type ColonizedDecorator struct {
	BaseBodyDecorator
	Population          int64
	ColonyName          string
	InfrastructureLevel int
}

// GrowPopulation increases colony count based on food/resources.
func (c *ColonizedDecorator) GrowPopulation(rate float64) {
	// TODO: Increase population by growth rate
}
