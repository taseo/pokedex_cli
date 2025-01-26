package main

import (
	"errors"
	"fmt"
)

func commandInspect(config *Config, arg string) error {
	if arg == "" {
		return errors.New("Pokémon name was not provided. Usage: inspect <pokémon-name>")
	}

	v, ok := config.Pokedex[arg]

	if ok {
		fmt.Printf("Name: %s\n", v.Name)
		fmt.Printf("Height: %d\n", v.Height)
		fmt.Printf("Weight: %d\n", v.Weight)

		fmt.Println("Stats:")

		for i := 0; i < len(v.Stats); i++ {
			fmt.Printf(" - %s: %d\n", v.Stats[i].Stat.Name, v.Stats[i].BaseStat)
		}

		fmt.Println("Types:")

		for i := 0; i < len(v.Types); i++ {
			fmt.Printf(" - %s\n", v.Types[i].Type.Name)
		}

		return nil
	}

	fmt.Println("You have not caught that pokemon")

	return nil
}
