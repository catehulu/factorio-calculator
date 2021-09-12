package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

type Recipe struct {
	Name        string        `json:"name"`
	Energy      float32       `json:"energy"`
	Ingredients []ingredients `json:"ingredients"`
	Products    []products    `json:"products"`
	MainProduct products      `json:"main_product"`
}
type products struct {
	Type        string  `json:"type"`
	Name        string  `json:"name"`
	Probability float32 `json:"probability"`
	Amount      float32 `json:"amount"`
}

type ingredients struct {
	Type   string  `json:"type"`
	Name   string  `json:"name"`
	Amount float32 `json:"amount"`
}

var recipeList map[string]Recipe

func InitRecipe(path string) {
	f, err := os.Open(path + "/assets/data/recipe.json")
	if err != nil {
		log.Fatal("Recipe file not found")
	}

	defer f.Close()
	byteValue, _ := ioutil.ReadAll(f)

	json.Unmarshal([]byte(byteValue), &recipeList)
}

func GetRecipe(id string) (Recipe, error) {
	r, ok := recipeList[id]
	if !ok {
		return Recipe{}, errors.New("no recipe found")
	} else {
		return r, nil
	}
}

func GetAllRecipe() map[string]Recipe {
	return recipeList
}
