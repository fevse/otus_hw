package storage

import (
	"errors"
	"time"
)

var (
	ErrorEventAlreadyExist = errors.New("Event with this ID already exists")
	ErrorEventNotFound = errors.New("Event with this ID not found")
)

type Event struct {
	ID    		int64
	Title 		string
	Date 		time.Time
	Duration 	time.Duration
	Description string
	UserID 		int64
	Reminder 	time.Time
}

type Events map[int64]*Event
