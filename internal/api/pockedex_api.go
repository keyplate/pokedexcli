package api 

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/keyplate/pokedexcli/internal/cache"
    "io"
)

var baseUrl string = "https://pokeapi.co/api/v2"
var pokeCache = cache.NewCache(10)

type LocationAreaPageResp struct {
    Count    int    `json:"count"`
    Next     *string `json:"next"`
    Previous *string    `json:"previous"`
    Results  []struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"results"`
}

func FetchMaps(url *string) (LocationAreaPageResp, error) {
    fullUrl := baseUrl + "/location-area"
    if url != nil {
        fullUrl = *url
    }

    dat, ok := pokeCache.Get(fullUrl)
    if ok {
        locationAreaPage := LocationAreaPageResp{}
        err := json.Unmarshal(dat, &locationAreaPage)
        if err != nil {
            return LocationAreaPageResp{}, err
        }
        return locationAreaPage, nil
    }


    res, err := http.Get(fullUrl)
    if err != nil {
        return LocationAreaPageResp{}, fmt.Errorf("Fetching maps returned an error: %v", err)
    }

    dat, err = io.ReadAll(res.Body)
    if err != nil {
        return LocationAreaPageResp{}, err
    }

    locationAreaPage := LocationAreaPageResp{}
    err = json.Unmarshal(dat, &locationAreaPage)

    if err != nil {
        return LocationAreaPageResp{}, fmt.Errorf("Decoding maps was unsuccessful: %v", err)
    }

    pokeCache.Add(fullUrl, dat)

    return locationAreaPage, nil
}

type LocationAreaResp struct {
    ID                   int    `json:"id"`
    Name                 string `json:"name"`
    GameIndex            int    `json:"game_index"`
    EncounterMethodRates *string `json:"encounter_method_rates"`
    Location *string `json:"location"`
    Names *string `json:"names"`
    PokemonEncounters []struct {
        Pokemon struct {
            Name string `json:"name"`
            URL  string `json:"url"`
        } `json:"pokemon"`
        VersionDetails *string 
    } `json:"pokemon_encounters"` 
}

func FetchLocationArea(locationArea string) (LocationAreaResp, error) {
    fullUrl := baseUrl + "/location-area/" + locationArea

    dat, ok := pokeCache.Get(fullUrl)
    if ok {
        locationAreaResp := LocationAreaResp{}
        err := json.Unmarshal(dat, &locationAreaResp)
        if err != nil {
            return LocationAreaResp{}, err
        }
        return locationAreaResp, nil
    }

    res, err := http.Get(fullUrl)
    if err != nil {
        return LocationAreaResp{}, fmt.Errorf("Fetching location return an error: %v", err)
    }

    dat, err = io.ReadAll(res.Body)
    if err != nil {
        return LocationAreaResp{}, err
    } 

    locationAreaResp := LocationAreaResp{}
    err = json.Unmarshal(dat, &locationAreaResp)
    if err != nil {
        return LocationAreaResp{}, fmt.Errorf("Decoding location datat returned an error: %v", err)
    }

    pokeCache.Add(fullUrl, dat)

    return locationAreaResp, nil
}
