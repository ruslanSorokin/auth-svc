package config

import (
	"time"

	"github.com/spf13/viper"
)

type (
	app struct {
		MongoDBURI string `mapstructure:"MONGO_DB_URI"`

		InternalServicePort int `mapstructure:"INTERNAL_SERVICE_PORT"`
		ExternalServicePort int `mapstructure:"EXTERNAL_SERVICE_PORT"`

		InternalPassword string `mapstructure:"INTERNAL_PASSWORD"`
	}

	jwt struct {
		AccessTokenSecretKey string `mapstructure:"ACCESS_TOKEN_SECRET_KEY"`

		AccessTokenExpiry  time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRY"`
		RefreshTokenExpiry time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRY"`
	}
)

// Config stores general parametrs for the entire authentication service
type Config struct {
	App app
	JWT jwt
}

// Load config from path/name
func Load(path, name string) (*Config, error) {
	var cfg Config
	viper.AddConfigPath(path)
	viper.SetConfigName(name)

	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg.App)
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg.JWT)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
