package notificationd

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// TrackerStore stores the notification tracker.
type TrackerStore interface {
	TrackerOf(name string) (*Tracker, error)
	TrackMostRecentPublishedNotification(*Tracker, []*Notification) error
}

// Tracker is the notification tracker.
type Tracker struct {
	id                                TrackerID
	name                              string
	mostRecentPublishedNotificationID int64
}

// ID of the tracker.
func (t Tracker) ID() TrackerID {
	return t.id
}

// Name of the tracker.
func (t Tracker) Name() string {
	return t.name
}

// MostRecentPublishedNotificationID on the tracker.
func (t Tracker) MostRecentPublishedNotificationID() int64 {
	return t.mostRecentPublishedNotificationID
}

// SetMostRecentPublishedNotificationID of the tracker.
func (t *Tracker) SetMostRecentPublishedNotificationID(id int64) {
	t.mostRecentPublishedNotificationID = id
}

// WithID duplicate the tracker with assigned ID.
func (t *Tracker) WithID(id TrackerID) (*Tracker, error) {
	if !id.OK() {
		return nil, errors.New("notificationd: invalid id")
	}

	return &Tracker{
		id:   id,
		name: t.Name(),
		mostRecentPublishedNotificationID: t.MostRecentPublishedNotificationID(),
	}, nil
}

// NewTracker constructs new tracker.
func NewTracker(id TrackerID, name string, mostRecentPublishedNotificationID int64) (*Tracker, error) {
	if !id.OK() {
		return nil, errors.New("notificationd: invalid id")
	}

	trimmedName := strings.TrimSpace(name)
	if trimmedName == "" {
		return nil, errors.New("notificationd: empty name")
	}

	return &Tracker{
		id:   id,
		name: name,
		mostRecentPublishedNotificationID: mostRecentPublishedNotificationID,
	}, nil
}

// NewEmptyTracker constructs new empty tracker.
func NewEmptyTracker(name string) (*Tracker, error) {
	trimmedName := strings.TrimSpace(name)
	if trimmedName == "" {
		return nil, errors.New("notificationd: empty name")
	}

	return &Tracker{
		name: name,
	}, nil
}

// UninitializedTrackerID is the uninitialized id.
var UninitializedTrackerID = TrackerID("")

// TrackerID is the tracker identifier.
type TrackerID string

// OK check the validity.
func (id TrackerID) OK() bool {
	return strings.TrimSpace(string(id)) != ""
}

// Scan implements the sql.Scanner interface.
func (id *TrackerID) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		*id = TrackerID(string(v))
		return nil
	case string:
		*id = TrackerID(v)
		return nil
	default:
		return fmt.Errorf("notificationd: unable to scan tracker id %#v", src)
	}
}

// Value implements the driver.Valuer interface.
func (id TrackerID) Value() (driver.Value, error) {
	return string(id), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (id *TrackerID) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	*id = TrackerID(s)
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (id TrackerID) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(id))
}
