package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	commands := getCommands()
	for _, cmd := range commands {
		fmt.Printf("%v -> %v\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit() error {
	return errors.New("Exit")
}

fun commandMap() error {

}

fun commandMapb() error {

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
			callback:    commandMap,
		},
		"mapb": {
			name:        "map",
			description: "Display the names of the previous 20 locations in the Pokemon world",
			callback:    commandMapb,
		},
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
mainloop:
	for {
		fmt.Print("pokemon > ")
		for scanner.Scan() {
			cmd, ok := getCommands()[scanner.Text()]
			if ok {
				err := cmd.callback()
				if err != nil && err.Error() == "Exit" {
					break mainloop
				}
			} else {
				fmt.Println("Invalid command, use `help` to see available commands")
			}
			break
		}
	}
}
