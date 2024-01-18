package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/keshavsharma98/pokedexcli/poke_apis"
)

func callbackHelp(config *config) error {
	fmt.Println("Welcome, below are the available commands:")
	commands := getCommands()

	for _, value := range commands {
		fmt.Println(value.name+": ", value.description)
	}
	return nil
}

func callbackExit(config *config) error {
	fmt.Println("Exiting Pokedexcli")
	os.Exit(0)
	return nil
}

func callbackMap(config *config) error {
	client := poke_apis.NewClient()
	result, err := client.GetLocations(config.nextPageURL)
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

func callbackMapb(config *config) error {
	if config.prevPageURL == nil {
		return errors.New("you are on the first page")
	}

	client := poke_apis.NewClient()
	result, err := client.GetLocations(config.prevPageURL)
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
