package config

import (
	"github.com/kungfoome/pokedexcli/api/pokeapi"
	p "github.com/kungfoome/pokedexcli/api/types/pokemon"
)

type Config struct {
	PokeapiClient    pokeapi.Client
	NextLocationsURL *string
	PrevLocationsURL *string
	Pokedex          map[string]p.Pokemon
}
