package core_test

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/gildas/go-core"
)

func TestURLCanConvertToStandardURL(t *testing.T) {
	testURL, _ := url.Parse("https://www.acme.com")
	coreURL := (*URL)(testURL)
	assert.Equal(t, *testURL, coreURL.AsURL())
}

func TestURLCanConvertToString(t *testing.T) {
	testURL, _ := url.Parse("https://www.acme.com")
	coreURL := (*URL)(testURL)
	assert.Equal(t, "https://www.acme.com", coreURL.String())
}

func TestCanMarshalURLPtr(t *testing.T) {
	jsonString := `{"link":"https://www.acme.com"}`
	testURL, _ := url.Parse("https://www.acme.com")
	data := struct {
		Link *URL `json:"link"`
	}{(*URL)(testURL)}
	payload, err := json.Marshal(data)
	require.Nil(t, err, "Failed to marshal URL")
	assert.JSONEq(t, jsonString, string(payload))
}

func TestCanMarshalURL(t *testing.T) {
	jsonString := `{"link":"https://www.acme.com"}`
	testURL, _ := url.Parse("https://www.acme.com")
	data := struct {
		Link URL `json:"link"`
	}{URL(*testURL)}
	payload, err := json.Marshal(data)
	require.Nil(t, err, "Failed to marshal URL")
	assert.JSONEq(t, jsonString, string(payload))
}

func TestCanMarshalNilURL(t *testing.T) {
	jsonString := `{"id":"1234"}`
	data := struct {
		ID   string `json:"id"`
		Link *URL   `json:"link,omitempty"`
	}{ID: "1234", Link: nil}
	payload, err := json.Marshal(data)
	require.Nil(t, err, "Failed to marshal URL")
	assert.Equal(t, jsonString, string(payload))
	assert.JSONEq(t, jsonString, string(payload))
}

func TestCanUnmarshalURLPtr(t *testing.T) {
	jsonString := `{"link":"https://www.acme.com"}`
	testURL, _ := url.Parse("https://www.acme.com")
	data := struct {
		Link *URL `json:"link"`
	}{}
	err := json.Unmarshal([]byte(jsonString), &data)
	require.Nil(t, err, "Failed to unmarshal URL")
	parsedURL := (*url.URL)(data.Link)
	assert.Equal(t, testURL.String(), parsedURL.String())
}

func TestCanUnmarshalURL(t *testing.T) {
	jsonString := `{"link":"https://www.acme.com"}`
	testURL, _ := url.Parse("https://www.acme.com")
	data := struct {
		Link URL `json:"link"`
	}{}
	err := json.Unmarshal([]byte(jsonString), &data)
	require.Nil(t, err, "Failed to unmarshal URL")
	parsedURL := url.URL(data.Link)
	assert.Equal(t, testURL.String(), parsedURL.String())
}

func TestCanUnmarshalEmptyURL(t *testing.T) {
	jsonString := `{"id":"1234","link":""}`
	data := struct {
		ID   string `json:"1234"`
		Link *URL   `json:"link,omitempty"`
	}{}
	err := json.Unmarshal([]byte(jsonString), &data)
	require.Nil(t, err, "Failed to unmarshal URL")
	require.Empty(t, data.Link.String(), "Data Link should be empty")
}

func TestCanUnmarshalMissingURL(t *testing.T) {
	jsonString := `{"id":"1234"}`
	data := struct {
		ID   string `json:"1234"`
		Link *URL   `json:"link,omitempty"`
	}{}
	err := json.Unmarshal([]byte(jsonString), &data)
	require.Nil(t, err, "Failed to unmarshal URL")
	require.Nil(t, data.Link, "Data Link should be nil")
	parsedURL := (*url.URL)(data.Link)
	assert.Nil(t, parsedURL, "*url.URL equivalent should be nil")
}
