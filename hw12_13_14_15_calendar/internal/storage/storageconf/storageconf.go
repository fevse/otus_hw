package storageconf

import (
	"github.com/fevse/otus_hw/hw12_13_14_15_calendar/internal/app"
	"github.com/fevse/otus_hw/hw12_13_14_15_calendar/internal/config"
	memorystorage "github.com/fevse/otus_hw/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/fevse/otus_hw/hw12_13_14_15_calendar/internal/storage/sql"
)

func ChangeStorage(config config.Config, logg app.Logger) app.Storage {
	var storage app.Storage
	logg.Info("DB is used " + config.DB.Type)
	switch config.DB.Type {
	case "memorystorage":
		storage = memorystorage.New()
	case "sql":
		storage = sqlstorage.New()
	}
	return storage
}
