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

type Resource interface {
	Resource()
}

func (loc LocationResponse) Resource() {}
