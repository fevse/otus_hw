package memorystorage

import (
	"sync"
	"time"

	"github.com/fevse/otus_hw/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	// TODO
	mu     sync.RWMutex
	events storage.Events
}

func New() *Storage {
	events := make(storage.Events)
	return &Storage{
		mu:     sync.RWMutex{},
		events: events,
	}
}

func (s *Storage) CreateEvent(event *storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.events[event.ID]; !ok {
		s.events[event.ID] = event
		return nil
	}

	return storage.ErrorEventAlreadyExist
}

func (s *Storage) UpdateEvent(id int64, event *storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.events[id]; !ok {
		return storage.ErrorEventNotFound
	}

	s.events[id] = event
	return nil
}

func (s *Storage) DeleteEvent(id int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.events[id]; !ok {
		return storage.ErrorEventNotFound
	}

	delete(s.events, id)
	return nil
}

func (s *Storage) List() (events storage.Events) {
	return s.events
}

func (s *Storage) ListOfEventsDay(date time.Time) (events storage.Events, err error) {
	events = make(storage.Events)
	y1, m1, d1 := date.Date()
	for id, event := range s.events {
		y2, m2, d2 := event.Date.Date()
		if y1 == y2 &&
			m1 == m2 &&
			d1 == d2 {
				events[id] = event
		}
	}
	return events, nil
}

func (s *Storage) ListOfEventsWeek(date time.Time) (events storage.Events, err error) {
	events = make(storage.Events)
	y1, m1, d1 := date.Date()
	for id, event := range s.events {
		y2, m2, d2 := event.Date.Date()
		if y1 == y2 &&
			m1 == m2 &&
			d1 <= d2 &&
			d1+7 >= d2 {
				events[id] = event
		}
	}
	return events, nil
}

func (s *Storage) ListOfEventsMonth(date time.Time) (events storage.Events, err error) {
	events = make(storage.Events)
	y1, m1, _ := date.Date()
	for id, event := range s.events {
		y2, m2, _ := event.Date.Date()
		if y1 == y2 &&
			m1 == m2 {
			events[id] = event
		}
	}
	return events, nil
}
