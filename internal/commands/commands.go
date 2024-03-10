package commands

import (
	c "github.com/kungfoome/pokedexcli/internal/types/config"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*c.Config, []string) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"map": {
			Name:        "map",
			Description: "Returns a list of the next 20 locations",
			Callback:    CommandMapForward,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Returns a list of the previous 20 locations",
			Callback:    CommandMapBack,
		},
		"explore": {
			Name:        "explore",
			Description: "Returns a list of Pokemon based on the location provided",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Tries to catch the pokemon passed in as a parameter",
			Callback:    CommandCatch,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "List all the pokemon that were captured and added to the pokedex",
			Callback:    CommandPokedex,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect a pokemon that has been caught inside the pokedex",
			Callback:    CommandPokedexInspect,
		},
	}
}
