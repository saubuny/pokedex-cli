package main

import (
	"time"

	"github.com/saubuny/pokedex-cli/internal/pokeapi"
	"github.com/saubuny/pokedex-cli/internal/pokecache"
)

type cfg struct {
	cache    *pokecache.Cache
	nextPage *string
	prevPage *string
	arg      *string
	pokedex  map[string]pokeapi.Pokemon
}

func main() {
	baseUrl := "https://pokeapi.co/api/v2/location-area/"
	config := cfg{
		pokecache.NewCache(time.Duration(5 * time.Second)),
		&baseUrl,
		nil,
		nil,
		map[string]pokeapi.Pokemon{},
	}

	startRepl(&config)
}
