package config

import (
	"time"

	"github.com/spf13/viper"
)

type (
	// DB config
	DB struct {
		Mongo struct {
			Account struct {
				URI   string
				DB    string
				Table string
			}
			Session struct {
				URI   string
				DB    string
				Table string
			}
		}
	}

	// Server config
	Server struct {
		Internal struct {
			REST struct {
				Port int
			}

			RPC struct {
				Port int
			}
		}

		External struct {
			REST struct {
				Port int
			}

			RPC struct {
				Port int
			}
		}
	}

	// Service config
	Service struct {
		Secret struct {
			Password string
		}
	}

	// JWT config
	JWT struct {
		AccessTokenSecrets [3]string

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

// Load config from path
func Load(path, name string) (*Config, error) {
	var cfg Config
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
