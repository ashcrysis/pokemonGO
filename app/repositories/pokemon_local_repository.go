package repositories

import (
	"app/models"
	"errors"

	"gorm.io/gorm"
)

type PokemonLocalRepository struct {
	DB *gorm.DB
}

func (repo *PokemonLocalRepository) FetchAllAsJSON() map[string]interface{} {
	pokemons, err := new(models.Pokemon).All(repo.DB)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	return map[string]interface{}{
		"count":    len(pokemons),
		"next":     nil,
		"previous": nil,
		"results":  pokemons,
	}
}

func (repo *PokemonLocalRepository) FetchPokemonSpeciesData(name string) (map[string]interface{}, error) {
	pokemon, err := new(models.Pokemon).FindByName(repo.DB, name)
	if err != nil {
		return nil, errors.New("Pokémon not found")
	}

	return pokemon.AsJSON(), nil
}
