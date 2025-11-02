package core

import (
	"fmt"
)

// GetReference gets a reference of an identifiable
//
// If the element implements Identifiable, the reference will use the UUID ID.
//
// If the element implements StringIdentifiable, the reference will use the string ID.
//
// If the element implements fmt.Stringer, the reference will use the string representation.
//
// Otherwise, the element itself is returned.
func GetReference(element any) any {
	var ref struct {
		ID string `json:"id"`
	}

	switch actual := element.(type) {
	case Identifiable:
		ref.ID = actual.GetID().String()
	case StringIdentifiable:
		ref.ID = actual.GetID()
	case fmt.Stringer:
		ref.ID = actual.String()
	default:
		return element
	}
	return ref
}
