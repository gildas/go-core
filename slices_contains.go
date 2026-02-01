package core

import (
	slices "slices"
)

// Contains checks if a slice contains a value.
func Contains[S ~[]T, T comparable](slice S, value T) bool {
	return slices.Contains(S(slice), T(value))
}

// ContainsWithFunc checks if a slice contains a value using a custom function.
//
// The function f should return true if the slice item matches the value.
//
// If T implements [Identifiable] or [StringIdentifiable] or [Named],
// consider using [MatchIdentifiable], [MatchStringIdentifiable] or [MatchNamed]
// as the matching/equality function.
func ContainsWithFunc[S ~[]T, T any](slice S, value T, predicate func(item T, value T) bool) bool {
	for _, item := range slice {
		if predicate(item, value) {
			return true
		}
	}
	return false
}
