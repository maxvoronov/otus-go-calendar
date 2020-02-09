// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package grpc

import (
	"github.com/maxvoronov/otus-go-calendar/internal/logger"
	"github.com/maxvoronov/otus-go-calendar/internal/service"
	"github.com/maxvoronov/otus-go-calendar/internal/storage/sql"
)

// Injectors from wire.go:

func InitializeServer() (*server, error) {
	storage, err := sql.InitializeStorage()
	if err != nil {
		return nil, err
	}
	calendarService := service.NewCalendarService(storage)
	logrusLogger := logger.NewLogger()
	grpcServer := newServer(calendarService, storage, logrusLogger)
	return grpcServer, nil
}
