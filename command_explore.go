package main

import (
    "fmt"
    "github.com/keyplate/pokedexcli/internal/api"
)

func commandExplore(args []string, config *Config) error {
    locationAreaName := args[0]
    fmt.Printf("Exploring %s...\n", locationAreaName)

    locationArea, err := api.FetchLocationArea(locationAreaName)
    if err != nil {
        return err
    }

    fmt.Println("Found Pokemon: ")
    for _, pokemon := range locationArea.PokemonEncounters {
        fmt.Println(pokemon.Pokemon.Name)
    }

    return nil
}
