package commands

import (
	"fmt"
	"log"

	c "github.com/kungfoome/pokedexcli/internal/types/config"
)

func CommandExplore(cfg *c.Config, params []string) error {
	if len(params) == 0 {
		return fmt.Errorf("location parameter missing from explore command")
	}

	fmt.Printf("Exploring %s...\n", params[0])

	areas, err := cfg.PokeapiClient.FetchLocationArea(params[0])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found Pokemon:")
	areas.PrintPokemonEncounters()
	return nil
}
