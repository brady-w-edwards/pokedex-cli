package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/brady-w-edwards/pokedex-cli/internal"
)

func commandMap(c *internal.Config) error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		log.Fatal("Get failed: %s", err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal("Failed to read body: %s", err)
	}

	var locations internal.LocationsArea
	err = json.Unmarshal(body, &locations)
	if err != nil {
		fmt.Println("Error unmarshalling data: %s", err)
	}

	c.Next = locations.Next
	c.Previous = locations.Previous

	fmt.Print(locations.Results)
}
