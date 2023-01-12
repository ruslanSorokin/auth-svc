package main

import (
	"log"

	config "github.com/ruslanSorokin/auth-service/config"
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
