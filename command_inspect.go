package main

import (
    "fmt"
    "encoding/json"
)

func commandInspect(args []string, config *Config) error {
    if len(args) != 1 {
        return fmt.Errorf("You didnt' provide pokemon name\n")
    }
    pokemonName := args[0]
    pokemon, ok := config.pokedex[pokemonName]
    if !ok {
        return fmt.Errorf("You haven't caught %s yet\n", pokemonName)
    }
    pokemonString, _ := json.MarshalIndent(pokemon, "", "\t")
    fmt.Printf("s%\n", string(pokemonString))
    return nil
}
