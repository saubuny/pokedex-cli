package main

import (
	"github.com/saubuny/pokedex-cli/internal/pokecache"
	"time"
)

type cfg struct {
	cache    *pokecache.Cache
	nextPage *string
	prevPage *string
	location *string
}

func main() {
	baseUrl := "https://pokeapi.co/api/v2/location-area/"
	config := cfg{
		pokecache.NewCache(time.Duration(5 * time.Second)),
		&baseUrl,
		nil,
		nil,
	}

	startRepl(&config)
}
