package notificationd

import (
	"encoding/json"

	exchange "github.com/uudashr/go-exchange"
	mq "github.com/uudashr/go-mq"
)

// MQPublisher is publisher implementation to messaging queue.
type MQPublisher struct {
	Topic     string
	Publisher mq.Publisher
}

// Publish the notification.
func (p *MQPublisher) Publish(n *Notification) error {
	msg, err := toMessage(n)
	if err != nil {
		return err
	}
	return exchange.Publish(p.Publisher, p.Topic, msg)
}

func toMessage(n *Notification) (*exchange.Message, error) {
	body, err := json.Marshal(n)
	if err != nil {
		return nil, err
	}

	return &exchange.Message{
		Type: n.Name,
		Body: body,
	}, nil
}
