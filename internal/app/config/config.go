package config

import (
	"time"

	"github.com/spf13/viper"
)

type (
	// DB config
	DB struct {

		// Mongo config
		Mongo struct {
			URI      string
			Username string
			Password string
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

		// Validation config
		Validation struct {

			// UsernameLength config
			UsernameLength struct {
				Min int
				Max int
			}

			// PasswordLength config
			PasswordLength struct {
				Min int
				Max int
			}
		}

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
