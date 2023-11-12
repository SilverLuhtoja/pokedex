package main

import (
	pokeapi "github.com/SilverLuhtoja/pokedex/internal/pokeapi"
	pokecache "github.com/SilverLuhtoja/pokedex/internal/pokecache"
)

type CliCommand struct {
	name        string
	description string
	callback    func(*Config, string) error
}

type Config struct {
	Cache    *pokecache.Cache
	Client   pokeapi.Client
	Next     *string
	Previous *string
}
