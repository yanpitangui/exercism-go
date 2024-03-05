package lasagna

func PreparationTime(layers []string, averagePreparationTime int) int {
	if averagePreparationTime == 0 {
		averagePreparationTime = 2
	}
	return len(layers) * (averagePreparationTime)
}

func Quantities(layers []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		} else if layers[i] == "sauce" {
			sauce += 0.2
		}
	}
	return
}

func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, numberOfPortions int) []float64 {
	var scaledRecipe []float64
	factor := float64(numberOfPortions) / 2
	for i := 0; i < len(quantities); i++ {
		scaledRecipe = append(scaledRecipe, quantities[i]*factor)
	}
	return scaledRecipe

}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
