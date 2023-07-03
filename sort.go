package core

// Sort sorts a slice of items using the Quick Sort algorithm
//
// Note: The items slice is modified in place.
//
// Example:
//
//	// Sort a slice of numbers
//	Sort(numbers, func(a, b int) bool {
//		return a < b
//	})
func Sort[T any](items []T, sorter func(T, T) bool) {
	index := partition(items, len(items)/2, sorter)
	if index > 1 {
		Sort(items[:index], sorter)
	}
	if index+1 < len(items) {
		Sort(items[index+1:], sorter)
	}
}

func partition[T any](items []T, pivot int, sorter func(T, T) bool) int {
	pivotValue := items[pivot]
	items[pivot], items[len(items)-1] = items[len(items)-1], items[pivot]
	storeIndex := 0
	for i := 0; i < len(items)-1; i++ {
		if sorter(items[i], pivotValue) {
			items[i], items[storeIndex] = items[storeIndex], items[i]
			storeIndex++
		}
	}
	items[storeIndex], items[len(items)-1] = items[len(items)-1], items[storeIndex]
	return storeIndex
}
