package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	if val, ok := c.cache.Get(url); ok {
		locationsResp := LocationAreasResponse{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationAreasResponse{}, err
		}

		return locationsResp, nil
	}

	fmt.Println("Fetching locations...")

	resp, err := http.Get(url)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer resp.Body.Close()

	locationResponseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationResponse := LocationAreasResponse{}
	err = json.Unmarshal(locationResponseData, &locationResponse)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Add(url, locationResponseData)
	return locationResponse, nil
}
