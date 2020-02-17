package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// EventStatusNew Status for all new events
const EventStatusNew = "new"

// EventStatusNotified Status for processed events
const EventStatusNotified = "notified"

// Event struct
type Event struct {
	ID       uuid.UUID
	Title    string
	Status   string
	DateFrom time.Time
	DateTo   time.Time
}

// NewEvent Create and return new event
func NewEvent(title string, from, to time.Time) *Event {
	return &Event{
		ID:       uuid.NewV4(),
		Title:    title,
		Status:   EventStatusNew,
		DateFrom: from,
		DateTo:   to,
	}
}
