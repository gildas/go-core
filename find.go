package core

// Find finds a value in a slice
func Find[T comparable](items []T, value T) (t T, found bool) {
	for _, item := range items {
		if item == value {
			return item, true
		}
	}
	return t, false
}

// FindWithFunc is a function that finds a value in a slice
func FindWithFunc[T any](items []T, predicate func(T) bool) (t T, found bool) {
	for _, item := range items {
		if predicate(item) {
			return item, true
		}
	}
	return t, false
}
