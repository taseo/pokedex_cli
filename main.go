package main

import (
	"time"

	"github.com/taseo/pokedexcli/internal/api"
	"github.com/taseo/pokedexcli/internal/cache"
)

type Config struct {
	Next     *string
	Previous *string
	Cache    *cache.Cache
	Pokedex  map[string]api.Pokemon
}

func main() {
	next := "https://pokeapi.co/api/v2/location-area"
	cache := cache.NewCache(42 * time.Second)
	pokedex := make(map[string]api.Pokemon)

	config := Config{
		Next:    &next,
		Cache:   cache,
		Pokedex: pokedex,
	}

	startRepl(&config)
}
