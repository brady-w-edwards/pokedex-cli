package pokeapi

import (
	"pokedex-cli/internal/pokecache"
	"time"
)

var cache *pokecache.Cache

func init() {
	// Initialize cache with a reasonable interval (e.g., 5 minutes)
	cache = pokecache.NewCache(5 * time.Minute)
}

func makeRequest(url string) ([]byte, error) {
	// Check if response is in cache first
	if data, found := cache.Get(url); found {
		return data, nil
	}

	// Otherwise make the actual request
	// ...

	// And store the response in cache before returning
	cache.Add(url, responseData)
	return responseData, nil
}
