package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	MongoURI string `mapstructure:"MONGO_URI"`
	DBName   string `mapstructure:"DB_NAME"`
	Port     string `mapstructure:"PORT"`
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}