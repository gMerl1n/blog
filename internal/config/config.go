package config

import "github.com/spf13/viper"

type Config struct {
	Port         string
	ReadTimeOut  int
	WriteTimeOut int
}

func NewConfig() *Config {
	return &Config{
		Port:         viper.GetString("server.port"),
		ReadTimeOut:  viper.GetInt("server.read_timeout"),
		WriteTimeOut: viper.GetInt("server.write_timeout"),
	}
}
