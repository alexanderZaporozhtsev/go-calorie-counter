package internal

import (
	"encoding/json"
	"os"
)

func SaveFoods(foods []Food, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(foods)
}

func LoadFoods(filename string) ([]Food, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Food{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var foods []Food
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&foods)
	return foods, err
}
