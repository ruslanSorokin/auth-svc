package main

import (
	"log"

	"github.com/ruslanSorokin/auth-service/internal/app/config"
)

const DefaultConfigPath = "configs"

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	_, err := config.Load(DefaultConfigPath, "default.env")
	if err != nil {
		return err
	}

	return nil
}
