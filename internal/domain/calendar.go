package domain

import "time"

// CalendarInterface Interface of calendar
type CalendarInterface interface {
	GetEvents() ([]*Event, error)
	GetEventByID(id string) (*Event, error)
	AddEvent(title string, from, to *time.Time) (*Event, error)
	UpdateEvent(event *Event) error
	RemoveEvent(event *Event) error
}

// Calendar struct
type Calendar struct {
	Title   string
	storage StorageInterface
}

// NewCalendar Init new calendar instance
func NewCalendar(title string, s StorageInterface) *Calendar {
	return &Calendar{Title: title, storage: s}
}

// GetEvents Return list of all events
func (cal *Calendar) GetEvents() ([]*Event, error) {
	return cal.storage.GetAll()
}

// GetEventByID Return event by ID
func (cal *Calendar) GetEventByID(id string) (*Event, error) {
	return cal.storage.GetByID(id)
}

// CreateEvent Create new event and save it to storage
func (cal *Calendar) CreateEvent(title string, from, to time.Time) (*Event, error) {
	event := NewEvent(title, from, to)
	if err := cal.storage.Save(event); err != nil {
		return nil, err
	}

	return event, nil
}

// UpdateEvent Update existing event
func (cal *Calendar) UpdateEvent(event *Event) error {
	return cal.storage.Save(event)
}

// RemoveEvent Remove existing event
func (cal *Calendar) RemoveEvent(event *Event) error {
	return cal.storage.Remove(event)
}
