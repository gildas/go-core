package core

import "github.com/google/uuid"

// Identifiable describes that can get their Identifier as a UUID
type Identifiable interface {
	GetID() uuid.UUID
}

// StringIdentifiable describes that can get their Identifier as a string
type StringIdentifiable interface {
	GetID() string
}

// EqualIdentifiable compares two values for equality by their Identifier
func EqualIdentifiable[T Identifiable](a, b T) bool {
	return a.GetID() == b.GetID()
}

// MatchIdentifiable returns a function that matches an Identifiable value
func MatchIdentifiable[T Identifiable](search T) func(T) bool {
	return func(item T) bool {
		return item.GetID() == search.GetID()
	}
}

// EqualStringIdentifiable compares two values for equality by their Identifier
func EqualStringIdentifiable[T StringIdentifiable](a, b T) bool {
	return a.GetID() == b.GetID()
}

// MatchStringIdentifiable returns a function that matches a StringIdentifiable value
func MatchStringIdentifiable[T StringIdentifiable](search T) func(T) bool {
	return func(item T) bool {
		return item.GetID() == search.GetID()
	}
}
