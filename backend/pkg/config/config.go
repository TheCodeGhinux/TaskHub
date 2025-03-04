package config

import (
	"log"

	"github.com/TheCodeGhinux/TaskHub/taskhub/config"
	"github.com/spf13/viper"
)

var configurations config.AppConfig

func LoadConfig() config.AppConfig {

	viper.SetConfigName("appConfig")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config")
	}

	err := viper.Unmarshal(&configurations)
	if err != nil {
		log.Fatalf("unable to decode into struct: %v", err)
	}

	return configurations
}
