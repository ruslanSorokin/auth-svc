package main

import (
	"log"

	"github.com/ruslanSorokin/auth-svc/cmd/authorizer/config"
)

const configName = "default"
const configPath = "cmd/authorizer/configs"

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	_, err := config.Load(configPath, configName)
	if err != nil {
		return err
	}

	return nil
}
