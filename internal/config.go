package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Username string `mapstructure:"WS_USERNAME"`
	Password string `mapstructure:"WS_PASSWORD"`
}

var AppConfig *Config

func init() {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")
	err := v.ReadInConfig()
	if err != nil {
		failOnError(err, "Failed to read config")
	}
	v.AutomaticEnv()
	err = v.Unmarshal(&AppConfig)
	if err != nil {
		failOnError(err, "Failed to read environment")
	}
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
