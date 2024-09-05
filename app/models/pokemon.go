package models

import (
	"gorm.io/gorm"
)

type Pokemon struct {
	ID     uint   `gorm:"primaryKey"`
	Nome   string `gorm:"unique"`
	Moves  string
	Tipo   string
	Imagem string
}

func (p *Pokemon) FindByName(db *gorm.DB, name string) (*Pokemon, error) {
	var pokemon Pokemon
	if err := db.Where("nome = ?", name).First(&pokemon).Error; err != nil {
		return nil, err
	}
	return &pokemon, nil
}

func (p *Pokemon) All(db *gorm.DB) ([]Pokemon, error) {
	var pokemons []Pokemon
	if err := db.Find(&pokemons).Error; err != nil {
		return nil, err
	}
	return pokemons, nil
}

func (p *Pokemon) AsJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":     p.ID,
		"nome":   p.Nome,
		"moves":  p.Moves,
		"tipo":   p.Tipo,
		"imagem": p.Imagem,
	}
}
