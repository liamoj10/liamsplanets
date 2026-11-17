# Liam's Planets: Engineering Roadmap & Linear Issue Plan

Team Prefix: **SOL** (from `SOL-02`, `SOL-03`)  
Linear Workspace: [linear.app/liamsplanets](https://linear.app/liamsplanets/welcome)

---

## Phase 1: Core Physics & Numerical Integration (Strategy Pattern)
*Focus: Get stable orbital mechanics running in Go.*

### LIA-11: Implement Vector2D Kinematics & Operations
- **Status**: [LIA-11](https://linear.app/liamsplanets/issue/LIA-11/implement-vector2d-kinematics-and-operations)
- **Size**: `S` | **Milestone**: Milestone 1: MVP Physics | **Labels**: `physics`, `backend`
- **Component**: `internal/domain`
- **Description**: Complete `Vector2D` math operations in [internal/domain/vector.go](file:///Users/Liam/Desktop/liamsplanets/internal/domain/vector.go) (`Add`, `Sub`, `Mul`, `Magnitude`, `MagnitudeSquared`, `Normalize`, `Distance`).

### LIA-12: Implement Strategy Pattern for Physics Integrators
- **Status**: [LIA-12](https://linear.app/liamsplanets/issue/LIA-12/implement-strategy-pattern-for-physics-integrators-euler-verlet-rk4)
- **Size**: `M` | **Milestone**: Milestone 1: MVP Physics | **Labels**: `design-pattern`, `physics`, `backend`
- **Pattern**: [Strategy Pattern](https://refactoring.guru/design-patterns/strategy)
- **Component**: `internal/physics`
- **Description**: Implement the `Integrator` interface in [internal/physics/strategy.go](file:///Users/Liam/Desktop/liamsplanets/internal/physics/strategy.go) (Euler, Velocity-Verlet, Runge-Kutta 4).

---

## Phase 2: Domain Modeling & Celestial Mechanics

### LIA-13: Implement Gravitational Force & Circular Orbital Velocity
- **Status**: [LIA-13](https://linear.app/liamsplanets/issue/LIA-13/implement-gravitational-force-and-circular-orbital-velocity)
- **Size**: `S` | **Milestone**: Milestone 1: MVP Physics | **Labels**: `physics`, `backend`
- **Component**: `internal/domain`
- **Description**: In [internal/domain/planet.go](file:///Users/Liam/Desktop/liamsplanets/internal/domain/planet.go), implement gravitational force and circular orbital velocity formulas.

### LIA-14: Implement Shatter Collision Detection & Debris Resolution
- **Status**: [LIA-14](https://linear.app/liamsplanets/issue/LIA-14/implement-shatter-collision-detection-and-debris-resolution)
- **Size**: `M` | **Milestone**: Milestone 1: MVP Physics | **Labels**: `physics`, `gameplay`, `backend`
- **Component**: `internal/domain`
- **Description**: Complete `ShatterCollisionHandler.Resolve` in [internal/domain/collision.go](file:///Users/Liam/Desktop/liamsplanets/internal/domain/collision.go).

---

## Phase 3: Game Mechanics & Extensibility (Decorator Pattern)

### LIA-15: Implement Planet Decorators (Atmosphere, Resources, Colonies)
- **Status**: [LIA-15](https://linear.app/liamsplanets/issue/LIA-15/implement-decorator-pattern-for-planet-attributes-atmosphere-resources)
- **Size**: `L` | **Milestone**: Milestone 2: Game Mechanics | **Labels**: `design-pattern`, `gameplay`, `backend`
- **Pattern**: [Decorator Pattern](https://refactoring.guru/design-patterns/decorator)
- **Component**: `internal/domain`
- **Description**: Complete decorator wrappers in [internal/domain/decorator.go](file:///Users/Liam/Desktop/liamsplanets/internal/domain/decorator.go).

---

## Phase 4: System Creation & Control (Factory, Builder, Command)

### LIA-16: Implement SolarSystemBuilder & Factories
- **Status**: [LIA-16](https://linear.app/liamsplanets/issue/LIA-16/implement-solarsystembuilder-and-factory-methods)
- **Size**: `M` | **Milestone**: Milestone 3: Control & Architecture | **Labels**: `design-pattern`, `backend`
- **Pattern**: [Builder](https://refactoring.guru/design-patterns/builder) & [Factory Method](https://refactoring.guru/design-patterns/factory-method)
- **Component**: `internal/factory`

### LIA-17: Implement Command Pattern & Undo History
- **Status**: [LIA-17](https://linear.app/liamsplanets/issue/LIA-17/implement-command-pattern-and-undo-history-stack)
- **Size**: `M` | **Milestone**: Milestone 3: Control & Architecture | **Labels**: `design-pattern`, `gameplay`, `backend`
- **Pattern**: [Command Pattern](https://refactoring.guru/design-patterns/command)
- **Component**: `internal/commands`

---

## Phase 5: State Machine & Event Dispatcher (State & Observer)

### LIA-18: Implement Simulation State Machine
- **Status**: [LIA-18](https://linear.app/liamsplanets/issue/LIA-18/implement-simulation-state-machine-running-paused)
- **Size**: `S` | **Milestone**: Milestone 3: Control & Architecture | **Labels**: `design-pattern`, `backend`
- **Pattern**: [State Pattern](https://refactoring.guru/design-patterns/state)
- **Component**: `internal/state`

### LIA-19: Implement Observer Event Dispatcher & WebSocket Stream
- **Status**: [LIA-19](https://linear.app/liamsplanets/issue/LIA-19/implement-observer-event-dispatcher-and-websocket-stream)
- **Size**: `L` | **Milestone**: Milestone 3: Control & Architecture | **Labels**: `design-pattern`, `backend`
- **Pattern**: [Observer Pattern](https://refactoring.guru/design-patterns/observer)
- **Component**: `internal/events`, `internal/transport`
