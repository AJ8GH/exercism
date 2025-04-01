package knapsack

import "reflect"

type Item struct {
	Weight, Value int
}

type knapsack struct {
	weight, value, numItems int
	items                   []Item
}

// Knapsack takes in a maximum carrying capacity and a collection of items
// and returns the maximum value that can be carried by the knapsack
// given that the knapsack can only carry a maximum weight given by maximumWeight
func Knapsack(maximumWeight int, items []Item) int {
	knapsacks := []knapsack{}
	for _, v := range getCombos(items) {
		knapsacks = append(knapsacks, knapsackOf(v))
	}

	maxValue := 0
	for _, v := range knapsacks {
		if v.weight <= maximumWeight && v.value > maxValue {
			maxValue = v.value
		}
	}

	return maxValue
}

func knapsackOf(items map[int]Item) knapsack {
	knapsack := knapsack{}
	for _, v := range items {
		knapsack.numItems++
		knapsack.items = append(knapsack.items, v)
		knapsack.weight += v.Weight
		knapsack.value += v.Value
	}
	return knapsack
}

func getCombos(items []Item) (out []map[int]Item) {
	if len(items) == 0 {
		return out
	}

	for i, v := range items {
		out = append(out, map[int]Item{i: v})
	}

	for i := 0; i < len(items)-1; i++ {
		for j, item := range items {
			for _, m := range out {
				if !mapContains(m, j) {
					mCopy := copyMap(m)
					mCopy[j] = item
					if !sliceContains(out, mCopy) {
						out = append(out, mCopy)
					}
				}
			}
		}
	}

	return out
}

func copyMap(m map[int]Item) map[int]Item {
	out := map[int]Item{}
	for k, v := range m {
		out[k] = v
	}
	return out
}

func mapContains(m map[int]Item, k int) bool {
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
