package main

import (
	"bufio"
	"fmt"
	"os"
	config "pokedex-cli/internal"
	"strings"
)

func startRepl(c *config.Preview) {
	// REPL LOOP
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]

		if exists {
			err := command.callback(c)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	cleaned := strings.Fields(lowerCase)
	return cleaned
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config.Preview) error
}

// COMMAND REGISTRY
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 map locations in the world",
			callback:    commandMap,
		},
	}
}
