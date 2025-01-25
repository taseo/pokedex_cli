package main

import (
	"time"

	"github.com/taseo/pokedexcli/internal/cache"
)

type Config struct {
	Next     *string
	Previous *string
	Cache    *cache.Cache
}

func main() {
	next := "https://pokeapi.co/api/v2/location-area"

	cache := cache.NewCache(42 * time.Second)

	config := Config{
		Next:  &next,
		Cache: cache,
	}

	startRepl(&config)
}
