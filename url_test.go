package core_test

import (
	"encoding/json"
	"testing"
	"net/url"

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
	assert.Equal(t, jsonString, string(payload))
}

func TestCanMarshalURL(t *testing.T) {
	jsonString := `{"link":"https://www.acme.com"}`
	testURL, _ := url.Parse("https://www.acme.com")
	data := struct {
		Link URL `json:"link"`
	}{URL(*testURL)}
	payload, err := json.Marshal(data)
	require.Nil(t, err, "Failed to marshal URL")
	assert.Equal(t, jsonString, string(payload))
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