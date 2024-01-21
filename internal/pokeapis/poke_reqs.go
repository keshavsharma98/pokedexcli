package pokeapis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var url string = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
var locationAreaUrl string = "https://pokeapi.co/api/v2/location-area/"
var getPokemonUrl string = "https://pokeapi.co/api/v2/pokemon/"

func (c *Client) GetLocations(requestURL *string) (LocationsApiResponse, error) {
	if requestURL != nil {
		url = *requestURL
	}

	cached, ok := c.cache.Get(url)
	if ok {
		response := LocationsApiResponse{}
		if err := json.Unmarshal(cached, &response); err != nil {
			return LocationsApiResponse{}, err
		}
		return response, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationsApiResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationsApiResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationsApiResponse{}, err
	}

	if res.StatusCode > 399 {
		return LocationsApiResponse{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	var responseBody LocationsApiResponse
	if err := json.Unmarshal(body, &responseBody); err != nil {
		return LocationsApiResponse{}, err
	}

	c.cache.Add(url, body)
	return responseBody, nil
}

func (c *Client) GetPokemonsInArea(areaName string) (LocationsAreaApiResponse, error) {

	cached, ok := c.cache.Get(areaName)
	if ok {
		response := LocationsAreaApiResponse{}
		if err := json.Unmarshal(cached, &response); err != nil {
			return LocationsAreaApiResponse{}, err
		}
		return response, nil
	}

	req, err := http.NewRequest("GET", locationAreaUrl+areaName, nil)
	if err != nil {
		return LocationsAreaApiResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationsAreaApiResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationsAreaApiResponse{}, err
	}

	if res.StatusCode > 399 {
		if res.StatusCode == 404 {
			return LocationsAreaApiResponse{}, fmt.Errorf("invalid location")
		}
		return LocationsAreaApiResponse{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	var responseBody LocationsAreaApiResponse
	if err := json.Unmarshal(body, &responseBody); err != nil {
		return LocationsAreaApiResponse{}, err
	}

	c.cache.Add(areaName, body)
	return responseBody, nil
}

func (c *Client) GetPokemonInfo(pokemonName string) (PokemonInfoApiResponse, error) {
	cached, ok := c.cache.Get(pokemonName)
	if ok {
		response := PokemonInfoApiResponse{}
		if err := json.Unmarshal(cached, &response); err != nil {
			return PokemonInfoApiResponse{}, err
		}
		return response, nil
	}

	req, err := http.NewRequest("GET", getPokemonUrl+pokemonName, nil)
	if err != nil {
		return PokemonInfoApiResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfoApiResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonInfoApiResponse{}, err
	}

	if res.StatusCode > 399 {
		if res.StatusCode == 404 {
			return PokemonInfoApiResponse{}, fmt.Errorf("invalid location")
		}
		return PokemonInfoApiResponse{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	var responseBody PokemonInfoApiResponse
	if err := json.Unmarshal(body, &responseBody); err != nil {
		return PokemonInfoApiResponse{}, err
	}
	c.cache.Add(pokemonName, body)
	return responseBody, nil
}
