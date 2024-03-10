package commands

import (
	"log"

	c "github.com/kungfoome/pokedexcli/internal/types/config"
)

func CommandMapForward(cfg *c.Config, params []string) error {
	areas, err := cfg.PokeapiClient.FetchLocationAreas(cfg.NextLocationsURL)

	if err != nil {
		log.Fatal(err)
	}

	cfg.NextLocationsURL = areas.Next
	cfg.PrevLocationsURL = areas.Previous

	areas.PrintAreas()
	return nil
}

func CommandMapBack(cfg *c.Config, params []string) error {
	areas, err := cfg.PokeapiClient.FetchLocationAreas(cfg.PrevLocationsURL)

	if err != nil {
		log.Fatal(err)
	}

	cfg.NextLocationsURL = areas.Next
	cfg.PrevLocationsURL = areas.Previous

	areas.PrintAreas()
	return nil
}
