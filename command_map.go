package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	config "pokedex-cli/internal"
)

func commandMap(c *config.Preview) error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		fmt.Printf("Get failed: %s", err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		fmt.Printf("Failed to read body: %s", err)
	}

	var locations config.LocationsArea
	err = json.Unmarshal(body, &locations)
	if err != nil {
		fmt.Printf("Error unmarshalling data: %s", err)
	}

	c.Next = locations.Next
	c.Previous = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
