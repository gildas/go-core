package core

// TypeCarrier represents object that carries a Type
type TypeCarrier interface {
	// GetType tells the type of this object
	GetType() string
}
