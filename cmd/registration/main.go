package main

import (
	"log"

	"github.com/ruslanSorokin/authentication-service/cmd/registration/config"
)

const configName = "default"
const configPath = "cmd/registration/configs"

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