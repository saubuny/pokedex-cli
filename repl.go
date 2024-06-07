package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*cfg) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the names of the next 20 locations in the Pokemon world",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of the previous 20 locations in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a given pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a given pokemon that you have caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List the pokemon in your pokedex",
			callback:    commandPokedex,
		},
	}
}

func startRepl(config *cfg) {
	scanner := bufio.NewScanner(os.Stdin)
mainloop:
	for {
		fmt.Print("pokemon > ")
		for scanner.Scan() {
			input := strings.Split(scanner.Text(), " ")
			cmd, ok := getCommands()[input[0]]

			if cmd.name == "explore" || cmd.name == "catch" || cmd.name == "inspect" {
				if len(input) != 2 || input[1] == "" {
					fmt.Println("Command requires a single argument")
					break
				}
				config.arg = &input[1]
			}

			if ok {
				err := cmd.callback(config)
				if err != nil {
					if err.Error() == "Exit" {
						break mainloop
					}
					fmt.Println(err)
				}
			} else {
				fmt.Println("Invalid command, use `help` to see available commands")
			}
			break
		}
	}
}
