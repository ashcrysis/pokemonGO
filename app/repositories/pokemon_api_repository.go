package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	POKEMON_API      = "https://pokeapi.co/api/v2/pokemon"
	CacheFilePath    = "cached_pokemons.json" 
)

type PokemonApiRepository struct {
	Limit int
}

func (repo *PokemonApiRepository) FetchAllAsJSON() map[string]interface{} {
	if _, err := os.Stat(CacheFilePath); err == nil {
		return loadCachedData(CacheFilePath)
	}

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

	response := map[string]interface{}{
		"count":   len(allPokemons),
		"results": allPokemons,
	}

	cacheData(CacheFilePath, response)

	return response
}

func cacheData(filePath string, data map[string]interface{}) {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}

	err = ioutil.WriteFile(filePath, file, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func loadCachedData(filePath string) map[string]interface{} {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return map[string]interface{}{"error": "Cache not found"}
	}

	var data map[string]interface{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Error unmarshalling data:", err)
		return map[string]interface{}{"error": "Invalid cache format"}
	}

	return data
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