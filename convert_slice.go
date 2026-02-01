package core

import "encoding/json"

// ConvertSliceFromAny converts a slice of any interface to a slice of specific type T.
//
// If an item cannot be converted to type T, it is skipped.
//
// If an item is not of type T, it attempts to marshal and unmarshal it as JSON.
// This allows to convert a JSON payload with an array of objects into a slice of T.
// Note that this requires that the item can be marshaled to JSON and that T can be unmarshaled
// from JSON and is slower than a direct type assertion.
//
// Example:
//
//	// Convert a slice of any to a slice of strings
//	anySlice := []any{"one", "two", "three"}
//	stringSlice := ConvertSliceFromAny[string](anySlice)
func ConvertSliceFromAny[T any](items []any) []T {
	slice := make([]T, 0, len(items))
	for _, item := range items {
		if v, ok := item.(T); ok {
			slice = append(slice, v)
		} else {
			if payload, err := json.Marshal(item); err == nil {
				if err := json.Unmarshal(payload, &v); err == nil {
					slice = append(slice, v)
				}
			}
		}
	}
	return slice
}

// ConvertSliceToAny converts a slice of specific type T to a slice of any interface.
//
// Example:
//
//	// Convert a slice of strings to a slice of any
//	stringSlice := []string{"one", "two", "three"}
//	anySlice := ConvertSliceToAny[string](stringSlice)
func ConvertSliceToAny[T any](items []T) []any {
	slice := make([]any, 0, len(items))
	for _, item := range items {
		slice = append(slice, item)
	}
	return slice
}
