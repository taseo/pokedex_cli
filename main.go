package main

type Config struct {
	Next     *string
	Previous *string
}

func main() {
	next := "https://pokeapi.co/api/v2/location-area"

	config := Config{
		Next: &next,
	}

	startRepl(&config)
}
