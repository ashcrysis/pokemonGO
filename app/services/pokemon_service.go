package services

import (
	"app/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type PokemonService struct {
	repo repositories.PokemonApiRepository
}

func NewPokemonService(repo repositories.PokemonApiRepository) *PokemonService {
	return &PokemonService{repo: repo}
}

const cacheFilePath = "cached_pokemons.json"

func (ps *PokemonService) AllPokemons() []map[string]interface{} {
	if _, err := os.Stat(cacheFilePath); err == nil {
		return loadCachedData(cacheFilePath)
	}

	pokemons := ps.fetchAndCachePokemons()
	return pokemons
}

func (ps *PokemonService) fetchAndCachePokemons() []map[string]interface{} {
	pokemonData := ps.repo.FetchAllAsJSON()

	if results, ok := pokemonData["results"].([]interface{}); ok {
		pokemons := make([]map[string]interface{}, len(results))
		for i, result := range results {
			pokemons[i] = result.(map[string]interface{})
		}

		cacheData(cacheFilePath, pokemonData) 
		return pokemons
	}

	return nil
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

func loadCachedData(filePath string) []map[string]interface{} {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	var data map[string]interface{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Error unmarshalling data:", err)
		return nil
	}

	if results, ok := data["results"].([]interface{}); ok {
		pokemons := make([]map[string]interface{}, len(results))
		for i, result := range results {
			pokemons[i] = result.(map[string]interface{})
		}
		return pokemons
	}

	return nil
}

func (ps *PokemonService) FetchPokemonSpeciesData(name string) (map[string]interface{}, error) {
	return ps.repo.FetchPokemonSpeciesData(name)
}

func (ps *PokemonService) FindEnglishDescription(data map[string]interface{}) string {
	if flavorTextEntries, ok := data["flavor_text_entries"].([]interface{}); ok {
		for _, entry := range flavorTextEntries {
			entryMap := entry.(map[string]interface{})
			if language, ok := entryMap["language"].(map[string]interface{}); ok {
				if langName, ok := language["name"].(string); ok && langName == "en" {
					if description, ok := entryMap["flavor_text"].(string); ok {
						return description
					}
				}
			}
		}
	}
	return "No English description available."
}
