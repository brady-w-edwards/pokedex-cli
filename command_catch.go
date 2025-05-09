package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, params []string) error {
	if len(params) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := params[0]
	pokemon, err := cfg.pokeapiClient.ListPokemonInfo(name)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	catchRate := rand.Intn(pokemon.BaseExperience)

	if catchRate < 40 {
		cfg.pokedex[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
		return nil
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
}
