# PokedexCLI

A REPL-style Pokedex application that uses the [PokéAPI](https://pokeapi.co) to fetch data about Pokémon

## Usage

- `go build && ./pokedexcli` will:
    - Compile the source code into an executable binary
    - Run the generated binary
- `go test ./...` will:
    - Run unit tests

## Supported commands

- `help`    Displays help message
- `exit`    Exits the Pokedex
- `map`     Displays the names of 20 location areas. Each subsequent call to `map` shows the next 20 locations
- `mapb`    Displays the names of 20 location areas. Each subsequent call to `mapb` shows the previous 20 locations
- `explore` Explores an area and returns the Pokémon found there. Usage: `explore <area-name>`
- `catch`   Makes an attempt to catch Pokémon. Usage: `catch <pokémon-name>`
- `inspect` Displays details about a Pokémon if it was caught before. Useage: `inspect <pokémon-name>`
- `pokedex` Displays list of all the names of the Pokémon that has been caught

## Acknowledgments

Based on a guided project found on [boot.dev](https://www.boot.dev)

