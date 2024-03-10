package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kungfoome/pokedexcli/internal/commands"
	c "github.com/kungfoome/pokedexcli/internal/types/config"
)

func StartRepl(cfg *c.Config) {
	const prefix = "Pokedex > "

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prefix)
		scanner.Scan()

		cmdTokens := strings.Fields(strings.ToLower(scanner.Text()))
		cmdName := cmdTokens[0]
		cmdParameters := cmdTokens[1:]
		if cmdName == "" {
			commands.CommandHelp(cfg, cmdParameters)
			continue
		}

		cmd, ok := commands.GetCommands()[cmdName]
		if ok {
			err := cmd.Callback(cfg, cmdParameters)
			if err != nil {
				fmt.Println(err)
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
	}
}
