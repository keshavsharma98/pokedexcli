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
	callback    func(*config, ...string) error
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
			description: "Exit the PokedexCli",
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
		"explore": {
			name:        "explore {location_name}",
			description: "Shows all Pokemon in the given location",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Catches the given pokemon and adds them to users pokedex.",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "Shows the details of the given pokemon if caught",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows all pokemon caught by the user",
			callback:    callbackPokedex,
		},
	}
}

func cleanCommand(cmd string) []string {
	cmd = strings.TrimSpace(cmd)
	cmd = strings.ToLower(cmd)
	return strings.Fields(cmd)
}

func StartRepl(config *config) {
	input := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nPodexcli >")
		command, err := input.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input: ", err)
			return
		}

		cleancmd := cleanCommand(command)
		command = cleancmd[0]
		args := cleancmd[1:]

		commands := getCommands()

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
