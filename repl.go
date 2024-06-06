package main

import (
	"bufio"
	"fmt"
	"os"
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
	}
}

func startRepl(config *cfg) {
	scanner := bufio.NewScanner(os.Stdin)
mainloop:
	for {
		fmt.Print("pokemon > ")
		for scanner.Scan() {
			cmd, ok := getCommands()[scanner.Text()]
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
