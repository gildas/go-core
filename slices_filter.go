package core

// Filter filters a slice of items based on a filter function
//
// Note: The result is a new slice, the original is not modified
//
// Example:
//
//	// Filter all positive numbers in a slice
//	numbers := Filter(numbers, func(number int) bool {
//		return number > 0
//	})
func Filter[S ~[]T, T any](items S, filter func(T) bool) (result []T) {
	for _, item := range items {
		if filter(item) {
			result = append(result, item)
		}
	}
	return result
}
