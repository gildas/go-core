package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"golang.org/x/exp/slices"
)

// CaseInsensitiveTypeRegistry contains a map of identifier vs Type
// The Registry is case insensitive, so "something" and "Something" are the same
type CaseInsensitiveTypeRegistry map[string]reflect.Type

// Add adds one or more TypeCarriers to the CaseInsensitiveTypeRegistry
func (registry CaseInsensitiveTypeRegistry) Add(classes ...TypeCarrier) CaseInsensitiveTypeRegistry {
	for _, class := range classes {
		registry[strings.ToLower(class.GetType())] = reflect.TypeOf(class)
	}
	return registry
}

// SupportedTypes returns a list of supported types in the registry
func (registry CaseInsensitiveTypeRegistry) SupportedTypes() []string {
	supportedTypes := make([]string, 0, len(registry))
	for key := range registry {
		supportedTypes = append(supportedTypes, key)
	}
	slices.Sort(supportedTypes)
	return supportedTypes
}

// UnmarshalJSON unmarshal a payload into a Type Carrier
//
// The interface that is returned contains a pointer to the TypeCarrier structure.
//
// The default typetag is "type", but you can replace it by one or more of your own.
//
// Examples:
//
//	object, err := registry.UnmarshalJSON(payload)
//	object, err := registry.UnmarshalJSON(payload, "__type", "Type")
func (registry CaseInsensitiveTypeRegistry) UnmarshalJSON(payload []byte, typetag ...string) (interface{}, error) {
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

	if valueType, found := registry[strings.ToLower(objectType)]; found {
		value := reflect.New(valueType).Interface()
		if err := json.Unmarshal(payload, value); err != nil {
			return nil, err
		}
		return value, nil
	}
	return nil, fmt.Errorf(`Unsupported Type "%s"`, objectType)
}
