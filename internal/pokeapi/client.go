package pokeapi

import "net/http"

type Client struct {
	httpClient http.Client
}

func newClient() Client {
	return Client{httpClient: http.Client{}}
}
