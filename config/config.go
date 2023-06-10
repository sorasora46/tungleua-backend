package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseDSN        string
	AccessTokenSecret  string
	RefreshTokenSecret string
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read configuration file: %v", err)
	}

	AppConfig = &Config{
		DatabaseDSN:        viper.GetString("DATABASE_DSN"),
		AccessTokenSecret:  viper.GetString("ACCESS_TOKEN_SECRET"),
		RefreshTokenSecret: viper.GetString("REFRESH_TOKEN_SECRET"),
	}
}

func GetDatabaseDSN() string {
	return AppConfig.DatabaseDSN
}

func GetAccessTokenSecret() string {
	return AppConfig.AccessTokenSecret
}

func GetRefreshTokenSecret() string {
	return AppConfig.RefreshTokenSecret
}
