package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

type Item struct {
	Name          string   `json:"name"`
	LocalizedName []string `json:"localised_name"`
	ItemType      string   `json:"type"`
	Order         string   `json:"order"`
	FuelValue     int      `json:"fuel_value"`
	StackSize     int      `json:"stack_size"`
	PlaceResult   string   `json:"place_result"`
}

var itemList map[string]Item

func InitItem(path string) {
	f, err := os.Open(path + "/assets/data/item.json")
	if err != nil {
		log.Fatal("Item file not found")
	}

	defer f.Close()
	byteValue, _ := ioutil.ReadAll(f)

	json.Unmarshal([]byte(byteValue), &itemList)
}

func GetItem(id string) (Item, error) {
	i, ok := itemList[id]
	if !ok {
		return Item{}, errors.New("no item found")
	} else {
		return i, nil
	}
}
