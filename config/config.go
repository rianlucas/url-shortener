package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string `json:"SERVER_PORT"`
	DbConn     string `json:"DB_URL"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("../")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	return
}
