package main

import (
	"encoding/json"

	"github.com/maxvoronov/otus-go-calendar/internal/domain"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type notificator struct {
	MessageBus domain.MessageBusInterface
	Logger     *logrus.Logger
}

func newNotificator(bus domain.MessageBusInterface, logger *logrus.Logger) *notificator {
	return &notificator{bus, logger}
}

func (ntf *notificator) Start() error {
	err := ntf.MessageBus.Subscribe("notificator.email", func(delivery amqp.Delivery) {
		event := &domain.Event{}
		if err := json.Unmarshal(delivery.Body, event); err != nil {
			ntf.Logger.Errorf("Failed to parse message: %s", err)
			return
		}

		ntf.sendEmail(event)
	})
	if err != nil {
		return err
	}

	return nil
}

// ToDo: Send email notification
func (ntf *notificator) sendEmail(event *domain.Event) {
	ntf.Logger.Infof("Received event: %s", event)
}
