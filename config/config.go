package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DbConn string `mapstructure:"DB_URL"`
}

func LoadConfig() (c Config, err error) {
	// Try multiple paths to find .env file
	viper.AddConfigPath(".") // current directory
	viper.SetConfigName(".env")
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

	return c, nil
}
