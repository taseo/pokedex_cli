package main

import (
	"errors"
	"fmt"

	"github.com/taseo/pokedexcli/internal/api"
)

func commandMap(config *Config) error {
	if config.Next == nil {
		return errors.New("You're on the last page")
	}

	areas, err := api.GetMap(*config.Next, config.Cache)

	if err != nil {
		return err
	}

	config.Next = areas.Next
	config.Previous = areas.Previous

	for i := 0; i < len(areas.Results); i++ {
		fmt.Println(areas.Results[i].Name)
	}

	return nil
}

func commandMapB(config *Config) error {
	if config.Previous == nil {
		return errors.New("You're on the first page")
	}

	areas, err := api.GetMap(*config.Previous, config.Cache)

	if err != nil {
		return err
	}

	config.Next = areas.Next
	config.Previous = areas.Previous

	for i := 0; i < len(areas.Results); i++ {
		fmt.Println(areas.Results[i].Name)
	}

	return nil
}
