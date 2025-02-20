package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type ConfigServer struct {
	Port         string
	LogLevel     int
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

func NewConfig() (*Config, error) {

	if err := fetchConfig(); err != nil {
		fmt.Printf("error initialization config %s", err.Error())
		return nil, err
	}

	return &Config{
		&ConfigServer{
			Port:         viper.GetString("server.port"),
			LogLevel:     viper.GetInt("server.log_level"),
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
	}, nil
}

func fetchConfig() error {

	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()

}
