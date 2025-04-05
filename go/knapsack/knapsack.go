package knapsack

import "reflect"

type Item struct {
	Weight, Value int
}

type knapsack struct {
	weight, value int
	indices       map[int]bool
}

// Knapsack takes in a maximum carrying capacity and a collection of items
// and returns the maximum value that can be carried by the knapsack
// given that the knapsack can only carry a maximum weight given by maximumWeight
func Knapsack(maximumWeight int, items []Item) int {
	maxValue := 0
	for _, v := range getCombos(items) {
		if v.weight <= maximumWeight && v.value > maxValue {
			maxValue = v.value
		}
	}

	return maxValue
}

func knapsackOf(items map[int]Item) knapsack {
	knapsack := knapsack{}
	for _, v := range items {
		knapsack.weight += v.Weight
		knapsack.value += v.Value
	}
	return knapsack
}

func getCombos(items []Item) (out []knapsack) {
	if len(items) == 0 {
		return out
	}

	for i, v := range items {
		out = append(
			out,
			knapsack{weight: v.Weight, value: v.Value, indices: map[int]bool{i: true}},
		)
	}

	currentKnapsacks := make([]knapsack, len(out))
	copy(currentKnapsacks, out)

	for i := 0; i < len(items)-1; i++ {
		newOnes := []knapsack{}

		for i, v := range currentKnapsacks {
			for j := i + 1; j < len(items); j++ {
				item := items[j]
				m := v.indices
				if !mapContains(m, j) {
					newM := copyMap(m)
					newM[j] = true
					newOnes = append(
						newOnes,
						knapsack{
							weight:  v.weight + item.Weight,
							value:   v.value + item.Value,
							indices: newM,
						},
					)
				}
			}
		}

		out = append(out, newOnes...)
		currentKnapsacks = newOnes
	}
	return out
}

func copyMap[V any](m map[int]V) map[int]V {
	out := map[int]V{}
	for k, v := range m {
		out[k] = v
	}
	return out
}

func mapContains[V any](m map[int]V, k int) bool {
	_, exists := m[k]
	return exists
}

func sliceContains(s []map[int]Item, m map[int]Item) bool {
	for _, v := range s {
		if reflect.DeepEqual(v, m) {
			return true
		}
	}
	return false
}

/*

a
b
c
d
e
a b
a c
a d
a e
b c
b d
b e
c d
c e
d e
a b c
a b d
a b e
a c d
a c e
a d e
b c d
b c e
b d e
c d e
a b c d
a b c e
a b d e
a c d e
b c d e
a b c d e


a
b
c
d
a b
a c
a d
b c
b d
c d
a b c
a b d
a c d
b c d
a b c d

a
b
c
a b
a c
b c
a b c


a
b
a b

*/
