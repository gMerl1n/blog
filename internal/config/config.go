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

type ConfigTokens struct {
	JWTsecret       string
	OneDayInSeconds int
	AccessTokenTTL  int
	RefreshTokenTTL int
}

type Config struct {
	ConfigServer *ConfigServer
	ConfigDB     *ConfigDB
	ConfigToken  *ConfigTokens
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
		&ConfigTokens{
			JWTsecret:       viper.GetString("token.jwt_secret"),
			OneDayInSeconds: viper.GetInt("token.one_day_in_seconds"),
			AccessTokenTTL:  viper.GetInt("token.access_tokenTTL"),
			RefreshTokenTTL: viper.GetInt("token.refresh_tokenTTL"),
		},
	}, nil
}

func fetchConfig() error {

	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()

}
