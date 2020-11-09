package core

import (
	"reflect"
)

// TypeRegistry contains a map of identifier vs Type
type TypeRegistry map[string]reflect.Type

// Add adds one or more TypeCarriers to the TypeRegistry
func (registry TypeRegistry) Add(classes ...TypeCarrier) TypeRegistry {
	for _, class := range classes {
		registry[class.GetType()] = reflect.TypeOf(class)
	}
	return registry
}
