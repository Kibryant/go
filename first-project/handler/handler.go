package handler

import (
	"first-project/config"

	"gorm.io/gorm"
)

var (
	LOGGER   *config.Logger
	DATABASE *gorm.DB
)

func InitializeHandler() {
	LOGGER = config.GetLogger("handler")
	DATABASE = config.GetDatabase()
}
