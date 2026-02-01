package core

// Find finds a value in a slice
func Find[S ~[]T, T comparable](items S, value T) (t T, found bool) {
	for _, item := range items {
		if item == value {
			return item, true
		}
	}
	return t, false
}

// FindWithFunc is a function that finds a value in a slice
//
// Note that the returned value is a copy of the item in the slice.
func FindWithFunc[S ~[]T, T any](items S, predicate func(T) bool) (t T, found bool) {
	for _, item := range items {
		if predicate(item) {
			return item, true
		}
	}
	return t, false
}
