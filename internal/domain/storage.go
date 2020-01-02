package domain

import "time"

// StorageInterface Interface of events storage
type StorageInterface interface {
	GetAll() ([]*Event, error)
	GetByID(id string) (*Event, error)
	GetByPeriod(from, to time.Time) ([]*Event, error)
	Save(event *Event) error
	Remove(event *Event) error
}
