package core

// Named describes types that can get their Name
type Named interface {
	GetName() string
}

// RqualNamed compares two values for equality by name
func EqualNamed[T Named](a, b T) bool {
	return a.GetName() == b.GetName()
}

// MatchNamed returns a function that matches a Named value
func MatchNamed[T Named](search T) func(T) bool {
	return func(item T) bool {
		return item.GetName() == search.GetName()
	}
}
