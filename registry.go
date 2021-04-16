package core

import (
	"encoding/json"
	"errors"
	"fmt"
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

// UnmarshalJSON unmarshal a payload into a Type Carrier
//
// The interface that is returned contains a pointer to the TypeCarrier structure
func (registry TypeRegistry) UnmarshalJSON(payload []byte) (interface{}, error) {
	header := struct{Type string `json:"type"`}{}
	if err := json.Unmarshal(payload, &header); err != nil {
		return nil, err
	}
	if len(header.Type) == 0 {
		return nil, errors.New(`Missing JSON Property "type"`)
	}

	if valueType, found := registry[header.Type]; found {
		value := reflect.New(valueType).Interface()
		if err := json.Unmarshal(payload, value); err != nil {
			return nil, err
		}
		return value, nil
	}
	return nil, fmt.Errorf(`Unsupported Type "%s"`, header.Type)
}
