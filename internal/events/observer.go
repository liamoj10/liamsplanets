package events

import "github.com/liamoj10/liamsplanets/internal/simulation"

// EventType categorizes simulation events.
type EventType string

const (
	EventTick      EventType = "TICK"
	EventCollision EventType = "COLLISION"
	EventSpawn     EventType = "SPAWN"
	EventStateWait EventType = "PAUSED"
)

// Event carries event data broadcast by the simulation.
type Event struct {
	Type    EventType
	Payload any
}

// Observer defines the subscriber interface for listening to simulation events.
// Refactoring Guru: https://refactoring.guru/design-patterns/observer
//
// Observer allows a subject to notify multiple interested objects about any
// events that happen to the subject without coupling the subject to specific classes.
type Observer interface {
	OnEvent(event Event)
}

// Subject manages observers and broadcasts events.
type Subject interface {
	Subscribe(observer Observer)
	Unsubscribe(observer Observer)
	Notify(event Event)
}

// EventDispatcher is the concrete Subject implementation.
type EventDispatcher struct {
	// TODO: Add observers slice or map and synchronization lock
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{}
}

func (d *EventDispatcher) Subscribe(observer Observer) {
	// TODO: Add observer to the collection
}

func (d *EventDispatcher) Unsubscribe(observer Observer) {
	// TODO: Remove observer from the collection
}

func (d *EventDispatcher) Notify(event Event) {
	// TODO: Iterate over observers and call observer.OnEvent(event)
}

// WebSocketBroadcaster is an example Observer that forwards events to connected clients.
type WebSocketBroadcaster struct {
	// TODO: Hold client connections / channels
}

func (w *WebSocketBroadcaster) OnEvent(event Event) {
	// TODO: Serialize event to JSON and push to client sockets
}

// CollisionLogger is an example Observer that logs planetary collisions.
type CollisionLogger struct{}

func (c *CollisionLogger) OnEvent(event Event) {
	// TODO: If event.Type == EventCollision, log collision details
}

// Ensure Body is imported for payload types
var _ simulation.Body
