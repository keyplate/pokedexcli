package api 

import (
    "fmt"
    "net/http"
    "encoding/json"
)

var baseUrl string = "https://pokeapi.co/api/v2"
var cache 
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

    res, err := http.Get(fullUrl)
    if err != nil {
        return LocationListRes{}, fmt.Errorf("Fetching maps returned an error: %v", err)
    }

    locations := LocationListRes{}
    dec := json.NewDecoder(res.Body)
    err = dec.Decode(&locations)

    if err != nil {
        return LocationListRes{}, fmt.Errorf("Decoding maps was unsuccessful: %v", err)
    }
    return locations, nil
}
