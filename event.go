package notificationd

import (
	"time"

	eventd "github.com/uudashr/go-eventd"
)

// EventStore stores the event.
type EventStore interface {
	Append(eventd.Event) (*StoredEvent, error)
	StoredEventSince(int64) ([]*StoredEvent, error)
}

// StoredEvent is the stored event.
type StoredEvent struct {
	ID          int64
	Name        string
	Body        []byte
	Version     int
	OccuredTime time.Time
}
