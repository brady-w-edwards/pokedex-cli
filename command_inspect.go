package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, params []string) error {
	if len(params) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := params[0]
	pokemon, err := cfg.pokedex[name]
	if !err {
		return errors.New("you have not caught this pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.Effort)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
