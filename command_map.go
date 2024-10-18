package main 

import (
    "fmt"
    "github.com/keyplate/pokedexcli/internal/api"
)

func commandMap(args []string, config *Config) error {
    url := config.next

    locations, err := api.FetchMaps(url)
    if err != nil {
        return err
    }

    for _, result := range locations.Results {
        fmt.Println(result.Name)
    }
    config.next = locations.Next
    config.prev = locations.Previous

    return nil
} 
