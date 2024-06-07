package main

import (
	"errors"
	"fmt"
	"github.com/saubuny/pokedex-cli/internal/pokeapi"
)

func commandHelp(config *cfg) error {
	commands := getCommands()
	fmt.Println("Commands:")
	for _, cmd := range commands {
		fmt.Printf("\t%v -> %v\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit(config *cfg) error {
	return errors.New("Exit")
}

func commandMapf(config *cfg) error {
	locations, err := pokeapi.GetLocations(*config.nextPage, config.cache)
	if err != nil {
		return err
	}

	config.nextPage = &locations.Next
	config.prevPage = &locations.Previous

	fmt.Println("Locations:")
	for _, location := range locations.Results {
		fmt.Println("\t-", location.Name)
	}

	return nil
}

func commandMapb(config *cfg) error {
	if config.prevPage == nil || *config.prevPage == "" {
		return errors.New("No Previous Map Available")
	}
	locations, err := pokeapi.GetLocations(*config.prevPage, config.cache)
	if err != nil {
		return err
	}

	config.nextPage = &locations.Next
	config.prevPage = &locations.Previous

	fmt.Println("Locations:")
	for _, location := range locations.Results {
		fmt.Println("\t-", location.Name)
	}

	return nil
}

func commandExplore(config *cfg) error {
	locationInfo, err := pokeapi.GetLocationInfo(*config.location, config.cache)
	if err != nil {
		return err
	}

	fmt.Println("Exploring", *config.location)
	fmt.Println("Found Pokemon:")
	for _, encounter := range locationInfo.PokemonEncounters {
		fmt.Println("\t-", encounter.Pokemon.Name)
	}

	return nil
}
