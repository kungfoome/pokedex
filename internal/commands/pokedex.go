package commands

import (
	"fmt"

	c "github.com/kungfoome/pokedexcli/internal/types/config"
)

func CommandPokedex(cfg *c.Config, params []string) error {
	for _, pokemon := range cfg.Pokedex {
		fmt.Printf("- %s\n", pokemon.Name)
	}

	return nil
}

func CommandPokedexInspect(cfg *c.Config, params []string) error {
	if len(params) == 0 {
		return fmt.Errorf("pokemon missing as a parameter")
	}

	pokemonName := params[0]
	pokemon, exists := cfg.Pokedex[pokemonName]

	if !exists {
		return fmt.Errorf("pokemon %s does not exist in the pokedex and has not been caught", pokemonName)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, pokemonStat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", pokemonStat.Stat.Name, pokemonStat.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokemonType.Type.Name)
	}

	return nil
}
