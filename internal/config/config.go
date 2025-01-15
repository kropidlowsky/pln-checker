package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	LogFile string `mapstructure:"LOG_FILE"`
}

func LoadConfig() Config {
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("error wile reading the config env variables: %s", err.Error())
	}

	return config
}

func init() {
	viper.AutomaticEnv()

	viper.SetDefault("LOG_FILE", "logs/log.txt")
}
