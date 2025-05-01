package main

import (
	"fmt"
	"os"
	"pokedex-cli/internal/config"
)

func commandExit(c *config.Preview) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
