package core

import "golang.org/x/exp/slices"

// Sort sorts a slice of items
//
// Note: The items slice is modified in place.
//
// calls slices.SortFunc from the golang.org/x/exp/slices package.
//
// Example:
//
//	// Sort a slice of numbers
//	Sort(numbers, func(a, b int) bool {
//		return a < b
//	})
func Sort[S ~[]T, T any](items S, sorter func(T, T) bool) {
	slices.SortFunc(items, func(a, b T) int {
		if sorter(a, b) {
			return -1
		} else if sorter(b, a) {
			return 1
		}
		return 0
	})

	/*
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
			for i := range len(items) - 1 {
				if sorter(items[i], pivotValue) {
					items[i], items[storeIndex] = items[storeIndex], items[i]
					storeIndex++
				}
			}
			items[storeIndex], items[len(items)-1] = items[len(items)-1], items[storeIndex]
			return storeIndex
	*/
}
