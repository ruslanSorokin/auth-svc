package main

import (
	"log"

	"github.com/ruslanSorokin/auth-service/internal/pkg/config"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	_, err := config.Load("default.env")
	if err != nil {
		return err
	}

	return nil
}
