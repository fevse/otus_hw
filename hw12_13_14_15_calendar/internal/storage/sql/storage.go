package sqlstorage

import (
	"context"
	"errors"
	"time"

	"github.com/fevse/otus_hw/hw12_13_14_15_calendar/internal/config"
	"github.com/fevse/otus_hw/hw12_13_14_15_calendar/internal/storage"
	_ "github.com/jackc/pgx/stdlib" // driver
	"github.com/jmoiron/sqlx"
)

type Storage struct { // TODO
	conf config.Config
	db *sqlx.DB
}

func New(conf config.Config) *Storage {
	return &Storage{conf: conf}
}

func (s *Storage) Connect(_ context.Context, dsn string) (err error) {
	// TODO
	s.db, err = sqlx.Open("pgx", dsn)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Close(_ context.Context) error {
	// TODO
	if s.db != nil {
		return s.db.Close()
	}
	return errors.New("DB is not open")
}

func (s *Storage) CreateEvent(event *storage.Event) error {
	err := s.Connect(context.Background(), s.conf.DB.DSN)
	if err != nil {
		return err
	}
	defer s.Close(context.Background())
	_, err = s.db.Exec(`INSERT INTO events(id, title, date, duration, description, userid, reminder)
		VALUES($1, $2, $3, $4, $5, $6, $7);`,
		event.ID, event.Title, event.Date, event.Duration, event.Description, event.UserID, event.Reminder)
	return err
}

func (s *Storage) UpdateEvent(id int64, event *storage.Event) error {
	err := s.Connect(context.Background(), s.conf.DB.DSN)
	if err != nil {
		return err
	}	
	defer s.Close(context.Background())
	_, err = s.db.Exec(`UPDATE events 
		SET title = $2, date=$3, duration=$4, description=$5, userid=$6, reminder=$7
		WHERE id=$1;`, id, event.Title, event.Date, event.Duration, event.Description, event.UserID, event.Reminder)
	return err
}

func (s *Storage) DeleteEvent(id int64) error {
	err := s.Connect(context.Background(), s.conf.DB.DSN)
	if err != nil {
		return err
	}
	defer s.Close(context.Background())
	_, err = s.db.Exec(`DELETE FROM events WHERE id=$1`, id)
	return err
}

func (s *Storage) List() (events storage.Events, err error) {
	data := make([]storage.Event, 0)
	events = make(storage.Events)
	err = s.Connect(context.Background(), s.conf.DB.DSN)
	if err != nil {
		return
	}
	defer s.Close(context.Background())
	err = s.db.Select(&data, `SELECT * FROM events`)
	if err != nil {
		return 
	}
	for _, v := range data {
		events[v.ID] = &v
	}
	return
}

func (s *Storage) ListOfEventsDay(date time.Time) (events storage.Events, err error) {
	data := make([]storage.Event, 0)
	events = make(storage.Events)
	err = s.Connect(context.Background(), s.conf.DB.DSN)
	if err != nil {
		return
	}
	defer s.Close(context.Background())
	err = s.db.Select(&data, `SELECT * FROM events WHERE date = $1`, date)
	for _, v := range data {
		events[v.ID] = &v
	}
	return
}

func (s *Storage) ListOfEventsWeek(date time.Time) (events storage.Events, err error) {
	data := make([]storage.Event, 0)
	events = make(storage.Events)
	err = s.Connect(context.Background(), s.conf.DB.DSN)
	if err != nil {
		return
	}
	defer s.Close(context.Background())
	err = s.db.Select(&data, `SELECT * FROM events WHERE date >= $1 AND date <= $2`, date, date.AddDate(0, 0, 7))
	for _, v := range data {
		events[v.ID] = &v
	}
	return
}

func (s *Storage) ListOfEventsMonth(date time.Time) (events storage.Events, err error) {
	data := make([]storage.Event, 0)
	events = make(storage.Events)
	err = s.Connect(context.Background(), s.conf.DB.DSN)
	if err != nil {
		return
	}
	defer s.Close(context.Background())
	err = s.db.Select(&data, `SELECT * FROM events WHERE date >= $1 AND date <= $2`, date, date.AddDate(0, 1, 0))
	for _, v := range data {
		events[v.ID] = &v
	}
	return
}
