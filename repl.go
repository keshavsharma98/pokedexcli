package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	nextPageURL *string
	prevPageURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Below are the next 20 locations:",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Below are the prev 20 locations",
			callback:    callbackMapb,
		},
	}
}

func cleanCommand(command string) string {
	command = strings.TrimSpace(command)
	command = strings.ToLower(command)
	return command
}

func StartRepl() {
	config := config{}
	input := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nPodexcli >")
		command, err := input.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input: ", err)
			return
		}

		command = cleanCommand(command)

		commands := getCommands()

		if command == "" {
			continue
		}

		c, ok := commands[command]
		if !ok {
			fmt.Println("Command not found. Please ender help to see available commands")
			continue
		}

		err = c.callback(&config)
		if err != nil {
			fmt.Println(fmt.Errorf("error executing command map: %s", err))
		}
	}
}
