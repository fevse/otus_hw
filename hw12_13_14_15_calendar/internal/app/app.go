package app

import (
	"time"

	"github.com/fevse/otus_hw/hw12_13_14_15_calendar/internal/storage"
)

type App struct { // TODO
	Logger  Logger
	Storage Storage
}

type Logger interface { // TODO
	Info(string)
	Error(string)
}

type Storage interface { // TODO
	CreateEvent(*storage.Event) error
	UpdateEvent(int64, *storage.Event) error
	DeleteEvent(int64) error
	List() (storage.Events, error)
	ListOfEventsDay(time.Time) (storage.Events, error)
	ListOfEventsWeek(time.Time) (storage.Events, error)
	ListOfEventsMonth(time.Time) (storage.Events, error)
}

func New(logger Logger, storage Storage) *App {
	return &App{
		Logger:  logger,
		Storage: storage,
	}
}
