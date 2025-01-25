package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/taseo/pokedexcli/internal/cache"
)

type PokeApiModel interface{}

func PokeApiGet[T PokeApiModel](url string, cache *cache.Cache) (T, error) {
	var data T

	v, ok := cache.Get(url)

	if ok {
		return ParseData[T](v, &data)
	}

	res, err := http.Get(url)

	if err != nil {
		return data, err
	}

	body, err := io.ReadAll(res.Body)

	defer res.Body.Close()

	if res.StatusCode == 404 {
		return data, errors.New("Not Found")
	}

	if res.StatusCode != 200 {
		return data, fmt.Errorf("API Error: %d\n", res.StatusCode)
	}

	if err != nil {
		return data, err
	}

	cache.Add(url, body)

	return ParseData[T](body, &data)
}
