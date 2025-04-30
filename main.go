package main

import config "pokedex-cli/internal"

func main() {
	c := new(config.Preview)

	startRepl(c)
}
