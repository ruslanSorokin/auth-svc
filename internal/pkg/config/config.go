package config

import (
	"time"

	"github.com/spf13/viper"
)

const config_path = "configs/"

type Config struct {
	PrivateAuthServerHostname string `mapstructure:"PRIVATE_SERVER_HOST"`
	PrivateAuthServerPort     string `mapstructure:"PRIVATE_SERVER_PORT"`

	PublicAuthServerHostname string `mapstructure:"PUBLIC_SERVER_HOST"`
	PublicAuthServerPort     string `mapstructure:"PUBLIC_SERVER_PORT"`

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
