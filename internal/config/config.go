package config

import (
	"os"

	"github.com/spf13/viper"
)

type ConfigServer struct {
	Port         string
	ReadTimeOut  int
	WriteTimeOut int
}

type ConfigDB struct {
	Host     string
	Port     string
	User     string
	Password string
	NameDB   string
	SSLMode  string
}

type Config struct {
	ConfigServer *ConfigServer
	ConfigDB     *ConfigDB
}

func NewConfig() *Config {
	return &Config{
		&ConfigServer{
			Port:         viper.GetString("server.port"),
			ReadTimeOut:  viper.GetInt("server.read_timeout"),
			WriteTimeOut: viper.GetInt("server.write_timeout"),
		},
		&ConfigDB{
			User:     os.Getenv("POSTGRES_USER"),
			NameDB:   os.Getenv("POSTGRES_DB"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			SSLMode:  os.Getenv("SSLMODE"),
		},
	}
}
