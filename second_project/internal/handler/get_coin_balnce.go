package handler

import (
	"encoding/json"
	"net/http"

	"second_project/api"
	"second_project/internal/tools"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalServerErrorHandler(w)
		return
	}

	var database *tools.Database

	database, err = tools.GetDatabase()

	if err != nil {
		log.Error(err)
		api.InternalServerErrorHandler(w)
		return
	}

	var tokenDetails *tools.LoginDetails

	tokenDetails, err = (*database).GetUserLoginDetails(params.Username)

	if tokenDetails == nil {
		api.InternalServerErrorHandler(w)
		return
	}

	var coinDetails *tools.CoinDetails

	coinDetails, err = (*database).GetUserCoinDetails(params.Username)

	if err != nil {
		log.Error(err)
		api.InternalServerErrorHandler(w)
		return
	}

	var response = api.CoinBalanceResponse{
		StatusCode: http.StatusOK,
		Balance:    coinDetails.Coin,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Error(err)
		api.InternalServerErrorHandler(w)
	}
}
