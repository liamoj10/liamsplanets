package factory

import "github.com/liamoj10/liamsplanets/internal/simulation"

// BodyType categorizes celestial body templates.
type BodyType string

const (
	TypeStar        BodyType = "STAR"
	TypeTerrestrial BodyType = "TERRESTRIAL"
	TypeGasGiant    BodyType = "GAS_GIANT"
	TypeMoon        BodyType = "MOON"
)

// CelestialBodyFactory defines the Factory Method interface.
// Refactoring Guru: https://refactoring.guru/design-patterns/factory-method
//
// Factory Method provides an interface for creating objects in a superclass,
// but allows subclasses or factory functions to alter the type of objects that will be created.
type CelestialBodyFactory interface {
	CreateBody(name string, x, y, vx, vy float64) simulation.Body
}

// StarFactory creates massive luminous central bodies.
type StarFactory struct{}

func (f *StarFactory) CreateBody(name string, x, y, vx, vy float64) simulation.Body {
	// TODO: Return simulation.Body with star defaults (high mass, warm color like #facc15, large radius)
	return simulation.Body{
		Name:   name,
		X:      x,
		Y:      y,
		VX:     vx,
		VY:     vy,
		Radius: 18,
		Color:  "#facc15",
		Mass:   50000,
	}
}

// TerrestrialFactory creates rocky planets (like Earth/Mars).
type TerrestrialFactory struct{}

func (f *TerrestrialFactory) CreateBody(name string, x, y, vx, vy float64) simulation.Body {
	// TODO: Return simulation.Body with rocky planet defaults (moderate mass, custom color, medium radius)
	return simulation.Body{
		Name:   name,
		X:      x,
		Y:      y,
		VX:     vx,
		VY:     vy,
		Radius: 6,
		Color:  "#38bdf8",
		Mass:   10,
	}
}

// SolarSystemBuilder defines the Builder pattern interface.
// Refactoring Guru: https://refactoring.guru/design-patterns/builder
//
// Builder is a creational design pattern that lets you produce complex objects
// step by step, allowing you to produce different types and representations
// of an object using the same construction code.
type SolarSystemBuilder interface {
	Reset() SolarSystemBuilder
	SetCentralStar(name string, mass, radius float64) SolarSystemBuilder
	AddPlanet(name string, orbitalRadius, orbitalSpeed float64, bodyType BodyType) SolarSystemBuilder
	AddMoon(planetName, moonName string, orbitalRadius, orbitalSpeed float64) SolarSystemBuilder
	Build() []simulation.Body
}

// ConcreteSolarSystemBuilder implements SolarSystemBuilder step by step.
type ConcreteSolarSystemBuilder struct {
	bodies []simulation.Body
}

func NewSolarSystemBuilder() *ConcreteSolarSystemBuilder {
	return &ConcreteSolarSystemBuilder{bodies: make([]simulation.Body, 0)}
}

func (b *ConcreteSolarSystemBuilder) Reset() SolarSystemBuilder {
	b.bodies = make([]simulation.Body, 0)
	return b
}

func (b *ConcreteSolarSystemBuilder) SetCentralStar(name string, mass, radius float64) SolarSystemBuilder {
	// TODO: Add central star positioned at (0, 0) with zero velocity
	return b
}

func (b *ConcreteSolarSystemBuilder) AddPlanet(name string, orbitalRadius, orbitalSpeed float64, bodyType BodyType) SolarSystemBuilder {
	// TODO: Position planet at (orbitalRadius, 0) with tangent velocity (0, orbitalSpeed)
	return b
}

func (b *ConcreteSolarSystemBuilder) AddMoon(planetName, moonName string, orbitalRadius, orbitalSpeed float64) SolarSystemBuilder {
	// TODO: Find parent planet and compute relative moon position & combined velocity
	return b
}

func (b *ConcreteSolarSystemBuilder) Build() []simulation.Body {
	// TODO: Return constructed planetary bodies slice
	return b.bodies
}
