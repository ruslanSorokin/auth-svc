package config

import (
	"time"

	"github.com/spf13/viper"
)

type (
	// Network part of the config contains all network stuff
	Network struct {
		MongoDBURI string `mapstructure:"MONGO_DB_URI"`

		InternalServicePort int `mapstructure:"INTERNAL_SERVICE_PORT"`
		ExternalServicePort int `mapstructure:"EXTERNAL_SERVICE_PORT"`
	}

	// Service part of the config contains everything used in usecases
	Service struct {
		MinUsernameLength int `mapstructure:"MIN_USERNAME_LENGTH"`
		MinPasswordLength int `mapstructure:"MIN_PASSWORD_LENGTH"`

		MaxUsernameLength int `mapstructure:"MAX_USERNAME_LENGTH"`
		MaxPasswordLength int `mapstructure:"MAX_PASSWORD_LENGTH"`
	}

	// JWT contains everething for JWT configuration
	JWT struct {
		AccessTokenSecretKey string `mapstructure:"ACCESS_TOKEN_SECRET_KEY"`

		AccessTokenExpiry  time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRY"`
		RefreshTokenExpiry time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRY"`
	}

	// Secret part of the config contains the service's secrets
	Secret struct {
		InternalPassword string `mapstructure:"PASSWORD"`
	}
)

// Config stores general parametrs for the authentication service
type Config struct {
	Network Network `mapstructure:",squash"`
	Service Service `mapstructure:",squash"`
	JWT     JWT     `mapstructure:",squash"`
	Secret  Secret  `mapstructure:",squash"`
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

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
