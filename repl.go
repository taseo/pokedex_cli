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
	callback    func(config *Config) error
}

var commands map[string]cliCommand

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl(config *Config) {
	commands = map[string]cliCommand{
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
			description: "Displays names of 20 location areas. Each subsequent call to map displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays names of 20 location areas. Each subsequent call to mapb displays the previous 20 locations",
			callback:    commandMapB,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()

		input := cleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}

		command, ok := commands[input[0]]

		if !ok {
			fmt.Println("Unknown command")

			continue
		}

		fmt.Println("")

		err := command.callback(config)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("")
	}
}
