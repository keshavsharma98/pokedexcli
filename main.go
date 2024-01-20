package main

import (
	"time"

	"github.com/keshavsharma98/pokedexcli/internal/pokeapis"
)

type config struct {
	pokeapiClient *pokeapis.Client
	nextPageURL   *string
	prevPageURL   *string
}

func main() {
	config := config{
		pokeapiClient: pokeapis.NewCLient(time.Second * 30),
	}
	StartRepl(&config)
}
