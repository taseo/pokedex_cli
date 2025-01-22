package api

import (
	"encoding/json"
	"io"
	"net/http"
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

func GetMap(url string) (LocationAreas, error) {
	areas := LocationAreas{}

	res, err := http.Get(url)

	if err != nil {
		return areas, err
	}

	body, err := io.ReadAll(res.Body)

	res.Body.Close()

	if res.StatusCode != 200 {
		return areas, err
	}

	if err != nil {
		return areas, err
	}

	err = json.Unmarshal(body, &areas)

	if err != nil {
		return areas, err
	}

	return areas, nil
}
