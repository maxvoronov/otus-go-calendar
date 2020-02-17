package domain

import "github.com/streadway/amqp"

// MessageBusInterface Interface of message bus
type MessageBusInterface interface {
	Publish(message []byte) error
	Subscribe(consumerName string, handlerFunc func(amqp.Delivery)) error
}
