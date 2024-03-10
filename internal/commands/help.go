package commands

import (
	"fmt"
	"strings"

	c "github.com/kungfoome/pokedexcli/internal/types/config"
)

func CommandHelp(cfg *c.Config, params []string) error {
	var helpMsg strings.Builder
	helpMsg.WriteString("\nWelcome to the Pokedex!\n")
	helpMsg.WriteString("Usage:\n\n")

	for _, cmd := range GetCommands() {
		fmt.Fprintf(&helpMsg, "%s: %s\n", cmd.Name, cmd.Description)
	}

	helpMsg.WriteString("\n")

	fmt.Print(helpMsg.String())
	return nil
}
