package config

import (
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
		}
	}

	// Server config
	Server struct {
		REST struct {
			Port int
		}

		RPC struct {
			Port int
		}
	}
)

// Config stores general parametrs for the authentication service
type Config struct {
	DB     DB
	Server Server
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
