package main

import (
	"first-project/config"
	"first-project/router"
)

var (
	LOGGER *config.Logger
)

func main() {
	LOGGER = config.GetLogger("MAIN: ")

	err := config.Initialize()

	if err != nil {
		LOGGER.Err(err.Error())
		return
	}

	router.Initialize()
}
