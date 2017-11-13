package notificationd

import (
	"encoding/json"
	"time"
)

// Notification represents the notification.
type Notification struct {
	ID          int64
	Name        string
	Event       interface{}
	Version     int
	OccuredTime time.Time
}

// Versioned interface for versioned event.
// Event with no version will be assumed as version 1.
type Versioned interface {
	Version() int
}

// MarshalJSON implements the json.Marshaler interface.
func (n Notification) MarshalJSON() ([]byte, error) {
	event, err := json.Marshal(n.Event)
	if err != nil {
		return nil, err
	}

	version := 1
	if v, ok := n.Event.(Versioned); ok {
		version = v.Version()
	}

	return json.Marshal(jsonizeNotification{
		ID:          n.ID,
		Name:        n.Name,
		Event:       event,
		Version:     version,
		OccuredTime: n.OccuredTime.Format(time.RFC3339Nano),
	})
}

// ReadNotification from data as Reader.
func ReadNotification(data []byte) (*NotificationReader, error) {
	var v jsonizeNotification
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}

	occuredTime, err := time.Parse(time.RFC3339Nano, v.OccuredTime)
	if err != nil {
		return nil, err
	}

	return &NotificationReader{
		id:          v.ID,
		name:        v.Name,
		eventData:   v.Event,
		version:     v.Version,
		occuredTime: occuredTime,
	}, nil
}

// NotificationReader is notification reader.
type NotificationReader struct {
	id          int64
	name        string
	eventData   []byte
	version     int
	occuredTime time.Time
}

// ID of the notification.
func (r *NotificationReader) ID() int64 {
	return r.id
}

// Name of the notification.
func (r *NotificationReader) Name() string {
	return r.name
}

// ReadEvent from notification.
func (r *NotificationReader) ReadEvent(event interface{}) error {
	return json.Unmarshal(r.eventData, event)
}

// Version of the notification.
func (r *NotificationReader) Version() int {
	return r.version
}

// OccuredTime of the notification.
func (r *NotificationReader) OccuredTime() time.Time {
	return r.occuredTime
}

type jsonizeNotification struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	Event       json.RawMessage `json:"event"`
	Version     int             `json:"version"`
	OccuredTime string          `json:"occuredTime"`
}
