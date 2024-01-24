package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/keshavsharma98/pokedexcli/internal/pokeapis"
)

type config struct {
	pokeapiClient *pokeapis.Client
	nextPageURL   *string
	prevPageURL   *string
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading.env file: %v", err)
	}
	config := config{
		pokeapiClient: pokeapis.NewCLient(time.Second * 30),
	}
	StartRepl(&config)
}
