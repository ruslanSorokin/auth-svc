package config

import (
	"time"

	"github.com/spf13/viper"
)

const config_path = "config/configs"

type Config struct {
	MongoDBURI string `mapstructure:"MDB_URI"`

	InternalServicePort int `mapstructure:"INTERNAL_SERVICE_PORT"`
	ExternalServicePort int `mapstructure:"EXTERNAL_SERVICE_PORT"`

	AccessTokenSecretKey  string `mapstructure:"ACCESS_TOKEN_SECRET_KEY"`
	RefreshTokenSecretKey string `mapstructure:"REFRESH_TOKEN_SECRET_KEY"`

	AccessTokenExpiry  time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRY"`
	RefreshTokenExpiry time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRY"`

	InternalPassword string `mapstructure:"INTERNAL_PASSWORD"`
}

func Load(name string) (cfg Config, err error) {
	viper.AddConfigPath(config_path)
	viper.SetConfigName(name)

	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)
	return
}
