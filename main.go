package main

import (
	"time"

	pokeapi "github.com/SilverLuhtoja/pokedex/internal/pokeapi"
	pokecache "github.com/SilverLuhtoja/pokedex/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(10 * time.Second)
	client := pokeapi.NewClient()
	cfg := &Config{
		Client: client,
		Cache:  cache,
	}

	initApp(cfg)
}
