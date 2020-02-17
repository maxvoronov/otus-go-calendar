package amqp

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

// MessageBus struct
type MessageBus struct {
	Client  *amqp.Connection
	Channel *amqp.Channel
	Queue   amqp.Queue
	Logger  *logrus.Logger
}

// NewMessageBus Create new message broker (rabbitmq)
func NewMessageBus(conf *Config, logger *logrus.Logger) (*MessageBus, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", conf.User, conf.Password, conf.Host, conf.Port))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to AMQP broker")
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare AMQP channel")
	}

	queue, err := channel.QueueDeclare(
		"events_queue", // name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to declare AMQP queue")
	}

	return &MessageBus{conn, channel, queue, logger}, nil
}

// Publish Send message to bus
func (bus *MessageBus) Publish(message []byte) error {
	err := bus.Channel.Publish(
		"",             // exchange
		bus.Queue.Name, // routing key
		false,          // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         message,
		})
	if err != nil {
		return errors.Wrap(err, "failed to publish event")
	}

	return nil
}

// Subscribe to message bus queue
func (bus *MessageBus) Subscribe(consumerName string, handlerFunc func(amqp.Delivery)) error {
	messages, err := bus.Channel.Consume(
		bus.Queue.Name, // queue
		consumerName,   // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	if err != nil {
		return errors.Wrap(err, "failed to init AMQP consumer")
	}

	waitingChan := make(chan struct{})
	go func() {
		for msgDelivery := range messages {
			bus.Logger.Debugf("Received a message: %s", msgDelivery.Body)
			handlerFunc(msgDelivery)
		}
	}()
	<-waitingChan

	return nil
}
