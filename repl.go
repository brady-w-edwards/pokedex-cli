package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex-cli/internal/pokeapi"
	"strings"
)

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	pokedex       map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
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
			err := command.callback(cfg, words[1:])
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
	callback    func(*config, []string) error
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
		"mapf": {
			name:        "mapf",
			description: "Displays the next 20 map locations in the world",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 map locations in the world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explores the area typed after the command",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Throws a pokeball and makes an attemp to catch a pokemon",
			callback:    commandCatch,
		},
	}
}
