package config

import (
	"time"

	"github.com/spf13/viper"
)

const (
	// MaxLoginLength used for validations on service-layer
	MaxLoginLength = 16
	// MinLoginLength used for validations on service-layer
	MinLoginLength = 8

	// MaxPasswordLength used for validations on service-layer
	MaxPasswordLength = 16
	// MinPasswordLength used for validations on service-layer
	MinPasswordLength = 8
)

type (
	// DB config
	DB struct {

		// Mongo config
		Mongo struct {
			URI       string
			DBName    string
			TableName struct {
				User    string
				Session string
			}
		}
	}

	// Server config
	Server struct {

		// Internal server config
		Internal struct {

			// REST server config
			REST struct {
				Port int
			}

			// RPC server config
			RPC struct {
				Port int
			}
		}

		// External server config
		External struct {

			// REST server config
			REST struct {
				Port int
			}

			// RPC server config
			RPC struct {
				Port int
			}
		}
	}

	// Service config
	Service struct {

		// Secret config
		Secret struct {
			Password string
		}
	}

	// JWT config
	JWT struct {
		AccessTokenSecretKey string

		AccessTokenExpiry  time.Duration
		RefreshTokenExpiry time.Duration
	}
)

// Config stores general parametrs for the authentication service
type Config struct {
	DB      DB
	Server  Server
	Service Service
	JWT     JWT
}

// Load config from path/name
func Load(p, n string) (*Config, error) {
	var cfg Config
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
