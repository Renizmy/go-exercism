package lasagna

// PreparationTime calculate the time of preparation
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// Quantities calculate the portion of each ingredient
func Quantities(slice []string) (int, float64) {

	sauce := 0.0
	noodles := 0
	for i := 0; i < len(slice); i++ {
		switch slice[i] {
		case "sauce":
			sauce += 0.2
		case "noodles":
			noodles += 50
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(slice1 []string, slice2 []string) {
	slice2[len(slice2)-1] = slice1[len(slice1)-1]
}

func ScaleRecipe(ingredients []float64, number int) []float64 {
	var scaledIngredients []float64
	coef := float64(number) / 2.0
	for i := 0; i < len(ingredients); i++ {
		scaledIngredients = append(scaledIngredients, ingredients[i]*coef)
	}
	return scaledIngredients
}
