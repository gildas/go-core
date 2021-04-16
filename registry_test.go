package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/gildas/go-core"
)

type Something interface {
	TypeCarrier
}

type Something1 struct {
	Data string `json:"data"`
}

type Something2 struct {
	Data string `json:"data"`
}

func (s Something1) GetType() string {
	return "something1"
}

func (s Something2) GetType() string {
	return "something2"
}

func TestCanCreateTypeRegistry(t *testing.T) {
	registry := TypeRegistry{}.Add(Something1{}, Something2{})
	assert.Equal(t, 2, len(registry))
}

func TestCanUnmarshalTypeCarrier(t *testing.T) {
	registry := TypeRegistry{}.Add(Something1{}, Something2{})
	require.Equal(t, 2, len(registry))

	payload := []byte(`{"type": "something1", "data": "Hello"}`)
	object, err := registry.UnmarshalJSON(payload)
	require.Nilf(t, err, "Failed to Unmarshal payload: %s", err)
	require.NotNil(t, object, "Returned object cannot be nil")

	value, ok := object.(*Something1)
	require.True(t, ok, "Return object should carry a *Something1")
	require.NotNil(t, value, "Actual value cannot be nil")
	assert.Equal(t, "Hello", value.Data)
}

func TestShouldFailUnmarshalingTypeCarrierWithoutType(t *testing.T) {
	registry := TypeRegistry{}.Add(Something1{}, Something2{})
	require.Equal(t, 2, len(registry))

	payload := []byte(`{"data": "Hello"}`)
	_, err := registry.UnmarshalJSON(payload)
	require.NotNil(t, err)
	assert.Equal(t, `Missing JSON Property "type"`, err.Error())
}

func TestShouldFailUnmarshalingTypeCarrierWithInvalidType(t *testing.T) {
	registry := TypeRegistry{}.Add(Something1{}, Something2{})
	require.Equal(t, 2, len(registry))

	payload := []byte(`{"type": "something3", "data": "Hello"}`)
	_, err := registry.UnmarshalJSON(payload)
	require.NotNil(t, err)
	assert.Equal(t, `Unsupported Type "something3"`, err.Error())
}

func TestShouldFailUnmarshalingTypeCarrierWithInvalidJSON(t *testing.T) {
	registry := TypeRegistry{}.Add(Something1{}, Something2{})
	require.Equal(t, 2, len(registry))

	payload := []byte(`{"type": 2, "data": "Hello"}`)
	_, err := registry.UnmarshalJSON(payload)
	require.NotNil(t, err)
	assert.Equal(t, "json: cannot unmarshal number into Go struct field .type of type string", err.Error())

	payload = []byte(`{"type": "something1", "data": 2}`)
	_, err = registry.UnmarshalJSON(payload)
	require.NotNil(t, err)
	assert.Equal(t, "json: cannot unmarshal number into Go struct field Something1.data of type string", err.Error())
}
