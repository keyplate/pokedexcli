package main

import (
    "github.com/keyplate/pokedexcli/internal/api"
    "math/rand"
    "fmt"
)

func commandCatch(args []string, config *Config) error {
    pokemonName := args[0]
    pokemon, err := api.FetchPokemon(pokemonName)
    if err != nil {
        return err
    }
    pokemonExp := pokemon.BaseExperience
    fmt.Printf("Trhowing a Pokeball at %s...\n", pokemonName)
    isCatched := catch(pokemonExp)
    if !isCatched {
        fmt.Printf("%s escaped\n", pokemonName)
        return nil
    } 
    config.pokedex[pokemonName] = pokemon
    fmt.Printf("%s is caught!\n", pokemonName)
    return nil
}

func catch(pokemonExp int) bool {
    var chance float64
    switch {
        case pokemonExp > 200:
            chance = 0.2
        case pokemonExp > 100:
            chance = 0.5
        default:
            chance = 0.8
    }
    
    randomChance := rand.Float64()
    return randomChance < chance
}
