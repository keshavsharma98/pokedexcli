package poke_apis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

type LocationAreaResponse struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var url string = "https://pokeapi.co/api/v2/location-area"

func NewCLient() *Client {
	return &Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) GetLocations(requestURL *string) (LocationAreaResponse, error) {
	if requestURL != nil {
		url = *requestURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	if res.StatusCode > 399 {
		return LocationAreaResponse{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	var responseBody LocationAreaResponse
	if err := json.Unmarshal(body, &responseBody); err != nil {
		return LocationAreaResponse{}, err
	}

	return responseBody, nil
}
