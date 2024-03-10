package pokeapi

import (
	t "github.com/kungfoome/pokedexcli/api/types/common"
)

type LocationArea struct {
	Id                   int                     `json:"id"`
	Name                 string                  `json:"name"`
	GameIndex            int                     `json:"game_index"`
	EncounterMethodRates []t.EncounterMethodRate `json:"encounter_method_rates"`
	Location             t.NamedAPIResource      `json:"location"`
	Names                []t.Name                `json:"names"`
	PokemonEncounters    []t.PokemonEncounter    `json:"pokemon_encounters"`
}
