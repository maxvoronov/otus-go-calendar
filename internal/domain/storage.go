package domain

import (
	"context"
	"time"
)

// StorageInterface Interface of events storage
type StorageInterface interface {
	GetAll(ctx context.Context) ([]*Event, error)
	GetByID(ctx context.Context, id string) (*Event, error)
	GetByPeriod(ctx context.Context, from, to time.Time) ([]*Event, error)
	GetForNotification(ctx context.Context, from, to time.Time) ([]*Event, error)
	Save(ctx context.Context, event *Event) error
	Remove(ctx context.Context, event *Event) error
}
