package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
			callback:    mapCommand,
		},
		// "mapb": {
		// 	name:        "mapb",
		// 	description: "Displays previous 20 location areas",
		// 	callback:    commands.ExitCommand,
		// },
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

// func mapCommand(conf *config) error {

func mapCommand(cfg *Config) error {
	var location_endpoint string = "https://pokeapi.co/api/v2/location/"
	res, err := http.Get(location_endpoint)

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	response := LocationResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}

	// LocationResponse
	fmt.Println(response)

	return nil

}

// if on first page, give error
// func mapBackCommand(conf *config) error {
// return nil
// }
