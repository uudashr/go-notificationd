package notificationd

// Publishing service.
type Publishing struct {
	Name             string
	TrackerStore     TrackerStore
	Publisher        Publisher
	EventStore       EventStore
	EventUnmarshaler EventUnmarshaler
}

// PublishNotifications publish the notifications.
func (p *Publishing) PublishNotifications() error {
	tracker, err := p.TrackerStore.TrackerOf(p.Name)
	if err != nil {
		return err
	}

	notifications, err := p.listUnpublishNotificationsSince(tracker.MostRecentPublishedNotificationID())
	if err != nil {
		return err
	}

	for _, n := range notifications {
		if err = p.Publisher.Publish(n); err != nil {
			return err
		}
	}
	return p.TrackerStore.TrackMostRecentNotification(tracker, notifications)
}

func (p *Publishing) listUnpublishNotificationsSince(id int64) ([]*Notification, error) {
	events, err := p.EventStore.StoredEventSince(id)
	if err != nil {
		return nil, err
	}

	return toNotifications(events, p.EventUnmarshaler)
}

// EventUnmarshaler unmarshals event.
type EventUnmarshaler interface {
	UnmarshalEvent(name string, version int, data []byte) (interface{}, error)
}

func toNotifications(events []*StoredEvent, unmarshaler EventUnmarshaler) ([]*Notification, error) {
	notifications := make([]*Notification, len(events))
	for i, e := range events {
		event, err := unmarshaler.UnmarshalEvent(e.Name, e.Version, e.Body)
		if err != nil {
			return nil, err
		}

		notifications[i] = &Notification{
			ID:          e.ID,
			Name:        e.Name,
			Event:       event,
			Version:     e.Version,
			OccuredTime: e.OccuredTime,
		}
	}
	return notifications, nil
}

// Publisher publish notification.
type Publisher interface {
	Publish(*Notification) error
}
