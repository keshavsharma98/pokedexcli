package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func StartRepl(config *config) {
	input := bufio.NewReader(os.Stdin)
	fmt.Print("                            ****** WELCOME TO POKEDEX ******\n\n\n")
	err := config.pokeapiClient.LoadGame()
	if err != nil {
		log.Fatalln("Error loading game: ", err)
	}

	for {
		fmt.Print("\nPodexcli >")
		command, err := input.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input: ", err)
			return
		}

		cleancmd := CleanCommand(command)
		command = cleancmd[0]
		args := cleancmd[1:]

		commands := GetCommands()

		if command == "" {
			continue
		}

		c, ok := commands[command]
		if !ok {
			fmt.Println("Command not found. Please ender help to see available commands")
			continue
		}

		err = c.callback(config, args...)
		if err != nil {
			fmt.Println(fmt.Errorf("error executing command %s: %s", command, err))
		}
	}
}
