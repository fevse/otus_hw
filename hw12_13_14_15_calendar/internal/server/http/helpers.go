package internalhttp

import (
	"strconv"
	"time"

	"github.com/fevse/otus_hw/hw12_13_14_15_calendar/internal/storage"
)


type EventStr struct {
	ID          string `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Date        string `json:"date" db:"date"`
	Duration    string `json:"duration" db:"duration"`
	Description string `json:"description" db:"description"`
	UserID      string `json:"userid" db:"userid"`
	Reminder    string `json:"reminder" db:"reminder"`
}

type ID struct {
	ID string `json:"id"`
}

type Date struct {
	Date string `json:"date"`
}

func (e *EventStr) ParseEvent() (storage.Event, error) {
	event := storage.Event{}

	id, err := strconv.Atoi(e.ID)
	if err != nil {
		return event, err
	}
	date, err := time.Parse(time.DateTime, e.Date)
	if err != nil {
		return event, err
	}
	dur, err := strconv.Atoi(e.Duration)
	if err != nil {
		return event, err
	}
	uid, err := strconv.Atoi(e.UserID)
	if err != nil {
		return event, err
	}
	rem, err := time.Parse(time.DateTime, e.Reminder)
	if err != nil {
		return event, err
	}
	event = storage.Event{ID: int64(id), Title: e.Title, Date: date, Duration: int64(dur), Description: e.Description, UserID: int64(uid), Reminder: rem}
	return event, nil
}
