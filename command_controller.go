package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func getCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    ExitCommand,
		},
		"map": {
			name:        "map",
			description: "Displays names of 20 next location areas",
			callback:    mapForward,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 location areas",
			callback:    mapBack,
		},
		"explore": {
			name:        "explore",
			description: "Explore areas for pokemons",
			callback:    explore,
		},
	}
}

func helpCommand(cfg *Config, _ string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func ExitCommand(cfg *Config, _ string) error {
	os.Exit(1)
	return nil
}

func mapForward(cfg *Config, _ string) error {
	locations, err := cfg.Client.GetLocations(cfg.Next, cfg.Cache)
	if err != nil {
		return err
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func mapBack(cfg *Config, _ string) error {
	if cfg.Previous == nil {
		color.Red("On First page, cant go back.")
		return nil
	}

	locations, err := cfg.Client.GetLocations(cfg.Previous, cfg.Cache)
	if err != nil {
		return err
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func explore(cfg *Config, areaName string) error {
	pokemons, err := cfg.Client.GetExploreLocation(areaName, cfg.Cache)
	if err != nil {
		return err
	}

	for _, encounter := range pokemons.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
