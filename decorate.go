package core

import (
	"encoding/json"
	"net/url"
	"path"
	"reflect"
	"strings"
)

// DecoratedResource is a resource that contains a Self link
type DecoratedResource struct {
	Data    any
	SelfURI string `json:"selfURI,omitempty"`
}

// Decorate decorates a struct that can be identified
func Decorate(item any, rootpath string) *DecoratedResource {
	decorated := &DecoratedResource{Data: item}
	resourceName := pluralize(strings.ToLower(reflect.TypeOf(item).Name()))
	switch identifier := item.(type) {
	case Identifiable:
		decorated.SelfURI = path.Join(rootpath, resourceName, identifier.GetID().String())
	case StringIdentifiable:
		decorated.SelfURI = path.Join(rootpath, resourceName, identifier.GetID())
	}
	return decorated
}

// DecorateWithURL decorates a struct that can be identified with a URL
func DecorateWithURL(item any, root url.URL) *DecoratedResource {
	decorated := &DecoratedResource{Data: item}
	resourceName := pluralize(strings.ToLower(reflect.TypeOf(item).Name()))
	switch identifier := item.(type) {
	case Identifiable:
		decorated.SelfURI = root.JoinPath(resourceName, identifier.GetID().String()).String()
	case StringIdentifiable:
		decorated.SelfURI = root.JoinPath(resourceName, identifier.GetID()).String()
	}
	return decorated
}

// DecorateAll decorates all items in a slice of identifiable items
func DecorateAll[S ~[]T, T any](items S, rootpath string) []DecoratedResource {
	decorated := make([]DecoratedResource, len(items))
	for i, item := range items {
		decorated[i] = *Decorate(item, rootpath)
	}
	return decorated
}

// ResourceName returns the name of a resource
//
// If the data is a core.TypeCarrier, it will return the pluralized name of the type
// Otherwise, it will return the pluralized name of the type of the data, via reflection
func (resource DecoratedResource) ResourceName() string {
	if carrier, ok := resource.Data.(TypeCarrier); ok {
		return pluralize(strings.ToLower(carrier.GetType()))
	}
	dataType := reflect.TypeOf(resource.Data)
	if dataType.Kind() == reflect.Ptr {
		return pluralize(strings.ToLower(dataType.Elem().Name()))
	}
	return pluralize(strings.ToLower(dataType.Name()))
}

// MarshalJSON marshals the DecoratedResource to JSON
//
// implements the json.Marshaler interface
func (resource DecoratedResource) MarshalJSON() ([]byte, error) {
	payload, err := json.Marshal(resource.Data)
	if err != nil {
		return nil, err
	}
	decoration, _ := json.Marshal(struct {
		SelfURI string `json:"selfURI"`
	}{
		SelfURI: resource.SelfURI,
	})
	return []byte(strings.TrimRight(string(payload), "}") + ", " + string(decoration)[1:]), nil
}

// pluralize returns the plural form of a given string
//
// this is a very simple implementation and may not work for all cases
// It does not handle irregular plurals, such as "child" -> "children"
func pluralize(name string) string {
	if strings.HasSuffix(name, "ch") || strings.HasSuffix(name, "sh") {
		return name + "es"
	}
	if strings.HasSuffix(name, "o") || strings.HasSuffix(name, "s") || strings.HasSuffix(name, "x") {
		return name + "es"
	}
	if strings.HasSuffix(name, "y") {
		return name[:len(name)-1] + "ies"
	}
	if strings.HasSuffix(name, "f") {
		return name[:len(name)-1] + "ves"
	}
	if strings.HasSuffix(name, "z") {
		return name + "zes"
	}

	return name + "s"
}
