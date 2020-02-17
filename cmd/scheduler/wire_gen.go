// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/maxvoronov/otus-go-calendar/internal/logger"
	"github.com/maxvoronov/otus-go-calendar/internal/messaging/amqp"
	"github.com/maxvoronov/otus-go-calendar/internal/service"
	"github.com/maxvoronov/otus-go-calendar/internal/storage/sql"
)

// Injectors from wire.go:

func InitializeScheduler() (*scheduler, error) {
	storage, err := sql.InitializeStorage()
	if err != nil {
		return nil, err
	}
	calendarService := service.NewCalendarService(storage)
	messageBus, err := amqp.InitializeAmqp()
	if err != nil {
		return nil, err
	}
	logrusLogger := logger.NewLogger()
	mainScheduler := newScheduler(calendarService, storage, messageBus, logrusLogger)
	return mainScheduler, nil
}
