package commands

import (
	"fmt"
	"log"
	"math/rand"

	c "github.com/kungfoome/pokedexcli/internal/types/config"
)

func CommandCatch(cfg *c.Config, params []string) error {
	if len(params) == 0 {
		return fmt.Errorf("pokemon missing as a parameter")
	}

	pokemonToCatch := params[0]
	pokemonResponse, err := cfg.PokeapiClient.FetchPokemon(pokemonToCatch)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonToCatch)
	if rand.Intn(pokemonResponse.Pokemon.BaseExperience) > 40 {
		fmt.Printf("%s escaped!\n", pokemonToCatch)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonToCatch)
	cfg.Pokedex[pokemonToCatch] = pokemonResponse.Pokemon

	return nil
}
