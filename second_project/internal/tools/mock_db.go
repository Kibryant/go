package tools

type mockDb struct{}

var MOCK_LOGIN_DETAILS = map[string]LoginDetails{
	"test_user": {
		Token:    "1234567890",
		Username: "test_user",
	},
	"test_user_2": {
		Token:    "1234567890",
		Username: "test_user_2",
	},
}

var MOCK_COIN_DETAILS = map[string]CoinDetails{
	"test_user": {
		Coin:     100,
		Username: "test_user",
	},
	"test_user_2": {
		Coin:     200,
		Username: "test_user_2",
	},
}

func (db *mockDb) GetUserLoginDetails(username string) (*LoginDetails, error) {
	var loginDetails = LoginDetails{}

	loginDetails, ok := MOCK_LOGIN_DETAILS[username]

	if !ok {
		return nil, nil
	}

	return &loginDetails, nil
}

func (db *mockDb) GetUserCoinDetails(username string) (*CoinDetails, error) {
	var coinDetails = CoinDetails{}

	coinDetails, ok := MOCK_COIN_DETAILS[username]

	if !ok {
		return nil, nil
	}

	return &coinDetails, nil
}

func (db *mockDb) SetupDatabase() error {
	return nil
}
