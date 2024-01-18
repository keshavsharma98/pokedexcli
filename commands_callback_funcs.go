package main

import (
	"fmt"
	"os"
)

func callbackHelp() {
	fmt.Println("Welcome, below are the available commands:")
	commands := getCommands()

	for _, value := range commands {
		fmt.Println(value.name+": ", value.description)
	}
}

func callbackExit() {
	fmt.Println("Exiting Pokedexcli")
	os.Exit(0)
}
