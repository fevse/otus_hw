package memorystorage

import (
	"fmt"
	"testing"
	"time"

	"github.com/fevse/otus_hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	// TODO
	stor := New()
	t.Run("test create event", func(t *testing.T) {
		err := stor.CreateEvent(&storage.Event{
			ID:          1,
			Title:       "Check",
			Date:        time.Now(),
			Duration:    4,
			Description: "test data",
			UserID:      1,
			Reminder:    time.Now().Add(999),
		})
		if err != nil {
			fmt.Println(err)
		}
		err = stor.CreateEvent(&storage.Event{
			ID:          2,
			Title:       "Hello",
			Date:        time.Now().Add(128),
			Duration:    8,
			Description: "some data",
			UserID:      2,
			Reminder:    time.Now().Add(222),
		})
		if err != nil {
			fmt.Println(err)
		}
		err = stor.CreateEvent(&storage.Event{
			ID:          3,
			Title:       "Test",
			Date:        time.Now().Add(256),
			Duration:    16,
			Description: "another data",
			UserID:      3,
			Reminder:    time.Now().Add(555),
		})
		if err != nil {
			fmt.Println(err)
		}
		require.Equal(t, "Check", stor.events[1].Title)
		require.Equal(t, 3, len(stor.events))
	})
	t.Run("test update event", func(t *testing.T) {
		err := stor.UpdateEvent(1, &storage.Event{
			ID:          1,
			Title:       "Update",
			Date:        time.Now(),
			Duration:    5,
			Description: "test data",
			UserID:      1,
			Reminder:    time.Now().Add(999),
		})
		if err != nil {
			fmt.Println(err)
		}
		require.Equal(t, "Update", stor.events[1].Title)
		require.Equal(t, 3, len(stor.events))
	})
	t.Run("test delete event", func(t *testing.T) {
		err := stor.DeleteEvent(1)
		if err != nil {
			fmt.Println(err)
		}
		require.Equal(t, 2, len(stor.events))
	})
}
