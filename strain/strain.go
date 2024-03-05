package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](s []T, predicate func(T) bool) []T {

	col := []T{}
	for _, val := range s {
		if predicate(val) {
			col = append(col, val)
		}
	}

	return col
}

func Discard[T any](s []T, predicate func(T) bool) []T {

	col := []T{}
	for _, val := range s {
		if !predicate(val) {
			col = append(col, val)
		}
	}

	return col
}
