package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedex-cli/internal/config"
)

func commandMapb(c *config.Preview) error {
	url := "https://pokeapi.co/api/v2/location-area"

	if c.Previous != nil {
		url = *c.Previous
	} else {
		fmt.Printf("You're on the first page\n")
		return nil
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("request failed: %s", err)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}
	if err != nil {
		return fmt.Errorf("failed to read body: %s", err)
	}

	var locations config.LocationsArea
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return fmt.Errorf("error unmarshalling data: %s", err)
	}

	c.Next = locations.Next
	c.Previous = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
