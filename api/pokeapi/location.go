package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	l "github.com/kungfoome/pokedexcli/api/types/locationarea"
)

type LocationAreaResponse struct {
	Count         int              `json:"count"`
	Next          *string          `json:"next"`
	Previous      *string          `json:"previous"`
	LocationAreas []l.LocationArea `json:"results"`
}

type LocationAreaPokemonResponse struct {
	LocationArea l.LocationArea
}

const endpointBase = "https://pokeapi.co/api/v2"

func (c *Client) FetchLocationAreas(url *string) (LocationAreaResponse, error) {
	var locationAreas LocationAreaResponse

	// If we don't have a prev or next url specified, we just use the default
	// first endpoint
	endpoint := fmt.Sprintf("%s/location-area", endpointBase)
	if url != nil {
		endpoint = *url
	}

	if val, ok := c.cache.Get(endpoint); ok {
		if err := json.Unmarshal(val, &locationAreas); err != nil {
			return locationAreas, fmt.Errorf("error unmarshalling API response: %w", err)
		}

		return locationAreas, nil
	}

	resp, err := http.Get(endpoint)
	if err != nil {
		return locationAreas, fmt.Errorf("error making API request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 299 {
		return locationAreas, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAreas, fmt.Errorf("error reading API response body: %w", err)
	}

	// fmt.Printf("%s", body)

	if err := json.Unmarshal(body, &locationAreas); err != nil {
		return locationAreas, fmt.Errorf("error unmarshalling API response: %w", err)
	}

	c.cache.Add(endpoint, body)

	return locationAreas, nil
}

func (c *Client) FetchLocationArea(location string) (LocationAreaPokemonResponse, error) {
	var locationArea LocationAreaPokemonResponse

	endpoint := fmt.Sprintf("%s/location-area/%s", endpointBase, location)

	if val, ok := c.cache.Get(endpoint); ok {
		if err := json.Unmarshal(val, &locationArea.LocationArea); err != nil {
			return locationArea, fmt.Errorf("error unmarshalling API response: %w", err)
		}

		return locationArea, nil
	}

	resp, err := http.Get(endpoint)
	if err != nil {
		return locationArea, fmt.Errorf("error making API request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 299 {
		return locationArea, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationArea, fmt.Errorf("error reading API response body: %w", err)
	}

	// fmt.Printf("%s", body)

	if err := json.Unmarshal(body, &locationArea.LocationArea); err != nil {
		return locationArea, fmt.Errorf("error unmarshalling API response: %w", err)
	}

	c.cache.Add(endpoint, body)

	return locationArea, nil
}

func (a LocationAreaResponse) PrintAreas() {
	for _, area := range a.LocationAreas {
		fmt.Println(area.Name)
	}
}

func (a LocationAreaPokemonResponse) PrintPokemonEncounters() {
	for _, pokemon := range a.LocationArea.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
}
