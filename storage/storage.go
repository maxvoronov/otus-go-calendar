package storage

import "github.com/maxvoronov/otus-go-calendar/entity"

type StorageInterface interface {
	GetAll() ([]*entity.Event, error)
	GetById(id string) (*entity.Event, error)
	Save(event *entity.Event) error
	Remove(event *entity.Event) error
}
