package core_test

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type TestError struct {
	ID string
	Value interface{}
}

func (err TestError) Error() string {
	return err.ID
}

func TestCanExtractFieldsFromError(t *testing.T) {
	err := TestError{ID: "id_this", Value: nil}

	var field reflect.Value
	errValue := reflect.ValueOf(err)
	require.NotNil(t, errValue)
	t.Logf("ErrValue: %#v", errValue)
	
	field = errValue.FieldByName("ID")
	require.True(t, field.IsValid(), "Field ID is not valid")
	require.True(t, field.Type().Kind().String() == "string", "Field ID is not a string, it is a %s", field.Type().Kind().String())
	assert.Equal(t, "id_this", field.String())

	field = errValue.FieldByName("Value")
	require.True(t, field.IsValid(), "Field Value is not valid")
	require.True(t, field.IsNil(), "Field Value has a value")
}