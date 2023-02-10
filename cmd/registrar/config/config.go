package config

import (
	"github.com/ruslanSorokin/authentication-service/internal/config"
	"github.com/spf13/viper"
)

type (
	// Repository config
	Repository struct {

		// Mongo config
		Mongo struct {
			Account config.DocTable
		}
	}
)

// RegistrarConfig stores general parametrs for the registration service
type RegistrarConfig struct {
	Repo    Repository
	Handler config.Handler
}

// Load config from path/name
func Load(p, n string) (*RegistrarConfig, error) {
	var cfg RegistrarConfig
	viper := viper.New()

	viper.SetConfigType("yaml")
	viper.SetConfigName(n)
	viper.AddConfigPath(p)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
