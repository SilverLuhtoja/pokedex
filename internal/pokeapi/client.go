package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	cache "github.com/SilverLuhtoja/pokedex/internal/pokecache"
)

func NewClient() Client {
	return Client{
		httpClient: http.Client{},
		BaseURL:    "https://pokeapi.co/api/v2/",
	}
}

func (c *Client) GetPokemon(cache *cache.Cache, pokemon_name string) (Pokemon, error) {
	url := c.BaseURL + "/pokemon/" + pokemon_name
	responseBody, ok := cache.Get(url)
	if !ok {
		res, err := c.GetRawResponse(url)
		if err != nil {
			return Pokemon{}, err
		}

		cache.Add(url, res)
		responseBody = res
	}
	pokemon, err := ConvertToDomain[Pokemon](responseBody)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}

func (c *Client) GetExploreLocation(cache *cache.Cache, areaName string) (ExploreResponse, error) {
	url := c.BaseURL + "/location-area/" + areaName

	responseBody, ok := cache.Get(url)

	if !ok {
		res, err := c.GetRawResponse(url)
		if err != nil {
			return ExploreResponse{}, err
		}

		cache.Add(url, res)
		responseBody = res
	}
	pokemons, err := ConvertToDomain[ExploreResponse](responseBody)
	if err != nil {
		return ExploreResponse{}, err
	}
	return pokemons, nil
}

func (c *Client) GetLocations(pageURL *string, cache *cache.Cache) (LocationResponse, error) {
	url := c.BaseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	responseBody, ok := cache.Get(url)
	if !ok {
		res, err := c.GetRawResponse(url)
		if err != nil {
			fmt.Printf("Error getting data: %s", err)
			return LocationResponse{}, err
		}

		cache.Add(url, res)
		responseBody = res
	}

	locations, err := ConvertToDomain[LocationResponse](responseBody)
	if err != nil {
		fmt.Printf("Error with unmarshaling: %s", err)
		return LocationResponse{}, err
	}
	return locations, nil
}

func (c *Client) GetRawResponse(pageURL string) (responseBody []byte, err error) {
	response, err := http.Get(pageURL)
	if err != nil {
		return responseBody, err
	}

	defer response.Body.Close()
	responseBody, err = io.ReadAll(response.Body)

	if response.StatusCode > 400 {
		errMessage := fmt.Sprintf("Error: couldn't get data from server. Code: %v - %s", response.StatusCode, responseBody)
		return responseBody, errors.New(errMessage)
	}

	if err != nil {
		return responseBody, err
	}

	return responseBody, nil
}

func ConvertToDomain[T Resource](source []byte) (T, error) {
	var target T

	if err := json.Unmarshal(source, &target); err != nil {
		return target, err
	}

	return target, nil
}
