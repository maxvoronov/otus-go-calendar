package entity

import (
	"github.com/satori/go.uuid"
	"time"
)

type Event struct {
	Id       uuid.UUID
	Title    string
	DateFrom *time.Time
	DateTo   *time.Time
}

func NewEvent(title string, from, to *time.Time) *Event {
	return &Event{
		Id:       uuid.NewV4(),
		Title:    title,
		DateFrom: from,
		DateTo:   to,
	}
}
