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

type LocationListRes struct {
    Count    int    `json:"count"`
    Next     *string `json:"next"`
    Previous *string    `json:"previous"`
    Results  []struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"results"`
}

func FetchMaps(url *string) (LocationListRes, error) {
    fullUrl := baseUrl + "/location-area"
    if url != nil {
        fullUrl = *url
    }

    dat, ok := pokeCache.Get(fullUrl)
    if ok {
        locations := LocationListRes{}
        err := json.Unmarshal(dat, &locations)
        if err != nil {
            return LocationListRes{}, err
        }
        return locations, nil
    }


    res, err := http.Get(fullUrl)
    if err != nil {
        return LocationListRes{}, fmt.Errorf("Fetching maps returned an error: %v", err)
    }

    locations := LocationListRes{}
    dat, err = io.ReadAll(res.Body)
    if err != nil {
        return LocationListRes{}, err
    }

    err = json.Unmarshal(dat, &locations)

    if err != nil {
        return LocationListRes{}, fmt.Errorf("Decoding maps was unsuccessful: %v", err)
    }

    pokeCache.Add(fullUrl, dat)

    return locations, nil
}
