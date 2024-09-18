package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const POKEMON_API = "https://pokeapi.co/api/v2/pokemon"

type PokemonRepository interface {
	FetchAllAsJSON() map[string]interface{}
	FetchPokemonSpeciesData(name string) (map[string]interface{}, error)
}

type PokemonApiRepository struct {
	Limit int
}

func NewPokemonApiRepository() *PokemonApiRepository {
	return &PokemonApiRepository{
		Limit: 10000,
	}
}

func (repo *PokemonApiRepository) FetchAllAsJSON() map[string]interface{} {
	allPokemons := make([]interface{}, 0)
	nextURL := fmt.Sprintf("%s?limit=%d", POKEMON_API, repo.Limit)

	for nextURL != "" {
		resp, err := http.Get(nextURL)
		if err != nil {
			return map[string]interface{}{"error": err.Error()}
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return map[string]interface{}{"error": err.Error()}
		}

		var result map[string]interface{}
		json.Unmarshal(body, &result)

		if results, ok := result["results"].([]interface{}); ok {
			allPokemons = append(allPokemons, results...)
		}

		if next, ok := result["next"].(string); ok {
			nextURL = next
		} else {
			nextURL = "" 
		}
	}

	return map[string]interface{}{
		"count":   len(allPokemons),
		"results": allPokemons,
	}
}
func (repo *PokemonApiRepository) FetchPokemonSpeciesData(name string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s-species/%s", POKEMON_API, name)
	println(url)
	resp, err := http.Get(url)
	
	if err != nil || resp.StatusCode != 200 {
		return nil, fmt.Errorf("Pokémon not found")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)
	return result, nil
}
