package services

import (
	"app/repositories"
)

type PokemonService struct {
	Repository repositories.PokemonRepository
}

func NewPokemonService(repo repositories.PokemonRepository) *PokemonService {
	return &PokemonService{
		Repository: repo,
	}
}

func (ps *PokemonService) AllPokemons() map[string]interface{} {
	return ps.Repository.FetchAllAsJSON()
}

func (ps *PokemonService) FetchPokemonSpeciesData(name string) (map[string]interface{}, error) {
	data, err := ps.Repository.FetchPokemonSpeciesData(name)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (ps *PokemonService) FindEnglishDescription(data map[string]interface{}) string {
	flavorTexts, ok := data["flavor_text_entries"].([]interface{})
	if !ok {
		return "Description not found"
	}

	for _, entry := range flavorTexts {
		entryMap := entry.(map[string]interface{})
		language := entryMap["language"].(map[string]interface{})
		if language["name"].(string) == "en" {
			return entryMap["flavor_text"].(string)
		}
	}
	return "Description not found"
}
