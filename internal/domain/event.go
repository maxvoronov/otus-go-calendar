package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Event struct
type Event struct {
	ID       uuid.UUID
	Title    string
	DateFrom time.Time
	DateTo   time.Time
}

// NewEvent Create and return new event
func NewEvent(title string, from, to time.Time) *Event {
	return &Event{
		ID:       uuid.NewV4(),
		Title:    title,
		DateFrom: from,
		DateTo:   to,
	}
}
