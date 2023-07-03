package core_test

import (
	"encoding/json"
	"encoding/xml"
	"testing"

	"github.com/gildas/go-core"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCanEncodeUUID(t *testing.T) {
	reqid, _ := uuid.Parse("b934e02c-96cb-4de3-ba7c-3b7daf593131")
	encoded := core.EncodeUUID(reqid)
	assert.NotEmpty(t, encoded, "encoded value is empty")
	assert.True(t, len(encoded) < 26, "encoded value is too long")
	t.Logf("Encoded reqid: %s (%d bytes)", encoded, len(encoded))

	decoded, err := core.DecodeUUID(encoded)
	assert.NoError(t, err, "Not a valid compressed UUID: %s", encoded)
	assert.Equal(t, reqid, decoded)
}

func TestCanJSONMarshalUUID(t *testing.T) {
	reqid, _ := uuid.Parse("b934e02c-96cb-4de3-ba7c-3b7daf593131")
	payload, err := json.Marshal(struct {
		ID core.UUID `json:"id"`
	}{ID: core.UUID(reqid)})
	require.NoError(t, err, "Cannot marshal UUID")
	assert.Equal(t, []byte(`{"id":"`+reqid.String()+`"}`), payload)

	payload, err = json.Marshal(struct {
		ID string `json:"id,omitempty"`
	}{ID: core.UUID(reqid).String()})
	require.NoError(t, err, "Cannot marshal UUID")
	assert.Equal(t, []byte(`{"id":"`+reqid.String()+`"}`), payload)

	reqid = uuid.Nil
	payload, err = json.Marshal(struct {
		ID core.UUID `json:"id"`
	}{ID: core.UUID(reqid)})
	require.NoError(t, err, "Cannot marshal UUID")
	assert.Equal(t, []byte(`{"id":""}`), payload)

	payload, err = json.Marshal(struct {
		ID string `json:"id,omitempty"`
	}{ID: core.UUID(reqid).String()})
	require.NoError(t, err, "Cannot marshal UUID")
	assert.Equal(t, []byte(`{}`), payload)
}

func TestCanJSONUnmarshalUUID(t *testing.T) {
	var decoded struct {
		ID core.UUID `json:"id"`
	}
	reqid, _ := uuid.Parse("b934e02c-96cb-4de3-ba7c-3b7daf593131")
	payload := []byte(`{"id":"` + reqid.String() + `"}`)

	err := json.Unmarshal(payload, &decoded)
	require.NoError(t, err, "Cannot unmarshal UUID")
	assert.Equal(t, reqid, uuid.UUID(decoded.ID))

	err = json.Unmarshal([]byte(`{"id": ""}`), &decoded)
	require.NoError(t, err, "Cannot unmarshal UUID")
	assert.Equal(t, uuid.Nil, uuid.UUID(decoded.ID))

	err = json.Unmarshal([]byte(`{"id": null}`), &decoded)
	require.NoError(t, err, "Cannot unmarshal UUID")
	assert.Equal(t, uuid.Nil, uuid.UUID(decoded.ID))

	err = json.Unmarshal([]byte(`{}`), &decoded)
	require.NoError(t, err, "Cannot unmarshal UUID")
	assert.Equal(t, uuid.Nil, uuid.UUID(decoded.ID))
}

func TestCanXMLAttributeMarshalUUID(t *testing.T) {
	type Data struct {
		ID core.UUID `xml:"id,attr"`
	}
	reqid, _ := uuid.Parse("b934e02c-96cb-4de3-ba7c-3b7daf593131")
	payload, err := xml.Marshal(Data{ID: core.UUID(reqid)})
	require.NoError(t, err, "Cannot marshal UUID")
	assert.Equal(t, []byte(`<Data id="`+reqid.String()+`"></Data>`), payload)

	type Data02 struct {
		XMLName xml.Name `xml:"Data"`
		ID      string   `xml:"id,attr,omitempty"`
	}
	payload, err = xml.Marshal(Data02{ID: core.UUID(reqid).String()})
	require.NoError(t, err, "Cannot marshal UUID")
	assert.Equal(t, []byte(`<Data id="`+reqid.String()+`"></Data>`), payload)

	reqid = uuid.Nil
	payload, err = xml.Marshal(Data{ID: core.UUID(reqid)})
	require.NoError(t, err, "Cannot marshal UUID")
	assert.Equal(t, []byte(`<Data id=""></Data>`), payload)

	payload, err = xml.Marshal(Data02{ID: core.UUID(reqid).String()})
	require.NoError(t, err, "Cannot marshal UUID")
	assert.Equal(t, []byte(`<Data></Data>`), payload)
}

func TestCanXMLEntityMarshalUUID(t *testing.T) {
	type Data struct {
		ID core.UUID `xml:"id"`
	}
	reqid, _ := uuid.Parse("b934e02c-96cb-4de3-ba7c-3b7daf593131")
	payload, err := xml.Marshal(Data{ID: core.UUID(reqid)})
	require.NoError(t, err, "Cannot marshal UUID")
	assert.Equal(t, []byte(`<Data><id>`+reqid.String()+`</id></Data>`), payload)

	type Data02 struct {
		XMLName xml.Name `xml:"Data"`
		ID      string   `xml:"id,omitempty"`
	}
	payload, err = xml.Marshal(Data02{ID: core.UUID(reqid).String()})
	require.NoError(t, err, "Cannot marshal UUID")
	assert.Equal(t, []byte(`<Data><id>`+reqid.String()+`</id></Data>`), payload)

	reqid = uuid.Nil
	payload, err = xml.Marshal(Data{ID: core.UUID(reqid)})
	require.NoError(t, err, "Cannot marshal UUID")
	assert.Equal(t, []byte(`<Data><id></id></Data>`), payload)

	payload, err = xml.Marshal(Data02{ID: core.UUID(reqid).String()})
	require.NoError(t, err, "Cannot marshal UUID")
	assert.Equal(t, []byte(`<Data></Data>`), payload)
}

func TestCanXMLAttributeUnmarshalUUID(t *testing.T) {
	var decoded struct {
		ID core.UUID `xml:"id,attr"`
	}
	reqid, _ := uuid.Parse("b934e02c-96cb-4de3-ba7c-3b7daf593131")
	payload := []byte(`<Data id="` + reqid.String() + `" />}`)

	err := xml.Unmarshal(payload, &decoded)
	require.NoError(t, err, "Cannot unmarshal UUID")
	assert.Equal(t, reqid, uuid.UUID(decoded.ID))

	err = xml.Unmarshal([]byte(`<Data id="" />`), &decoded)
	require.NoError(t, err, "Cannot unmarshal UUID")
	assert.Equal(t, uuid.Nil, uuid.UUID(decoded.ID))

	err = xml.Unmarshal([]byte(`<Data/>`), &decoded)
	require.NoError(t, err, "Cannot unmarshal UUID")
	assert.Equal(t, uuid.Nil, uuid.UUID(decoded.ID))
}

func TestCanXMLEntityUnmarshalUUID(t *testing.T) {
	var decoded struct {
		ID core.UUID `xml:"id"`
	}
	reqid, _ := uuid.Parse("b934e02c-96cb-4de3-ba7c-3b7daf593131")
	payload := []byte(`<Data><id>` + reqid.String() + `</id></Data>}`)

	err := xml.Unmarshal(payload, &decoded)
	require.NoError(t, err, "Cannot unmarshal UUID")
	assert.Equal(t, reqid, uuid.UUID(decoded.ID))

	err = xml.Unmarshal([]byte(`<Data><id>""</id></Data>`), &decoded)
	require.NoError(t, err, "Cannot unmarshal UUID")
	assert.Equal(t, uuid.Nil, uuid.UUID(decoded.ID))

	err = xml.Unmarshal([]byte(`<Data><id></id></Data>`), &decoded)
	require.NoError(t, err, "Cannot unmarshal UUID")
	assert.Equal(t, uuid.Nil, uuid.UUID(decoded.ID))

	err = xml.Unmarshal([]byte(`<Data/>`), &decoded)
	require.NoError(t, err, "Cannot unmarshal UUID")
	assert.Equal(t, uuid.Nil, uuid.UUID(decoded.ID))
}
