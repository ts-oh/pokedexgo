package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaRes, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaRes{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaRes{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaRes{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaRes{}, err
	}

	locationAreaRes := LocationAreaRes{}
	err = json.Unmarshal(dat, &locationAreaRes)
	if err != nil {
		return LocationAreaRes{}, err
	}

	return locationAreaRes, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	return locationArea, nil
}
