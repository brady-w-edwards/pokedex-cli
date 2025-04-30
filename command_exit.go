package main

import (
	"fmt"
	"os"
	config "pokedex-cli/internal"
)

func commandExit(c *config.Preview) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
