package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	Token    string
	Username string
}

type CoinDetails struct {
	Coin     int64
	Username string
}

type Database interface {
	GetUserLoginDetails(username string) (*LoginDetails, error)
	GetUserCoinDetails(username string) (*CoinDetails, error)
	SetupDatabase() error
}

func GetDatabase() (*Database, error) {
	var database Database = &mockDb{}

	var err error = database.SetupDatabase()

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}
