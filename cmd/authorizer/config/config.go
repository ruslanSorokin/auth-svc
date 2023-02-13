package config

import (
	"time"

	"github.com/ruslanSorokin/auth-svc/internal/config"
	"github.com/spf13/viper"
)

type (
	// Repository config
	Repository struct {

		// Mongo config
		Mongo struct {
			Account config.DocTable
			Session config.DocTable
		}
	}

	// Service config
	Service struct {
		AccessTokenSecrets [3]string

		TTL struct {
			AccessToken  time.Duration
			RefreshToken time.Duration
		}
	}
)

// AuthorizorConfig stores general parametrs for the authentication service
type AuthorizorConfig struct {
	Repo    Repository
	Handler config.Handler
	Service Service
}

// Load config from path
func Load(path, name string) (*AuthorizorConfig, error) {
	var cfg AuthorizorConfig
	viper := viper.New()

	viper.SetConfigType("yaml")
	viper.SetConfigName(name)
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
