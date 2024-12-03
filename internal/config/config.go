package config

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Env struct {
		DatabaseURL string `mapstructure:"DB_URL"`
		Port        int64  `mapsctructure:"PORT"`
	}
	DB struct {
		DSM string `mapstructure:"DSN"`
	}
}

func EnvLoad() {
	viper.AutomaticEnv()
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}

	return &config, nil
}

func SetupDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return db, nil
}
