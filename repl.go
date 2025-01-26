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
	callback    func(config *Config, arg string) error
}

var commands map[string]cliCommand

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl(config *Config) {
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas. Each subsequent call to 'map' shows the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of 20 location areas. Each subsequent call to 'mapb' shows the previous 20 locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Explores an area and returns the Pokémon found there. Usage: explore <area-name>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Makes an attempt to catch Pokémon. Usage: catch <pokémon-name>",
			callback:    commandCatch,
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

		arg := ""

		if len(input) > 1 {
			arg = input[1]
		}

		fmt.Println("")

		err := command.callback(config, arg)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("")
	}
}
