package chance

import "math/rand"

var animals = []string{
	"ant",
	"beaver",
	"cat",
	"dog",
	"elephant",
	"fox",
	"giraffe",
	"hedgehog",
}

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return rand.Intn(19) + 2
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return rand.Float64() + float64(rand.Intn(12))
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := []string{
		"ant",
		"beaver",
		"cat",
		"dog",
		"elephant",
		"fox",
		"giraffe",
		"hedgehog",
	}

	rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})

	return animals
}
