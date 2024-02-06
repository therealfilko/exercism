package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTimePerLayer int) int {
    if prepTimePerLayer == 0 {
        prepTimePerLayer = 2 // Verwende den Standardwert, falls 0 übergeben wird
    }
    return len(layers) * prepTimePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64){
    sumNoodles := 0
    sumSauce := 0.0

    for i := 0; i < len(layers); i++ {
        if layers[i] == "noodles" {
            sumNoodles += 50
        } else if layers[i] == "sauce" {
            sumSauce += 0.2
        }
    }
    return sumNoodles, sumSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
    if len(friendsList) > 0 {
        // Setze das letzte Element von myList auf das letzte Element von friendsList
        myList[len(myList)-1] = friendsList[len(friendsList)-1]
    }
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantitiesForTwo []float64, portions int) []float64 {
    scaleFactor := float64(portions) / 2
    scaleQuantities := make([]float64, len(quantitiesForTwo))

    for i, quantity := range quantitiesForTwo{
        scaleQuantities[i] = quantity * scaleFactor
    }

    return scaleQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
