package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
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

// UnmarshalJSON unmarshal a payload into a Type Carrier
//
// The interface that is returned contains a pointer to the TypeCarrier structure.
//
// The default typetag is "type", but you can replace it by one or more of your own.
//
// Examples:
//   object, err := registry.UnmarshalJSON(payload)
//   object, err := registry.UnmarshalJSON(payload, "__type", "Type")
func (registry TypeRegistry) UnmarshalJSON(payload []byte, typetag... string) (interface{}, error) {
	if len(typetag) == 0 {
		typetag = []string{"type"}
	}
	guts := map[string]json.RawMessage{}
	if err := json.Unmarshal(payload, &guts); err != nil {
		return nil, err
	}
	objectType := ""
	for _, tag := range typetag {
		if value, found := guts[tag]; found {
			objectType = strings.Trim(string(value), "\"")
		}
	}
	if len(objectType) == 0 {
		return nil, errors.New(`Missing JSON Property "type"`)
	}

	if valueType, found := registry[objectType]; found {
		value := reflect.New(valueType).Interface()
		if err := json.Unmarshal(payload, value); err != nil {
			return nil, err
		}
		return value, nil
	}
	return nil, fmt.Errorf(`Unsupported Type "%s"`, objectType)
}
