package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"

	"github.com/SilverLuhtoja/pokedex/internal/pokeapi"
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
		"inspect": {
			name:        "inspect <pokemon-name>",
			description: "Inspect your pokeon",
			callback:    inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows pokedex content",
			callback:    pokedex,
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
	locations, err := cfg.Client.GetLocations(cfg.Next)
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

	locations, err := cfg.Client.GetLocations(cfg.Previous)
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
	pokemons, err := cfg.Client.GetExploreLocation(areaName)
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

	pokemon, err := cfg.Client.GetPokemon(name)
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

func inspect(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("must provide a pokemon name")
	}

	name := args[0]
	pokemon, ok := cfg.Pokedex[name]
	if ok {
		logPokemonStats(pokemon)
		return nil
	}

	return errors.New("you dont have this pokemon with provided name")
}

func pokedex(cfg *Config, args ...string) error {
	if len(cfg.Pokedex) < 1 {
		fmt.Println("Pokedex is empty.")
	}
	for key := range cfg.Pokedex {
		fmt.Printf("- %s\n", key)
	}
	return nil
}

func logPokemonStats(pokemon pokeapi.Pokemon) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, types := range pokemon.Types {
		fmt.Printf("- %s\n", types.Type.Name)
	}
}
