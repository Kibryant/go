package config

import (
	"first-project/schemas"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	LOGGER := GetLogger("SQLITE: ")

	const DB_PATH = "./db/sqlite.db"

	_, err := os.Stat(DB_PATH)

	if os.IsNotExist(err) {
		LOGGER.Info("sqlite db not found, creating new one")

		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			LOGGER.Errf("sqlite db creation error: %v", err.Error())
			return nil, err
		}

		file, err := os.Create(DB_PATH)

		if err != nil {
			LOGGER.Errf("sqlite db creation error: %v", err.Error())
			return nil, err
		}

		file.Close()
	}

	database, err := gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})

	if err != nil {
		LOGGER.Errf("sqlite opening error: %v", err.Error())
		return nil, err
	}

	LOGGER.Info("sqlite connected")

	err = database.AutoMigrate(&schemas.User{})

	if err != nil {
		LOGGER.Errf("sqlite migration error: %v", err.Error())
		return nil, err
	}

	LOGGER.Info("sqlite migrated")

	return database, nil
}
