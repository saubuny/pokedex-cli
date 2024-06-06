package pokeapi

import (
	"encoding/json"
	"github.com/saubuny/pokedex-cli/internal/pokecache"
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

func GetLocations(url string, cache *pokecache.Cache) (locationJson, error) {
	location := locationJson{}
	var err error

	data, ok := cache.Get(url)
	if !ok {
		data, err = fetch(url, cache)
		if err != nil {
			return location, err
		}
	}

	err = json.Unmarshal(data, &location)
	if err != nil {
		return location, err
	}

	return location, nil
}

func fetch(url string, cache *pokecache.Cache) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	cache.Add(url, data)
	return data, nil
}
