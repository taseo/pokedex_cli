package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/taseo/pokedexcli/internal/api"
)

func commandCatch(config *Config, arg string) error {
	if arg == "" {
		return errors.New("Pokémon name was not provided. Usage: catch <pokémon-name>")
	}

	fmt.Printf("Throwing a Pokeball at %s ...\n", arg)

	url := "https://pokeapi.co/api/v2/pokemon/" + arg

	pokemon, err := api.PokeApiGet[api.Pokemon](url, config.Cache)

	if err != nil {
		return err
	}

	rand.Seed(time.Now().UnixNano())

	chance := rand.Float64()
	catchRate := float64(pokemon.BaseExperience) / 420.0

	if catchRate > 1.0 {
		catchRate = 1.0
	}

	if chance < catchRate {
		fmt.Printf("%s was caught\n", arg)

		config.Pokedex[arg] = pokemon

		return nil
	}

	fmt.Printf("%s escaped\n", arg)

	return nil
}
