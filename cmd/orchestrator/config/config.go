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
			Session config.DocTable
		}
	}

	// Service config
	Service struct {
		Password           string
		AccessTokenSecrets [3]string
	}
)

// OrchestratorConfig stores general parametrs for the orchestration service
type OrchestratorConfig struct {
	Repo    Repository
	Handler config.Handler
	Service Service
}

// Load config from path
func Load(path, name string) (*OrchestratorConfig, error) {
	var cfg OrchestratorConfig
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
