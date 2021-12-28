package event

type EventPlugin interface {
	EventType() string
	Start() error
}
