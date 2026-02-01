package core

// Map maps a slice of items into a new slice
//
// Note: The result is a new slice, the original is not modified
//
// Example:
//
//		// Map all numbers in a slice to their square
//	 numbers := []int{1, 2, 3, 4, 5}
//		squares := Map(numbers, func(number int) int64 {
//			return int64(number * number)
//		})
func Map[S ~[]T, T any, R any](items S, mapper func(T) R) (result []R) {
	for _, item := range items {
		result = append(result, mapper(item))
	}
	return result
}
