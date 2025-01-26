package main

import (
	"fmt"
)

func commandPokedex(config *Config, arg string) error {
	if len(config.Pokedex) == 0 {
		fmt.Println("Your Pokedex is empty. Use catch <pokémon-name> command to catch and add Pokémon to your Pokedex")

		return nil
	}

	fmt.Println("Your Pokedex:")

	for k := range config.Pokedex {
		fmt.Printf(" - %s\n", config.Pokedex[k].Name)
	}

	return nil
}
