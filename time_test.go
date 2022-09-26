package core_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/gildas/go-core"
)

func TestTimeCanConvertToStandardTime(t *testing.T) {
	coreTime := Now()
	assert.False(t, coreTime.IsZero(), "Core Time from Now should not be zero")
}

func TestTimeCanConvertToString(t *testing.T) {
	loc, err := time.LoadLocation("America/Denver")
	assert.Nil(t, err, "Failed to load location")
	coreTime := Date(2006, time.January, 2, 15, 4, 5, 0, loc)
	assert.Equal(t, "2006-01-02 15:04:05 -0700 MST", coreTime.String())
}

func TestCanMarshalTime(t *testing.T) {
	jsonString := `{"time":"2006-01-02T22:04:05Z"}`
	loc, _ := time.LoadLocation("America/Denver")
	coreTime := Date(2006, time.January, 2, 15, 4, 5, 0, loc)
	data := struct {
		Time Time `json:"time"`
	}{Time: coreTime}
	payload, err := json.Marshal(data)
	require.Nil(t, err, "Failed to marshal Time")
	assert.Equal(t, jsonString, string(payload))
}

func TestCanUnmarshalTime(t *testing.T) {
	jsonString := `{"time":"2006-01-02T22:04:05Z"}`
	loc, _ := time.LoadLocation("America/Denver")
	coreTime := Date(2006, time.January, 2, 15, 4, 5, 0, loc)
	data := struct {
		Time Time `json:"time"`
	}{}
	err := json.Unmarshal([]byte(jsonString), &data)
	require.Nil(t, err, "Failed to unmarshal Time")
	assert.Equal(t, coreTime.UTC().String(), data.Time.String())
}
