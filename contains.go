package core

// Contains checks if a slice contains a value.
func Contains[T comparable](slice []T, value T) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

// ContainsWithFunc checks if a slice contains a value using a custom function.
func ContainsWithFunc[T any](slice []T, value T, f func(T, T) bool) bool {
	for _, item := range slice {
		if f(item, value) {
			return true
		}
	}
	return false
}
