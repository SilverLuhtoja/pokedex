package main

import (
	"errors"
	"fmt"
	"math/rand"
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
			name:        "explore <location-name>",
			description: "Explore areas for pokemons",
			callback:    explore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Try to catch a pokemon",
			callback:    catch,
		},
	}
}

func helpCommand(cfg *Config, args ...string) error {
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

func ExitCommand(cfg *Config, args ...string) error {
	os.Exit(1)
	return nil
}

func mapForward(cfg *Config, args ...string) error {
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

func mapBack(cfg *Config, args ...string) error {
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

func explore(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("must provide a location name")
	}
	areaName := args[0]
	pokemons, err := cfg.Client.GetExploreLocation(cfg.Cache, areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", areaName)
	if len(pokemons.PokemonEncounters) > 0 {
		fmt.Println("Found Pokemon:")
	}
	for _, encounter := range pokemons.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func catch(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("must provide a pokemon name")
	}

	name := args[0]
	if _, ok := cfg.Pokedex[name]; ok {
		fmt.Println("Pokemon is already Caught")
		return nil
	}
	pokemon, err := cfg.Client.GetPokemon(cfg.Cache, name)
	if err != nil {
		return err
	}

	randomInt := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if randomInt < pokemon.BaseExperience/2 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	cfg.Pokedex[pokemon.Name] = pokemon

	return nil
}
