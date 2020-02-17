//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/maxvoronov/otus-go-calendar/internal/domain"
	"github.com/maxvoronov/otus-go-calendar/internal/logger"
	"github.com/maxvoronov/otus-go-calendar/internal/messaging/amqp"
	"github.com/maxvoronov/otus-go-calendar/internal/service"
	"github.com/maxvoronov/otus-go-calendar/internal/storage/sql"
)

func InitializeScheduler() (*scheduler, error) {
	wire.Build(
		newScheduler,
		sql.InitializeStorage,
		amqp.InitializeAmqp,
		wire.Bind(new(domain.StorageInterface), new(*sql.Storage)),
		wire.Bind(new(domain.MessageBusInterface), new(*amqp.MessageBus)),
		service.NewCalendarService,
		logger.NewLogger,
	)

	return &scheduler{}, nil
}
