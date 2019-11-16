package inmemory

import (
	"github.com/maxvoronov/otus-go-calendar/entity"
	"sync"
)

type InMemoryStorage struct {
	sync.Mutex
	events map[string]*entity.Event
}

func NewInMemoryStorage() *InMemoryStorage {
	eventStorage := make(map[string]*entity.Event)
	return &InMemoryStorage{events: eventStorage}
}

func (storage *InMemoryStorage) GetAll() ([]*entity.Event, error) {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	result := make([]*entity.Event, 0, len(storage.events))
	for _, event := range storage.events {
		result = append(result, event)
	}

	return result, nil
}

func (storage *InMemoryStorage) GetById(id string) (*entity.Event, error) {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	if event, ok := storage.events[id]; ok {
		return event, nil
	}

	return nil, nil
}

func (storage *InMemoryStorage) Save(event *entity.Event) error {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	storage.events[event.Id.String()] = event

	return nil
}

func (storage *InMemoryStorage) Remove(event *entity.Event) error {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	delete(storage.events, event.Id.String())

	return nil
}
