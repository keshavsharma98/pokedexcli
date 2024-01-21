package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func callbackHelp(config *config, args ...string) error {
	fmt.Println("Welcome, below are the available commands:")
	commands := getCommands()

	for _, value := range commands {
		fmt.Println(value.name+": ", value.description)
	}
	return nil
}

func callbackExit(config *config, args ...string) error {
	fmt.Println("Exiting Pokedexcli")
	os.Exit(0)
	return nil
}

func callbackMap(config *config, args ...string) error {
	result, err := config.pokeapiClient.GetLocations(config.nextPageURL)
	if err != nil {
		log.Fatalln(err)
	}

	for _, obj := range result.Results {
		fmt.Println(obj.Name)
	}
	config.nextPageURL = result.Next
	config.prevPageURL = result.Previous
	return nil
}

func callbackMapb(config *config, args ...string) error {
	if config.prevPageURL == nil {
		return errors.New("you are on the first page")
	}

	result, err := config.pokeapiClient.GetLocations(config.prevPageURL)
	if err != nil {
		log.Fatalln(err)
	}

	for _, obj := range result.Results {
		fmt.Println(obj.Name)
	}
	config.nextPageURL = result.Next
	config.prevPageURL = result.Previous
	return nil
}

func callbackExplore(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New(" Invalid location")
	}

	result, err := config.pokeapiClient.GetPokemonsInArea(args[0])
	if err != nil {
		if err.Error() == "invalid location" {
			return err
		}
		log.Fatalln(err)
	}

	fmt.Println("Found Pokemon:")
	for _, obj := range result.PokemonEncounters {
		fmt.Println("- ", obj.Pokemon.Name)
	}

	return nil
}

func callbackCatch(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New(" Invalid pokemon name")
	}

	fmt.Printf("Throwing a Pokeball at %s\n", args[0])
	isCatched, err := config.pokeapiClient.CatchPokemon(args[0])
	if err != nil {
		if err.Error() == "invalid pokemon" {
			return err
		}
		log.Fatalln(err)
	}

	if isCatched {
		fmt.Printf("%s was caught!\n", args[0])
		return nil
	}

	fmt.Printf("%s escaped!\n", args[0])

	return nil
}

func callbackInspect(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New(" Invalid pokemon name")
	}

	isCaught, pokemon, err := config.pokeapiClient.InspectPokemon(args[0])
	if err != nil {
		log.Fatalln(err)
		return err
	}

	if !isCaught {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats: ")
	for _, value := range pokemon.Stats {
		fmt.Println("	-", value.Stat.Name, ": ", value.BaseStat)
	}
	fmt.Println("Types: ")
	for _, value := range pokemon.Types {
		fmt.Println("	-", value.Type.Name)
	}

	return nil
}

func callbackPokedex(config *config, args ...string) error {
	pokemons, err := config.pokeapiClient.GetAllPokedexPokemons()
	if err != nil {
		log.Fatalln(err)
		return err
	}

	fmt.Println("Your Pokedex:")
	if len(pokemons) == 0 {
		fmt.Println("	- (empty)")
		return nil
	}
	for _, pokemon := range pokemons {
		fmt.Println("	-", pokemon)
	}

	return nil
}
