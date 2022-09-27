package core

// Reduce reduces a slice of items into a single value
//
// Example:
// 	// Sum all numbers in a slice
// 	sum := Reduce(numbers, 0, func(sum, number int) int {
// 		return sum + number
// 	})
func Reduce[T any, R any](items []T, initial R, reducer func(R, T) R) (result R) {
	result = initial
	for _, item := range items {
		result = reducer(result, item)
	}
	return result
}
