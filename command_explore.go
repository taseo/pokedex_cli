package main

import (
	"errors"
	"fmt"

	"github.com/taseo/pokedexcli/internal/api"
)

func commandExplore(config *Config, arg string) error {
	if arg == "" {
		return errors.New("Area name was not provided. Usage: explore <area-name>")
	}

	fmt.Printf("Exploring %s ...\n", arg)

	url := "https://pokeapi.co/api/v2/location-area/" + arg

	area, err := api.PokeApiGet[api.LocationArea](url, config.Cache)

	if err != nil {
		return err
	}

	fmt.Println("")
	fmt.Println("Found Pokemon:")

	for i := 0; i < len(area.PokemonEncounters); i++ {
		fmt.Printf(" - %s\n", area.PokemonEncounters[i].Pokemon.Name)
	}

	return nil
}
