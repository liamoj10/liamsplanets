package state

// SimulationState defines the State pattern interface.
// Refactoring Guru: https://refactoring.guru/design-patterns/state
//
// The State pattern lets an object alter its behavior when its internal state
// changes. It appears as if the object changed its class.
type SimulationState interface {
	// Name returns the descriptive name of the state (Running, Paused, etc.)
	Name() string

	// Step handles a tick progression. Returns the state to transition to, or nil if no transition.
	Step(context StateContext, dt float64) SimulationState

	// Play transitions towards running.
	Play(context StateContext) SimulationState

	// Pause transitions towards paused.
	Pause(context StateContext) SimulationState
}

// StateContext provides the interface that states can use to interact with the engine.
type StateContext interface {
	AdvancePhysics(dt float64)
	BroadcastTick()
}

// RunningState represents active simulation execution.
type RunningState struct{}

func (s *RunningState) Name() string { return "RUNNING" }

func (s *RunningState) Step(ctx StateContext, dt float64) SimulationState {
	// TODO: Call ctx.AdvancePhysics(dt), ctx.BroadcastTick(), return s (remain in running)
	return s
}

func (s *RunningState) Play(ctx StateContext) SimulationState {
	// Already playing, return current state
	return s
}

func (s *RunningState) Pause(ctx StateContext) SimulationState {
	// TODO: Transition to &PausedState{}
	return s
}

// PausedState represents halted simulation time (inspect, orbit planning, etc.).
type PausedState struct{}

func (s *PausedState) Name() string { return "PAUSED" }

func (s *PausedState) Step(ctx StateContext, dt float64) SimulationState {
	// In paused state, physics is not advanced; return s
	return s
}

func (s *PausedState) Play(ctx StateContext) SimulationState {
	// TODO: Transition to &RunningState{}
	return s
}

func (s *PausedState) Pause(ctx StateContext) SimulationState {
	// Already paused
	return s
}
