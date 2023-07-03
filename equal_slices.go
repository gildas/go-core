package core

// EqualSlices checks if two slices are equal.
//
// Two slices are equal if they have the same length and
// all elements of the first slice exist in the second slice. And vice versa.
func EqualSlices[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for _, item := range a {
		if !Contains(b, item) {
			return false
		}
	}
	for _, item := range b {
		if !Contains(a, item) {
			return false
		}
	}
	return true
}

// EqualSlicesWithFunc checks if two slices are equal.
//
// Two slices are equal if they have the same length and
// all elements of the first slice exist in the second slice. And vice versa.
func EqualSlicesWithFunc[T any](a, b []T, compare func(T, T) bool) bool {
	if len(a) != len(b) {
		return false
	}
	for _, item := range a {
		if !ContainsWithFunc(b, item, compare) {
			return false
		}
	}
	for _, item := range b {
		if !ContainsWithFunc(a, item, compare) {
			return false
		}
	}
	return true
}
