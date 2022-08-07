package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"restapi-with-opentelemetry/config"
	"restapi-with-opentelemetry/internal/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msg("Error loading .env file")
	}
	config.Load()

	s := server.NewServer()
	s.Run()
}
