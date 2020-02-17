package main

import (
	"encoding/json"
	"time"

	"github.com/maxvoronov/otus-go-calendar/internal/domain"
	"github.com/maxvoronov/otus-go-calendar/internal/service"
	"github.com/sirupsen/logrus"
)

type scheduler struct {
	Calendar   *service.CalendarService
	Storage    domain.StorageInterface
	MessageBus domain.MessageBusInterface
	Logger     *logrus.Logger
}

func newScheduler(
	calendarSvc *service.CalendarService,
	storage domain.StorageInterface,
	bus domain.MessageBusInterface,
	logger *logrus.Logger,
) *scheduler {
	return &scheduler{calendarSvc, storage, bus, logger}
}

func (sch *scheduler) Start() {
	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {
		events, err := sch.Calendar.GetEventsForNotification(time.Now(), time.Now().Add(3*time.Minute))
		if err != nil {
			sch.Logger.Error(err)
		}

		for _, event := range events {
			msg, err := json.Marshal(event)
			if err != nil {
				sch.Logger.Errorf("Failed to encode event: %s", err)
				continue
			}

			if err := sch.MessageBus.Publish(msg); err != nil {
				sch.Logger.Errorf("Failed to publish event: %s", err)
				continue
			}

			event.Status = domain.EventStatusNotified
			if err := sch.Calendar.UpdateEvent(event); err != nil {
				sch.Logger.Errorf("Failed to update event status: %s", err)
			}
		}
	}
}
