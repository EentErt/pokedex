package Requests

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

var cache = pokecache.NewCache(5 * time.Minute)

func MakeRequest(url string) error {
	fmt.Println("Fetching data from:", url)

	// Check if the data is cached
	cachedData, ok := cache.Get(url)
	// If not cached, make the HTTP request
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		} else if res.StatusCode != http.StatusOK {
			return fmt.Errorf("error: %s", res.Status)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cache.Add(url, body)
		cachedData = body
	}

	err := json.Unmarshal(cachedData, &JsonMapData)
	if err != nil {
		return err
	}
	return nil
}

func ExploreRequest(arg string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + arg + "/"

	fmt.Printf("Exploring %s...\n", arg)

	// Check if the data is cached
	cachedData, ok := cache.Get(url)
	// If not cached, make the HTTP request
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		} else if res.StatusCode != http.StatusOK {
			return fmt.Errorf("error: %s", res.Status)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cache.Add(url, body)
		cachedData = body
	}

	err := json.Unmarshal(cachedData, &JsonExploreData)
	if err != nil {
		return err
	}
	return nil
}

func CatchRequest(arg string) error {
	url := "https://pokeapi.co/api/v2/pokemon/" + arg + "/"

	fmt.Printf("Throwing a Pokeball at %s...\n", arg)

	// Check if the data is cached
	cachedData, ok := cache.Get(url)
	// If not cached, make the HTTP request
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		} else if res.StatusCode != http.StatusOK {
			return fmt.Errorf("error: %s", res.Status)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cache.Add(url, body)
		cachedData = body
	}

	err := json.Unmarshal(cachedData, &JsonPokemonData)
	if err != nil {
		return err
	}

	// Try to catch the Pokemon
	catchChance := 50
	if rand.Intn(JsonPokemonData.BaseExperience) > catchChance {
		err = fmt.Errorf("%s escaped!", arg)
		return err
	} else {
		return nil
	}
}
