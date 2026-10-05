package lasagnamaster

const (
    defaultPortions float64 = 2.0
    defaultTimePerLayer int = 2
    noodlesPerLayer int = 50
    saucePerLayer float64 = 0.2
)

func PreparationTime(layers []string, timePerLayer int) int {
    if timePerLayer == 0 {
        timePerLayer = defaultTimePerLayer
    }
    return len(layers) * timePerLayer
}

func Quantities(layers []string) (noodles int, sauce float64) {
    for i := 0; i < len(layers); i++ {
        if layers[i] == "noodles" {
            noodles += noodlesPerLayer
        } else if layers[i] == "sauce" {
            sauce += saucePerLayer
        }
    }
    return
}

func AddSecretIngredient(friendRecipe []string, myRecipe []string) {
    myRecipe[len(myRecipe) - 1] = friendRecipe[len(friendRecipe) - 1]
}

func ScaleRecipe(amounts []float64, portions int) []float64 {
    scaledAmounts := make([]float64, len(amounts))
    for i := 0; i < len(amounts); i++ {
        scaledAmounts[i] = amounts[i] / defaultPortions * float64(portions)
    }
    return scaledAmounts
}
