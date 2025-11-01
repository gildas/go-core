package core_test

import (
	"encoding/json"
	"testing"

	"github.com/gildas/go-core"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestIdentifiableReference(t *testing.T) {
	id := uuid.New()
	stuff := &Stuff{ID: id}
	reference := core.GetReference(stuff)
	expected := `{"id":"` + id.String() + `"}`

	payload, err := json.Marshal(reference)
	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(payload))
}

func TestStringIdentifiableReference(t *testing.T) {
	stuff := &Something4{ID: "string-id-123"}
	reference := core.GetReference(stuff)
	expected := `{"id":"string-id-123"}`

	payload, err := json.Marshal(reference)
	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(payload))
}

func TestStringerReference(t *testing.T) {
	something := Something2{Data: "stringer-id-456"}
	reference := core.GetReference(something)
	expected := `{"id":"stringer-id-456"}`

	payload, err := json.Marshal(reference)
	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(payload))
}

func TestUnknownReference(t *testing.T) {
	something := struct {
		Value string `json:"value"`
	}{
		Value: "no-id",
	}
	reference := core.GetReference(something)
	expected := `{"value":"no-id"}`

	payload, err := json.Marshal(reference)
	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(payload))
}
