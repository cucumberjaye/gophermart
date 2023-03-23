package main

import (
	"github.com/cucumberjaye/gophermart/internal/pkg/app"
	"github.com/rs/zerolog/log"
)

func main() {
	martApp, err := app.New()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	err = martApp.Run()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}
