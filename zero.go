package core

// IsZeroer is an interface that can be implemented by types that can tell if their value is zero
type IsZeroer interface {
	IsZero() bool
}
