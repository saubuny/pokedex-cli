package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type locationJson struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocations(url string) (locationJson, error) {
	res, err := http.Get(url)
	location := locationJson{}
	if err != nil {
		return location, err
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return location, err
	}

	err = json.Unmarshal(data, &location)
	if err != nil {
		return location, err
	}

	return location, nil
}
