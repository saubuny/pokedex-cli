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

			if cmd.name == "explore" {
				if len(input) != 2 {
					fmt.Println("Command `explore` requires a single location argument")
					break
				}
				config.location = &input[1]
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
