# Go Backend Architecture & Design Patterns Learning Guide

Based on [Refactoring.Guru Design Patterns](https://refactoring.guru/design-patterns).

This project has been scaffolded with clean Go interfaces, doc comments, and `TODO` learning exercises for you to implement each pattern step by step.

---

## 0. Core Domain Model
- **Path**: [internal/domain/planet.go](file:///Users/Liam/Desktop/liamsplanets/internal/domain/planet.go), [internal/domain/vector.go](file:///Users/Liam/Desktop/liamsplanets/internal/domain/vector.go), [internal/domain/dto.go](file:///Users/Liam/Desktop/liamsplanets/internal/domain/dto.go)
- **Concept**: Rich Domain Model entity representing celestial objects (stars, planets, gas giants, moons, asteroids) with vector mathematics and physical properties.
- **Your Exercises**:
  1. `Vector2D`: Vector addition, subtraction, scalar multiplication, and distance in [internal/domain/vector.go](file:///Users/Liam/Desktop/liamsplanets/internal/domain/vector.go).
  2. `CelestialBody.GravitationalForce`: Calculate gravitational force vector exerted on another body in [internal/domain/planet.go](file:///Users/Liam/Desktop/liamsplanets/internal/domain/planet.go#L63).
  3. `OrbitalVelocity`: Calculate circular orbital velocity around a primary mass: $v = \sqrt{\frac{G \cdot M}{r}}$ in [internal/domain/planet.go](file:///Users/Liam/Desktop/liamsplanets/internal/domain/planet.go#L73).
  4. `ToDTO` / `FromDTO`: Data Transfer Object mapping between the rich domain model and the lightweight UI JSON format in [internal/domain/dto.go](file:///Users/Liam/Desktop/liamsplanets/internal/domain/dto.go).
  5. Run unit tests anytime via `go test ./internal/domain/...`.

---

## 0.1 Decorator Pattern (Structural)
- **Path**: [internal/domain/decorator.go](file:///Users/Liam/Desktop/liamsplanets/internal/domain/decorator.go)
- **Refactoring Guru**: [Decorator Pattern](https://refactoring.guru/design-patterns/decorator)
- **Concept**: Dynamically attach game mechanics (atmosphere, mining resources, population/colonies) to a base `CelestialBody` without bloating the core physics struct.
- **Your Exercises**:
  1. `AtmosphericDecorator.ComputeTemperature`: Calculate equilibrium surface temperature based on stellar luminosity and orbital distance.
  2. `ResourceBearingDecorator.Mine`: Extract finite resources (Iron, Water Ice, Minerals) from deposits.
  3. `ColonizedDecorator.GrowPopulation`: Model population expansion and infrastructure growth.

---

## 0.2 Collision Shatter System (Behavioral / Strategy)
- **Path**: [internal/domain/collision.go](file:///Users/Liam/Desktop/liamsplanets/internal/domain/collision.go)
- **Concept**: Destructive impact resolution where colliding bodies shatter into smaller asteroid debris fragments using conservation of momentum ($P_{cm}$).
- **Your Exercises**:
  1. `ShatterCollisionHandler.Resolve`: Calculate center of mass ($P_{cm}$) and center of mass velocity ($V_{cm}$), destroy parent bodies, and spawn scattered asteroid fragments.

---

## 1. Strategy Pattern (Behavioral)
- **Path**: [internal/physics/strategy.go](file:///Users/Liam/Desktop/liamsplanets/internal/physics/strategy.go)
- **Refactoring Guru**: [Strategy Pattern](https://refactoring.guru/design-patterns/strategy)
- **Concept**: Define a family of physics algorithms, encapsulate each one in its own struct (`Integrator`), and make them interchangeable at runtime.
- **Your Exercises**:
  1. `EulerIntegrator.Step(...)`: Standard first-order Euler integration.
  2. `VerletIntegrator.Step(...)`: Velocity Verlet integration (conserves orbital energy).
  3. `RungeKutta4Integrator.Step(...)`: 4th-order Runge-Kutta numerical integration.

---

## 2. Observer Pattern (Behavioral)
- **Path**: [internal/events/observer.go](file:///Users/Liam/Desktop/liamsplanets/internal/events/observer.go)
- **Refactoring Guru**: [Observer Pattern](https://refactoring.guru/design-patterns/observer)
- **Concept**: Decouple the simulation engine from consumers (WebSockets, loggers, collision detectors) by broadcasting events via a `Subject` and `Observer` interface.
- **Your Exercises**:
  1. Complete `EventDispatcher.Subscribe`, `Unsubscribe`, and `Notify`.
  2. Implement `WebSocketBroadcaster` and `CollisionLogger` observers.

---

## 3. Command Pattern (Behavioral)
- **Path**: [internal/commands/command.go](file:///Users/Liam/Desktop/liamsplanets/internal/commands/command.go)
- **Refactoring Guru**: [Command Pattern](https://refactoring.guru/design-patterns/command)
- **Concept**: Turn simulation requests (spawning planets, pausing, setting simulation speed) into stand-alone objects with `Execute` and `Undo` methods.
- **Your Exercises**:
  1. Implement `Execute` and `Undo` on `SpawnBodyCommand`.
  2. Implement `CommandHistory` stack to enable time-rewind / undo functionality.

---

## 4. State Pattern (Behavioral)
- **Path**: [internal/state/state.go](file:///Users/Liam/Desktop/liamsplanets/internal/state/state.go)
- **Refactoring Guru**: [State Pattern](https://refactoring.guru/design-patterns/state)
- **Concept**: Allow the simulation engine to alter its behavior dynamically when its internal state changes (`RunningState`, `PausedState`) without giant `if/switch` blocks.
- **Your Exercises**:
  1. Implement transitions between `Play()` and `Pause()`.
  2. Implement tick handling where `RunningState` advances physics while `PausedState` idles.

---

## 5. Factory Method & Builder (Creational)
- **Path**: [internal/factory/builder.go](file:///Users/Liam/Desktop/liamsplanets/internal/factory/builder.go)
- **Refactoring Guru**:
  - [Factory Method](https://refactoring.guru/design-patterns/factory-method)
  - [Builder Pattern](https://refactoring.guru/design-patterns/builder)
- **Concept**:
  - **Factory Method**: Use `StarFactory` and `TerrestrialFactory` to create pre-configured bodies.
  - **Builder**: Construct entire planetary systems step-by-step using a fluent interface (`SetCentralStar(...) -> AddPlanet(...) -> AddMoon(...) -> Build()`).
- **Your Exercises**:
  1. Implement `ConcreteSolarSystemBuilder.AddPlanet(...)` to compute tangent orbital velocities.
  2. Implement `AddMoon(...)` to position orbiting moons relative to their parent planet.
