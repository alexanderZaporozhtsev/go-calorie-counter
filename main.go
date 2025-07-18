package main

import (
	"fmt"
	"log"

	"github.com/alexanderZaporozhtsev/go-calorie-counter/internal"
)

func main() {
	foods := []internal.Food{
		{Name: "Яблоко", Calories: 52},
	}

	err := internal.SaveFoods(foods, "data.json")
	if err != nil {
		log.Fatal(err)
	}

	loadedFoods, err := internal.LoadFoods("data.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(loadedFoods)

}
