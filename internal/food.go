package internal

type Food struct {
	Name     string  `json:"name"`
	Calories float64 `json:"calories"`
}

func AddFood(foods []Food, newFood Food) []Food {
	return append(foods, newFood)
}

func TotalCalories(foods []Food) float64 {
	total := 0.0

	for _, food := range foods {
		total += food.Calories
	}

	return total
}
