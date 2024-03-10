package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	p "github.com/kungfoome/pokedexcli/api/types/pokemon"
)

type PokemonResponse struct {
	Pokemon p.Pokemon
}

func (c *Client) FetchPokemon(p string) (PokemonResponse, error) {
	var pokemon PokemonResponse

	endpoint := fmt.Sprintf("%s/pokemon/%s", endpointBase, p)

	if val, ok := c.cache.Get(endpoint); ok {
		if err := json.Unmarshal(val, &pokemon.Pokemon); err != nil {
			return pokemon, fmt.Errorf("error unmarshalling API response: %w", err)
		}

		return pokemon, nil
	}

	resp, err := http.Get(endpoint)
	if err != nil {
		return pokemon, fmt.Errorf("error making API request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 299 {
		return pokemon, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokemon, fmt.Errorf("error reading API response body: %w", err)
	}

	// fmt.Printf("%s", body)

	if err := json.Unmarshal(body, &pokemon.Pokemon); err != nil {
		return pokemon, fmt.Errorf("error unmarshalling API response: %w", err)
	}

	c.cache.Add(endpoint, body)

	return pokemon, nil
}
