package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func()
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
	}
}

func cleanCommand(command string) string {
	command = strings.TrimSpace(command)
	command = strings.ToLower(command)
	return command
}

func StartRepl() {

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

		c.callback()
	}
}
