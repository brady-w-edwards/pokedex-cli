package main

import "pokedex-cli/internal/config"

func main() {
	c := new(config.Preview)

	startRepl(c)
}
