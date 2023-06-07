package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseDSN string
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
		DatabaseDSN: viper.GetString("DATABASE_DSN"),
		// Initialize other configuration properties here
	}
}

func GetDatabaseDSN() string {
	return AppConfig.DatabaseDSN
}
