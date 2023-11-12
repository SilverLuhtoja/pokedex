package pokeapi

import "net/http"

type Client struct {
	httpClient http.Client
	BaseURL    string
}

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationResponse struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

type Pokemon struct {
	Pokemon struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon"`
}

type ExploreResponse struct {
	PokemonEncounters []Pokemon `json:"pokemon_encounters"`
}

type Resource interface {
	Resource()
}

func (loc LocationResponse) Resource() {}
func (loc ExploreResponse) Resource()  {}
