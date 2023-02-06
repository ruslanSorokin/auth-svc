package main

import (
	"log"

	"github.com/ruslanSorokin/auth-service/internal/app/config"
)

const defaultConfigPath = "configs"
const defaultConfigName = "default"

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	_, err := config.Load(defaultConfigPath, defaultConfigName)
	if err != nil {
		return err
	}

	return nil
}
