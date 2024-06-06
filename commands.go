package main

import (
	"errors"
	"fmt"
	"github.com/saubuny/pokedex-cli/internal/pokeapi"
)

func commandHelp(config *cfg) error {
	commands := getCommands()
	for _, cmd := range commands {
		fmt.Printf("%v -> %v\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit(config *cfg) error {
	return errors.New("Exit")
}

func commandMapf(config *cfg) error {
	locations, err := pokeapi.GetLocations(*config.nextPage)
	if err != nil {
		return err
	}

	config.nextPage = &locations.Next
	config.prevPage = &locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(config *cfg) error {
	if config.prevPage == nil || *config.prevPage == "" {
		return errors.New("No Previous Map Available")
	}
	locations, err := pokeapi.GetLocations(*config.prevPage)
	if err != nil {
		return err
	}

	config.nextPage = &locations.Next
	config.prevPage = &locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
