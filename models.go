package main

import (
	"github.com/SilverLuhtoja/pokedex/internal/pokeapi"
	"net/http"
)

type CliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type PokeApiClient struct {
	Client http.Client
}

type Config struct {
	Client   Client
	Next     *string
	Previous *string
}

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}
