//+build wireinject

package api

import (
	"github.com/google/wire"
	"github.com/maxvoronov/otus-go-calendar/internal/domain"
	"github.com/maxvoronov/otus-go-calendar/internal/logger"
	"github.com/maxvoronov/otus-go-calendar/storage/sql"
)

func InitializeServer() (*server, error) {
	wire.Build(
		newServer,
		sql.InitializeStorage,
		wire.Bind(new(domain.StorageInterface), new(*sql.Storage)),
		logger.NewLogger,
	)

	return nil, nil
}
