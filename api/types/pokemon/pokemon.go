package pokeapi

import (
	t "github.com/kungfoome/pokedexcli/api/types/common"
)

type Pokemon struct {
	Id                     int           `json:"id"`
	Name                   string        `json:"name"`
	BaseExperience         int           `json:"base_experience"`
	Height                 int           `json:"height"`
	Weight                 int           `json:"weight"`
	IsDefault              bool          `json:"is_default"`
	Order                  int           `json:"order"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Stats                  []PokemonStat `json:"stats"`
	Types                  []PokemonType `json:"types"`
}

type PokemonStat struct {
	Stat     t.NamedAPIResource `json:"stat"`
	Effort   int                `json:"effort"`
	BaseStat int                `json:"base_stat"`
}

type PokemonType struct {
	Type t.NamedAPIResource `json:"type"`
	Slot int                `json:"slot"`
}
