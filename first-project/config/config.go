package config

import "gorm.io/gorm"

var (
	DATABASE *gorm.DB
	LOGGER   *Logger
)

func Initialize() error {
	var err error

	LOGGER = GetLogger("CONFIG: ")

	DATABASE, err = InitializeSQLite()

	if err != nil {
		LOGGER.Errf("sqlite initialization error: %v", err.Error())
		return err
	}

	return nil
}

func GetLogger(prefix string) *Logger {
	LOGGER = NewLogger(prefix)
	return LOGGER
}

func GetDatabase() *gorm.DB {
	return DATABASE
}
