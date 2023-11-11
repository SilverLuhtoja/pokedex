package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	pokecache "github.com/SilverLuhtoja/pokedex/internal/pokecache"
)

func NewClient() Client {
	return Client{
		httpClient: http.Client{},
		BaseURL:    "https://pokeapi.co/api/v2/",
	}
}

func (c *Client) GetLocations(pageURL *string, cache *pokecache.Cache) (LocationResponse, error) {
	url := c.BaseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// check if in cache
	responseBody, ok := cache.Get(pageURL)
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
