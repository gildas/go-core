package core

// Map maps a slice of items into a new slice
//
// Note: The result is a new slice, the original is not modified
//
// Example:
// 	// Map all numbers in a slice to their square
// 	squares := Map(numbers, func(number int) int {
// 		return number * number
// 	})
func Map[T any, R any](items []T, mapper func(T) R) (result []R) {
	for _, item := range items {
		result = append(result, mapper(item))
	}
	return result
}
