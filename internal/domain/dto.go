package domain

// CelestialBodyDTO represents the serialized data format delivered to the frontend UI.
// Maps cleanly from our rich domain CelestialBody to the UI's expected schema.
type CelestialBodyDTO struct {
	Name   string  `json:"name"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	VX     float64 `json:"vx"`
	VY     float64 `json:"vy"`
	Radius float64 `json:"radius"`
	Color  string  `json:"color"`
	Mass   float64 `json:"mass,omitempty"`
}

// ToDTO maps a domain CelestialBody into a lightweight frontend DTO.
func (b *CelestialBody) ToDTO() CelestialBodyDTO {
	// TODO: Map domain model fields to DTO
	return CelestialBodyDTO{
		Name:   b.Name,
		X:      b.Position.X,
		Y:      b.Position.Y,
		VX:     b.Velocity.X,
		VY:     b.Velocity.Y,
		Radius: b.Radius,
		Color:  b.Color,
		Mass:   b.Mass,
	}
}

// FromDTO constructs a domain CelestialBody from a frontend DTO.
func FromDTO(dto CelestialBodyDTO, id string, class BodyClassification) *CelestialBody {
	// TODO: Construct domain entity from DTO
	return &CelestialBody{
		ID:             id,
		Name:           dto.Name,
		Classification: class,
		Position:       Vector2D{X: dto.X, Y: dto.Y},
		Velocity:       Vector2D{X: dto.VX, Y: dto.VY},
		Radius:         dto.Radius,
		Color:          dto.Color,
		Mass:           dto.Mass,
	}
}
