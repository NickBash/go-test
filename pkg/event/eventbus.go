package event

const (
	LinkVisitedEvent = "link.visited"
)

type Event struct {
	Type string
	Data any
}

type EventBus struct {
	bus chan Event
}

func NewEventBus() *EventBus {
	return &EventBus{
		bus: make(chan Event),
	}
}

func (eb *EventBus) Public(event Event) {
	eb.bus <- event
}

func (eb *EventBus) Subscribe() <-chan Event {
	return eb.bus
}
