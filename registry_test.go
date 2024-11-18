package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/gildas/go-core"
)

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

	value, ok := object.(Something)
	require.True(t, ok, "Return object should carry a *Something1")
	require.NotNil(t, value, "Actual value cannot be nil")
	assert.Equal(t, "Hello", value.GetData())
}

func TestCanUnmarshalTypeCarrierWithTypetag(t *testing.T) {
	registry := TypeRegistry{}.Add(Something1{}, Something2{})
	require.Equal(t, 2, len(registry))

	payload := []byte(`{"__type": "something1", "data": "Hello"}`)
	object, err := registry.UnmarshalJSON(payload, "__type")
	require.Nilf(t, err, "Failed to Unmarshal payload: %s", err)
	require.NotNil(t, object, "Returned object cannot be nil")

	value, ok := object.(Something)
	require.True(t, ok, "Return object should carry a *Something1")
	require.NotNil(t, value, "Actual value cannot be nil")
	assert.Equal(t, "Hello", value.GetData())
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

	payload := []byte(`{"type": 2", "data": "Hello"}`)
	_, err := registry.UnmarshalJSON(payload)
	require.NotNil(t, err)
	assert.Equal(t, "invalid character '\"' after object key:value pair", err.Error())
}

func TestShouldFailUnmarshalingTypeCarrierWithInvalidJSON2(t *testing.T) {
	registry := TypeRegistry{}.Add(Something1{}, Something2{})
	require.Equal(t, 2, len(registry))

	payload := []byte(`{"type": "something1", "data": 2}`)
	_, err := registry.UnmarshalJSON(payload)
	require.NotNil(t, err)
	assert.Equal(t, "json: cannot unmarshal number into Go struct field Something1.data of type string", err.Error())
}
