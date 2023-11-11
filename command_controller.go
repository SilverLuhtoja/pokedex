package main

import (
	"fmt"
	"os"
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
	}
}

func helpCommand(cfg *Config) error {
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

func ExitCommand(cfg *Config) error {
	os.Exit(1)
	return nil
}

func mapForward(cfg *Config) error {
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

func mapBack(cfg *Config) error {
	if cfg.Previous == nil {
		fmt.Println("On First page, cant go back.")
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
