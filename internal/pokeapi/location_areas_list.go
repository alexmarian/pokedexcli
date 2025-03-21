package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	fmt.Println("Fetching locations...")

	resp, err := http.Get(url)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	var locationResponse LocationAreaResponse
	err = json.NewDecoder(resp.Body).Decode(&locationResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	return locationResponse, nil
}
