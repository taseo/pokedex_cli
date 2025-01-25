package api

import (
	"io"
	"net/http"

	"github.com/taseo/pokedexcli/internal/cache"
)

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetMap(url string, cache *cache.Cache) (LocationAreas, error) {
	areas := LocationAreas{}

	v, ok := cache.Get(url)

	if ok {
		return ParseData[LocationAreas](v, &areas)
	}

	res, err := http.Get(url)

	if err != nil {
		return areas, err
	}

	body, err := io.ReadAll(res.Body)

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return areas, err
	}

	if err != nil {
		return areas, err
	}

	cache.Add(url, body)

	return ParseData[LocationAreas](body, &areas)
}
