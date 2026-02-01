package core

import "fmt"

// Join joins a slice of items into a string using a separator.
func Join[S ~[]T, T fmt.Stringer](items S, separator string) string {
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
func JoinWithFunc[S ~[]T, T any](items S, separator string, stringer func(item T) string) string {
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
