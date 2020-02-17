//+build wireinject

package amqp

import (
	"github.com/google/wire"
	"github.com/maxvoronov/otus-go-calendar/internal/logger"
)

func InitializeAmqp() (*MessageBus, error) {
	wire.Build(NewMessageBus, CreateConfigFromEnvironment, logger.NewLogger)

	return &MessageBus{}, nil
}
