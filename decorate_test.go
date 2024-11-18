package core_test

import (
	"encoding/json"
	"errors"
	"net/url"
	"testing"

	"github.com/gildas/go-core"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// Mock Identifiable and StringIdentifiable types for testing
type MockIdentifiable struct {
	ID uuid.UUID
}

// GetType returns the type of the MockIdentifiable
//
// Implements the TypeCarrier interface
func (m MockIdentifiable) GetType() string {
	return "MockIdentifiable"
}

// GetID returns the ID of the MockIdentifiable
//
// Implements the Identifiable interface
func (m MockIdentifiable) GetID() uuid.UUID {
	return m.ID
}

type MockStringIdentifiable struct {
	ID string
}

// GetID returns the ID of the MockStringIdentifiable
//
// Implements the StringIdentifiable interface
func (m MockStringIdentifiable) GetID() string {
	return m.ID
}

// BogusValue is a bogus value that fails to marshal
type BogusValue struct {
}

func (v BogusValue) MarshalJSON() ([]byte, error) {
	return nil, errors.ErrUnsupported
}

func TestDecorate(t *testing.T) {
	rootpath := "/api/v1"

	t.Run("Identifiable", func(t *testing.T) {
		item := MockIdentifiable{ID: uuid.New()}
		decorated := core.Decorate(item, rootpath)
		expectedURI := rootpath + "/mockidentifiables/" + item.ID.String()
		assert.Equal(t, expectedURI, decorated.SelfURI)
		assert.Equal(t, item, decorated.Data)
	})

	t.Run("StringIdentifiable", func(t *testing.T) {
		item := MockStringIdentifiable{ID: "abc"}
		decorated := core.Decorate(item, rootpath+"/")
		expectedURI := rootpath + "/mockstringidentifiables/abc"
		assert.Equal(t, expectedURI, decorated.SelfURI)
		assert.Equal(t, item, decorated.Data)
	})
}

func TestDecorateWithURL(t *testing.T) {
	root := core.Must(url.Parse("http://localhost:8080/api/v1"))

	t.Run("Identifiable", func(t *testing.T) {
		item := MockIdentifiable{ID: uuid.New()}
		decorated := core.DecorateWithURL(item, *root)
		expectedURI := core.Must(url.Parse("http://localhost:8080/api/v1/mockidentifiables/" + item.ID.String()))
		assert.Equal(t, expectedURI.String(), decorated.SelfURI)
		assert.Equal(t, item, decorated.Data)
	})

	t.Run("StringIdentifiable", func(t *testing.T) {
		item := MockStringIdentifiable{ID: "abc"}
		decorated := core.DecorateWithURL(item, *root)
		expectedURI := core.Must(url.Parse("http://localhost:8080/api/v1/mockstringidentifiables/" + item.ID))
		assert.Equal(t, expectedURI.String(), decorated.SelfURI)
		assert.Equal(t, item, decorated.Data)
	})
}

func TestDecorateAll(t *testing.T) {
	rootpath := "/api/v1"
	items := []MockStringIdentifiable{
		{ID: "123"},
		{ID: "456"},
	}
	decorated := core.DecorateAll(items, rootpath)
	assert.Len(t, decorated, 2)
	assert.Equal(t, rootpath+"/mockstringidentifiables/123", decorated[0].SelfURI)
	assert.Equal(t, rootpath+"/mockstringidentifiables/456", decorated[1].SelfURI)
}

func TestResourceName(t *testing.T) {
	t.Run("TypeCarrier", func(t *testing.T) {
		item := MockIdentifiable{}
		decorated := core.Decorate(item, "")
		assert.Equal(t, "mockidentifiables", decorated.ResourceName())
	})

	t.Run("NonTypeCarrier", func(t *testing.T) {
		item := MockStringIdentifiable{ID: "123"}
		decorated := core.Decorate(item, "")
		assert.Equal(t, "mockstringidentifiables", decorated.ResourceName())
	})

	t.Run("NonTypeCarrierPointer", func(t *testing.T) {
		item := &MockStringIdentifiable{ID: "123"}
		decorated := core.Decorate(item, "")
		assert.Equal(t, "mockstringidentifiables", decorated.ResourceName())
	})
}

func TestMarshalJSON(t *testing.T) {
	item := MockStringIdentifiable{ID: "abc"}
	decorated := core.Decorate(item, "/var/api")
	expectedJSON := `{"selfURI":"/var/api/mockstringidentifiables/abc","ID":"abc"}`

	data, err := json.Marshal(decorated)
	assert.NoError(t, err)
	assert.JSONEq(t, expectedJSON, string(data))

	_, err = json.Marshal(core.Decorate(BogusValue{}, "/var/api"))
	assert.Error(t, err)
}
