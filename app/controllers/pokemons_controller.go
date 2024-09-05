package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"app/repositories"
	"app/services"

	"github.com/gin-gonic/gin"
)

type PokemonController struct {
	UseSecondAPI   bool
	PokemonService *services.PokemonService 
}

func NewPokemonController() *PokemonController {
	return &PokemonController{
		UseSecondAPI:   true,
		PokemonService: services.NewPokemonService(&repositories.PokemonApiRepository{}),  
	}
}

func (pc *PokemonController) Search(c *gin.Context) {
	name := c.Param("name")

	if strings.TrimSpace(name) == "" || strings.ContainsAny(name, "0123456789") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Pokémon name"})
		return
	}

	data, err := pc.PokemonService.FetchPokemonSpeciesData(name)  
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pokémon not found"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (pc *PokemonController) FetchAllPokemonData(c *gin.Context) {
	response := pc.PokemonService.AllPokemons()  
	c.JSON(http.StatusOK, response)
}

func (pc *PokemonController) Species(c *gin.Context) {
	name := c.Param("name")
	data, err := pc.PokemonService.FetchPokemonSpeciesData(name) 
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	description := pc.PokemonService.FindEnglishDescription(data)  
	c.JSON(http.StatusOK, gin.H{"description": description})
}

func (pc *PokemonController) ToggleAPI(c *gin.Context) {
	pc.UseSecondAPI = !pc.UseSecondAPI
	message := fmt.Sprintf("API toggled to %v", pc.UseSecondAPI)
	c.JSON(http.StatusOK, gin.H{"message": message})
}
