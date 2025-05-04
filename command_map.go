package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	res, err := cfg.pokeapiClient.ListLocations(cfg.Next)

	if err != nil {
		return err
	}

	cfg.Next = res.Next
	cfg.Previous = res.Previous

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous == nil {
		return errors.New("you're on the first page")
	}
	res, err := cfg.pokeapiClient.ListLocations(cfg.Previous)

	if err != nil {
		return err
	}

	cfg.Next = res.Next
	cfg.Previous = res.Previous

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
