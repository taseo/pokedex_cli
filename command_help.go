package main

import (
	"fmt"
)

func commandHelp(config *Config, arg string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for k, v := range commands {
		fmt.Printf("%s: %s\n", k, v.description)
	}

	return nil
}
