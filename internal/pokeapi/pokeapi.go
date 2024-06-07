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

type PokemonI struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
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

func GetLocationInfo(name string, cache *pokecache.Cache) (PokemonI, error) {
	locationInfo := PokemonI{}
	var err error
	url := "https://pokeapi.co/api/v2/location-area/" + name

	data, ok := cache.Get(url)
	if !ok {
		data, err = fetch(url, cache)
		if err != nil {
			return locationInfo, err
		}
	}

	err = json.Unmarshal(data, &locationInfo)
	if err != nil {
		return locationInfo, err
	}

	return locationInfo, nil
}

// NOTE: We can probably combine all these functions and only change the url and type of data based on a switch case or something
func GetPokemonInfo(name string, cache *pokecache.Cache) (Pokemon, error) {
	pokemon := Pokemon{}
	var err error
	url := "https://pokeapi.co/api/v2/pokemon/" + name

	data, ok := cache.Get(url)
	if !ok {
		data, err = fetch(url, cache)
		if err != nil {
			return pokemon, err
		}
	}

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return pokemon, err
	}

	return pokemon, nil
}

type Pokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience         int    `json:"base_experience"`
	Height                 int    `json:"height"`
	HeldItems              []any  `json:"held_items"`
	ID                     int    `json:"id"`
	IsDefault              bool   `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
	} `json:"moves"`
	Name          string `json:"name"`
	Order         int    `json:"order"`
	PastAbilities []any  `json:"past_abilities"`
	PastTypes     []any  `json:"past_types"`
	Sprites       struct {
		BackDefault      string `json:"back_default"`
		BackFemale       any    `json:"back_female"`
		BackShiny        string `json:"back_shiny"`
		BackShinyFemale  any    `json:"back_shiny_female"`
		FrontDefault     string `json:"front_default"`
		FrontFemale      any    `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale any    `json:"front_shiny_female"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}
