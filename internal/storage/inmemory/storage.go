package inmemory

import (
	"context"
	"sync"
	"time"

	"github.com/maxvoronov/otus-go-calendar/internal/domain"
)

// Storage struct
type Storage struct {
	sync.Mutex
	events map[string]*domain.Event
}

// NewStorage Create new in-memory storage
func NewStorage() *Storage {
	eventStorage := make(map[string]*domain.Event)
	return &Storage{events: eventStorage}
}

// GetAll Return list of all events
func (storage *Storage) GetAll(_ context.Context) ([]*domain.Event, error) {
	storage.Lock()
	defer storage.Unlock()
	result := make([]*domain.Event, 0, len(storage.events))
	for _, event := range storage.events {
		result = append(result, event)
	}

	return result, nil
}

// GetByID Return event by ID
func (storage *Storage) GetByID(_ context.Context, id string) (*domain.Event, error) {
	storage.Lock()
	defer storage.Unlock()
	if event, ok := storage.events[id]; ok {
		return event, nil
	}

	return nil, nil
}

// GetByPeriod Return list of events by period
func (storage *Storage) GetByPeriod(_ context.Context, from, to time.Time) ([]*domain.Event, error) {
	storage.Lock()
	defer storage.Unlock()
	result := make([]*domain.Event, 0)

	for _, event := range storage.events {
		if event.DateFrom.Before(to) && event.DateTo.After(from) {
			result = append(result, event)
		}
	}

	return result, nil
}

// GetForNotification Return list of events for notifications
func (storage *Storage) GetForNotification(_ context.Context, from, to time.Time) ([]*domain.Event, error) {
	storage.Lock()
	defer storage.Unlock()
	result := make([]*domain.Event, 0)

	for _, event := range storage.events {
		if event.Status == domain.EventStatusNew && event.DateFrom.Before(to) && event.DateFrom.After(from) {
			result = append(result, event)
		}
	}

	return result, nil
}

// Save Create or update event in storage
func (storage *Storage) Save(_ context.Context, event *domain.Event) error {
	storage.Lock()
	defer storage.Unlock()
	storage.events[event.ID.String()] = event

	return nil
}

// Remove event from storage
func (storage *Storage) Remove(_ context.Context, event *domain.Event) error {
	storage.Lock()
	defer storage.Unlock()
	delete(storage.events, event.ID.String())

	return nil
}
