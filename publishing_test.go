package notificationd_test

import (
	mq "github.com/uudashr/go-mq"
	"github.com/uudashr/go-notificationd"
)

func ExamplePublishing() {
	var eventStore notificationd.EventStore
	var trackerStore notificationd.TrackerStore
	var unmarshaler notificationd.EventUnmarshaler
	var pub mq.Publisher
	// ...

	publisher := &notificationd.MQPublisher{
		Topic:     "events",
		Publisher: pub,
	}

	publishing := &notificationd.Publishing{
		Name:             "Default",
		EventStore:       eventStore,
		TrackerStore:     trackerStore,
		EventUnmarshaler: unmarshaler,
		Publisher:        publisher,
	}

	if err := publishing.PublishNotifications(); err != nil {
		// TODO: handle the error
	}
}
