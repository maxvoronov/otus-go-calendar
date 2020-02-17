//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/maxvoronov/otus-go-calendar/internal/domain"
	"github.com/maxvoronov/otus-go-calendar/internal/logger"
	"github.com/maxvoronov/otus-go-calendar/internal/messaging/amqp"
)

func InitializeNotificator() (*notificator, error) {
	wire.Build(
		newNotificator,
		amqp.InitializeAmqp,
		wire.Bind(new(domain.MessageBusInterface), new(*amqp.MessageBus)),
		logger.NewLogger,
	)

	return &notificator{}, nil
}
