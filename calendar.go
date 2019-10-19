package calendar

import (
	"github.com/maxvoronov/otus-go-calendar/entity"
	"github.com/maxvoronov/otus-go-calendar/storage"
	"time"
)

type CalendarInterface interface {
	GetEvents() ([]*entity.Event, error)
	GetEventById(id string) (*entity.Event, error)
	AddEvent(title string, from, to *time.Time) (*entity.Event, error)
	UpdateEvent(event *entity.Event) error
	RemoveEvent(event *entity.Event) error
}

type Calendar struct {
	Title   string
	storage storage.StorageInterface
}

func NewCalendar(title string, s storage.StorageInterface) *Calendar {
	return &Calendar{Title: title, storage: s}
}

func (cal *Calendar) GetEvents() ([]*entity.Event, error) {
	return cal.storage.GetAll()
}

func (cal *Calendar) GetEventById(id string) (*entity.Event, error) {
	return cal.storage.GetById(id)
}

func (cal *Calendar) CreateEvent(title string, from, to *time.Time) (*entity.Event, error) {
	event := entity.NewEvent(title, from, to)
	if err := cal.storage.Save(event); err != nil {
		return nil, err
	}

	return event, nil
}

func (cal *Calendar) UpdateEvent(event *entity.Event) error {
	return cal.storage.Save(event)
}

func (cal *Calendar) RemoveEvent(event *entity.Event) error {
	return cal.storage.Remove(event)
}
