package core

// Named describes types that can get their Name
type Named interface {
	GetName() string
}