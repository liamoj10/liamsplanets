package commands

import "github.com/liamoj10/liamsplanets/internal/simulation"

// Command declares an interface for executing operations on the simulation.
// Refactoring Guru: https://refactoring.guru/design-patterns/command
//
// The Command pattern turns a request into a stand-alone object containing
// all information about the request. This transformation lets you pass requests
// as arguments, delay or queue execution, and support undoable operations.
type Command interface {
	Execute(engine *simulation.Engine) error
	Undo(engine *simulation.Engine) error
}

// SpawnBodyCommand adds a new celestial body to the simulation.
type SpawnBodyCommand struct {
	Body simulation.Body
}

func (c *SpawnBodyCommand) Execute(engine *simulation.Engine) error {
	// TODO: Add body to the simulation engine
	return nil
}

func (c *SpawnBodyCommand) Undo(engine *simulation.Engine) error {
	// TODO: Remove the spawned body from the engine by name or ID
	return nil
}

// PauseCommand toggles or pauses simulation time.
type PauseCommand struct {
	wasPaused bool
}

func (c *PauseCommand) Execute(engine *simulation.Engine) error {
	// TODO: Record previous pause state and pause engine
	return nil
}

func (c *PauseCommand) Undo(engine *simulation.Engine) error {
	// TODO: Restore previous pause state
	return nil
}

// SetTimeScaleCommand alters the speed of simulation time.
type SetTimeScaleCommand struct {
	NewScale float64
	OldScale float64
}

func (c *SetTimeScaleCommand) Execute(engine *simulation.Engine) error {
	// TODO: Record old scale and set new scale on engine
	return nil
}

func (c *SetTimeScaleCommand) Undo(engine *simulation.Engine) error {
	// TODO: Restore old scale on engine
	return nil
}

// CommandHistory acts as the Invoker/History tracker for undo/redo capabilities.
type CommandHistory struct {
	history []Command
}

func (h *CommandHistory) Push(cmd Command) {
	// TODO: Append executed command to history stack
}

func (h *CommandHistory) Pop() (Command, bool) {
	// TODO: Pop latest command off history stack for undo
	return nil, false
}
