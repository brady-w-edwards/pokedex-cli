package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListAreaInfo(area string) (AreaInformation, error) {
	url := baseURL + "/location-area/" + area
	// if pageURL != nil {
	// 	url = *pageURL
	// }

	if val, ok := c.cache.Get(url); ok {
		locationsResp := AreaInformation{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return AreaInformation{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AreaInformation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AreaInformation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return AreaInformation{}, err
	}

	locationsResp := AreaInformation{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return AreaInformation{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
