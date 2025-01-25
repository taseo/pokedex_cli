package main

import (
	"fmt"
	"os"
)

func commandExit(config *Config, arg string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")

	config.Cache.Delete()

	os.Exit(0)

	return nil
}
