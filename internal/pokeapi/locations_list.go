package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationsArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := LocationsArea{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationsArea{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationsArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationsArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationsArea{}, err
	}

	locationsResp := LocationsArea{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationsArea{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
