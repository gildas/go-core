package core

import "fmt"

// Join joins a slice of items into a string using a separator.
func Join[T fmt.Stringer](items []T, separator string) string {
	if len(items) == 0 {
		return ""
	}
	if len(items) == 1 {
		return items[0].String()
	}
	var buf []byte
	for i, item := range items {
		if i > 0 {
			buf = append(buf, separator...)
		}
		buf = append(buf, item.String()...)
	}
	return string(buf)
}

// JoinWithFunc joins a slice of items into a string using a separator.
//
// The function is called for each item to get its string representation.
func JoinWithFunc[T any](items []T, separator string, stringer func(item T) string) string {
	if len(items) == 0 {
		return ""
	}
	if len(items) == 1 {
		return stringer(items[0])
	}
	var buf []byte
	for i, item := range items {
		if i > 0 {
			buf = append(buf, separator...)
		}
		buf = append(buf, stringer(item)...)
	}
	return string(buf)
}

// MapJoin joins two or more maps of the same kind
//
// None of the maps are modified, a new map is created.
//
// If two maps have the same key, the latter map overwrites the value from the former map.
func MapJoin[K comparable, T any](maps ...map[K]T) map[K]T {
	results := map[K]T{}

	for _, m := range maps {
		for key, value := range m {
			results[key] = value
		}
	}
	return results
}
