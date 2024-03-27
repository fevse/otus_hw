package sqlstorage

import (
	"context"
	"time"

	"github.com/fevse/otus_hw/hw12_13_14_15_calendar/internal/storage"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Storage struct { // TODO
	db *sqlx.DB
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Connect(ctx context.Context, dsn string) (err error) {
	// TODO
	s.db, err = sqlx.Open("pgx", dsn)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	// TODO
	return s.db.Close()
}

func (s *Storage) CreateEvent (event *storage.Event) error {
	_, err := s.db.Exec(`INSERT INTO events(
		id, title, date, duration, description, user_id, reminder)
		VALUES($1, $2, $3, $4, $5, $6, $7);`,
		event.ID, event.Title, event.Date, event.Duration, event.Description, event.UserID, event.Reminder)
	return err
}

func (s *Storage) UpdateEvent (id int64, event *storage.Event) error {
	_, err := s.db.Exec(`UPDATE events 
		SET title = $2, date=$3, duration=$4, description=$5, user_id=$6, reminder=$7)
		WHERE id=$1	;`,
		id, event.Title, event.Date, event.Duration, event.Description, event.UserID, event.Reminder)
	return err
}

func (s *Storage) DeleteEvent (id int64) error {
	_, err := s.db.Exec(`DELETE FROM events WHERE id=$1`, id)
	return err
}

func (s *Storage) ListOfEventsDay (date time.Time) (events []*storage.Event, err error) {
	err = s.db.Select(&events, `SELECT * FROM events WHERE date = $1`, date)
	return
}

func (s *Storage) ListOfEventsWeek (date time.Time) (events []*storage.Event, err error) {
	err = s.db.Select(&events, `SELECT * FROM events WHERE date > $1 AND date < $2`, date, date.AddDate(0, 0, 7))
	return
}

func (s *Storage) ListOfEventsMonth (date time.Time) (events []*storage.Event, err error) {
	err = s.db.Select(&events, `SELECT * FROM events WHERE date > $1 AND date < $2`, date, date.AddDate(0, 1, 0))
	return
}