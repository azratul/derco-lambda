package config

var Conf Config

type Config struct {
	JsonPlaceHolder struct {
		Url    string
		ApiKey string
	}
}
