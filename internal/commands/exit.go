package commands

import (
	"os"

	c "github.com/kungfoome/pokedexcli/internal/types/config"
)

func CommandExit(cfg *c.Config, params []string) error {
	os.Exit(0)
	return nil
}
