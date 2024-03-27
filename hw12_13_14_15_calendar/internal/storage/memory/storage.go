package memorystorage

import (
	"sync"
	"time"

	"github.com/fevse/otus_hw/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	// TODO
	mu sync.RWMutex 
	events storage.Events
}

func New() *Storage {
	events := make(storage.Events)
	return &Storage{
		mu: sync.RWMutex{},
		events: events,
	}
}

func (s *Storage) CreateEvent (event *storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.events[event.ID]; !ok {
		s.events[event.ID] = event
	}

	return storage.ErrorEventAlreadyExist
}

func (s *Storage) UpdateEvent (id int64, event *storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	if _, ok := s.events[id]; !ok {
		return storage.ErrorEventNotFound
	}

	s.events[id] = event
	return nil
}

func (s *Storage) DeleteEvent (id int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.events[id]; !ok {
		return storage.ErrorEventNotFound
	}

	delete(s.events, id)
	return nil
}

func (s *Storage) ListOfEventsDay (date time.Time) storage.Events {
	list := make(storage.Events)
	for id, event := range s.events {
		if event.Date.Year() == date.Year() && 
		event.Date.Month() == date.Month() &&
		event.Date.Day() == date.Day() {
			list[id] = event
		}
	}
	return list
}

func (s *Storage) ListOfEventsWeek (date time.Time) storage.Events {
	list := make(storage.Events)
	for id, event := range s.events {
		if event.Date.Year() == date.Year() && 
		event.Date.Month() == date.Month() &&
		event.Date.Day() >= date.Day() &&
		event.Date.Day() <= date.Day()+7 {
			list[id] = event
		}
	}
	return list
}

func (s *Storage) ListOfEventsMonth (date time.Time) storage.Events {
	list := make(storage.Events)
	for id, event := range s.events {
		if event.Date.Year() == date.Year() && 
		event.Date.Month() == date.Month() {
			list[id] = event
		}
	}
	return list
}
