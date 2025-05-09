package main

import "fmt"

func commandPokedex(cfg *config, params []string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("You haven't caugh any pokemon. Get to hunting!")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
