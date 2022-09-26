package core_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"html/template"

	"github.com/gildas/go-core"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type StructuredError struct {
	ID      string
	Message string
	What    string
	Value   string
}

func (err StructuredError) Error() string {
	return err.Message
}

type StructuredErrorWithIdentifiable struct {
	ID      string
	Message string
	What    string
	Value   *Stuff
}

func (err StructuredErrorWithIdentifiable) Error() string {
	return err.Message
}

type StructuredErrorWithUUID struct {
	ID      string
	Message string
	What    string
	Value   uuid.UUID
}

func (err StructuredErrorWithUUID) Error() string {
	return err.Message
}

type Stuff struct {
	ID uuid.UUID
}

func (stuff Stuff) GetID() uuid.UUID {
	return stuff.ID
}

func TestHTTPResponderWithJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		core.RespondWithJSON(res, http.StatusOK, map[string]string{"hello": "world"})
	}))
	defer server.Close()

	res, err := http.Get(server.URL)
	require.NoErrorf(t, err, "Failed to get %s", server.URL)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	var payload map[string]string
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	require.NoErrorf(t, err, "Failed to read body of %s", server.URL)
	err = json.Unmarshal(body, &payload)
	require.NoError(t, err, "Failed to unmarshal JSON")
	assert.Equal(t, map[string]string{"hello": "world"}, payload)
}

func TestHTTPResponderWithError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		core.RespondWithError(res, http.StatusNotFound, fmt.Errorf("Not found"))
	}))
	defer server.Close()

	res, err := http.Get(server.URL)
	require.NoErrorf(t, err, "Failed to get %s", server.URL)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	var payload map[string]string
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	require.NoErrorf(t, err, "Failed to read body of %s", server.URL)
	err = json.Unmarshal(body, &payload)
	require.NoError(t, err, "Failed to unmarshal JSON")
	assert.Equal(t, map[string]string{"error": "Not found", "http_status": "404"}, payload)
}

func TestHTTPResponderWithStructuredError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		err := StructuredError{ID: "1234", Message: "Not found", What: "something", Value: "53c1e09c-bc7e-4989-8bc9-d71acdb0a013"}
		core.RespondWithError(res, http.StatusNotFound, err)
	}))
	defer server.Close()

	res, err := http.Get(server.URL)
	require.NoErrorf(t, err, "Failed to get %s", server.URL)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	var payload map[string]string
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	require.NoErrorf(t, err, "Failed to read body of %s", server.URL)
	err = json.Unmarshal(body, &payload)
	require.NoError(t, err, "Failed to unmarshal JSON")
	assert.Equal(t, map[string]string{"error": "Not found", "http_status": "404", "id": "1234", "what": "something", "value": "53c1e09c-bc7e-4989-8bc9-d71acdb0a013"}, payload)
}

func TestHTTPResponderWithStructuredErrorWithIdentifiable(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		err := StructuredErrorWithIdentifiable{ID: "1234", Message: "Not found", What: "something", Value: &Stuff{ID: uuid.MustParse("53c1e09c-bc7e-4989-8bc9-d71acdb0a013")}}
		core.RespondWithError(res, http.StatusNotFound, err)
	}))
	defer server.Close()

	res, err := http.Get(server.URL)
	require.NoErrorf(t, err, "Failed to get %s", server.URL)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	var payload map[string]string
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	require.NoErrorf(t, err, "Failed to read body of %s", server.URL)
	err = json.Unmarshal(body, &payload)
	require.NoError(t, err, "Failed to unmarshal JSON")
	assert.Equal(t, map[string]string{"error": "Not found", "http_status": "404", "id": "1234", "what": "something", "value": "53c1e09c-bc7e-4989-8bc9-d71acdb0a013"}, payload)
}

func TestHTTPResponderWithStructuredErrorWithUUID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		err := StructuredErrorWithUUID{ID: "1234", Message: "Not found", What: "something", Value: uuid.MustParse("53c1e09c-bc7e-4989-8bc9-d71acdb0a013")}
		core.RespondWithError(res, http.StatusNotFound, err)
	}))
	defer server.Close()

	res, err := http.Get(server.URL)
	require.NoErrorf(t, err, "Failed to get %s", server.URL)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	var payload map[string]string
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	require.NoErrorf(t, err, "Failed to read body of %s", server.URL)
	err = json.Unmarshal(body, &payload)
	require.NoError(t, err, "Failed to unmarshal JSON")
	assert.Equal(t, map[string]string{"error": "Not found", "http_status": "404", "id": "1234", "what": "something", "value": "53c1e09c-bc7e-4989-8bc9-d71acdb0a013"}, payload)
}

func TestHTTPResponderWithHTMLTemplate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		page, err := template.New("page").Parse("<html><body><h1>Hello {{.Name}}</h1></body></html>")
		require.NoError(t, err, "Failed to parse template")
		core.RespondWithHTMLTemplate(res, http.StatusOK, page, "page", map[string]string{"Name": "world"})
	}))
	defer server.Close()

	res, err := http.Get(server.URL)
	require.NoErrorf(t, err, "Failed to get %s", server.URL)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	require.NoErrorf(t, err, "Failed to read body of %s", server.URL)
	assert.Equal(t, "<html><body><h1>Hello world</h1></body></html>", string(body))
}

func TestHTTPResponderWithHTMLTemplateError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		page, err := template.New("page").Parse("<html><body><h1>Hello {{.Name}}</h1></body></html>")
		require.NoError(t, err, "Failed to parse template")
		core.RespondWithHTMLTemplate(res, http.StatusOK, page, "nowhere", map[string]string{"Name": "world"})
	}))
	defer server.Close()

	res, err := http.Get(server.URL)
	require.NoErrorf(t, err, "Failed to get %s", server.URL)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	var payload map[string]string
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	require.NoErrorf(t, err, "Failed to read body of %s", server.URL)
	err = json.Unmarshal(body, &payload)
	require.NoError(t, err, "Failed to unmarshal JSON")
	assert.Equal(t, map[string]string{"error": `html/template: "nowhere" is undefined`, "http_status": "500"}, payload)
}
