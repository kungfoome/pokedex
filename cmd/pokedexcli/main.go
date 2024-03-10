package main

import (
	"time"

	"github.com/kungfoome/pokedexcli/api/pokeapi"
	"github.com/kungfoome/pokedexcli/internal/repl"

	p "github.com/kungfoome/pokedexcli/api/types/pokemon"
	c "github.com/kungfoome/pokedexcli/internal/types/config"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &c.Config{
		PokeapiClient: pokeClient,
		Pokedex:       make(map[string]p.Pokemon),
	}

	repl.StartRepl(cfg)
}
