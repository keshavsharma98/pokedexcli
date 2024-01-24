package main

import "strings"

type CliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
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
		"save": {
			name:        "save",
			description: "Saves the progress",
			callback:    callbackSave,
		},
	}
}

func CleanCommand(cmd string) []string {
	cmd = strings.TrimSpace(cmd)
	cmd = strings.ToLower(cmd)
	return strings.Fields(cmd)
}
