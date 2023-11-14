package main

import (
	pokeapi "github.com/SilverLuhtoja/pokedex/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient()
	cfg := &Config{
		Client:  client,
		Pokedex: map[string]pokeapi.Pokemon{},
	}

	initApp(cfg)
}
