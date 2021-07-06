package core_test

import (
	"fmt"

	"github.com/gildas/go-core"
)

type DataHolder interface {
	core.TypeCarrier
	GetData() string
}

type DataSpec1 struct {
	Data string `json:"data"`
}

type DataSpec2 struct {
	Data string `json:"data"`
}

func (s DataSpec1) GetType() string {
	return "dataspec1"
}

func (s DataSpec1) GetData() string {
	return s.Data
}

func (s DataSpec2) GetType() string {
	return "dataspec2"
}

func (s DataSpec2) GetData() string {
	return s.Data
}

func ExampleTypeRegistry_UnmarshalJSON() {
	// Typically, each struct would be declared in its own go file
	// and the core.TypeRegistry.Add() func would be done in the init() func of each file
	registry := core.TypeRegistry{}.Add(DataSpec1{}, DataSpec2{})
	payload := []byte(`{"type": "dataspec1", "data": "Hello"}`)

	object, err := registry.UnmarshalJSON(payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	value, ok := object.(DataHolder)
	if !ok {
		fmt.Println("Not a DataHolder")
		return
	}
	fmt.Println(value.GetData())

	// Here we will read the Type of the payload through a custom JSON property
	payload = []byte(`{"__type": "dataspec1", "data": "Hello"}`)

	object, err = registry.UnmarshalJSON(payload, "__type")
	if err != nil {
		fmt.Println(err)
		return
	}

	value, ok = object.(DataHolder)
	if !ok {
		fmt.Println("Not a DataHolder")
		return
	}
	fmt.Println(value.GetData())
	// Output:
	// Hello
	// Hello
}
