package service

import (
	"context"
	"time"

	"github.com/maxvoronov/otus-go-calendar/internal/domain"
)

// CalendarService struct
type CalendarService struct {
	storage domain.StorageInterface
}

// NewCalendarService Init new calendar service instance
func NewCalendarService(s domain.StorageInterface) *CalendarService {
	return &CalendarService{storage: s}
}

// GetEvents Return list of all events
func (cal *CalendarService) GetEvents() ([]*domain.Event, error) {
	return cal.storage.GetAll(context.Background())
}

// GetEventsByPeriod Return list of events by period
func (cal *CalendarService) GetEventsByPeriod(from, to time.Time) ([]*domain.Event, error) {
	return cal.storage.GetByPeriod(context.Background(), from, to)
}

// GetEventsForNotification Return list of events for notifications
func (cal *CalendarService) GetEventsForNotification(from, to time.Time) ([]*domain.Event, error) {
	return cal.storage.GetForNotification(context.Background(), from, to)
}

// GetEventByID Return event by ID
func (cal *CalendarService) GetEventByID(id string) (*domain.Event, error) {
	return cal.storage.GetByID(context.Background(), id)
}

// CreateEvent Create new event and save it to storage
func (cal *CalendarService) CreateEvent(title string, from, to time.Time) (*domain.Event, error) {
	event := domain.NewEvent(title, from, to)
	if err := cal.storage.Save(context.Background(), event); err != nil {
		return nil, err
	}

	return event, nil
}

// UpdateEvent Update existing event
func (cal *CalendarService) UpdateEvent(event *domain.Event) error {
	return cal.storage.Save(context.Background(), event)
}

// RemoveEvent Remove existing event
func (cal *CalendarService) RemoveEvent(event *domain.Event) error {
	return cal.storage.Remove(context.Background(), event)
}
