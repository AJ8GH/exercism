package strain

func Keep[T any](in []T, predicate func(t T) bool) []T {
	return filter(in, predicate, true)
}

func Discard[T any](in []T, predicate func(t T) bool) []T {
	return filter(in, predicate, false)
}

func filter[T any](in []T, predicate func(t T) bool, predicateResult bool) []T {
	out := []T{}
	for _, v := range in {
		if predicate(v) == predicateResult {
			out = append(out, v)
		}
	}
	return out
}

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
