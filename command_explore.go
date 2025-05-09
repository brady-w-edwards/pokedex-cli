package main

import "fmt"

func commandExplore(cfg *config, params []string) error {
	res, err := cfg.pokeapiClient.ListAreaInfo(params[0])

	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\nFound pokemon:\n", params[0])
	for _, loc := range res.PokemonEncounters {
		fmt.Println("- " + loc.Pokemon.Name)
	}
	return nil
}
