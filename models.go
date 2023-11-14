package main

import (
	pokeapi "github.com/SilverLuhtoja/pokedex/internal/pokeapi"
)

type CliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

type Config struct {
	Client   pokeapi.Client
	Next     *string
	Previous *string
	Pokedex  map[string]pokeapi.Pokemon
}
