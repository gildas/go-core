package core_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/gildas/go-core"
)

func TestCanStringifyDuration(t *testing.T) {
	duration := core.Duration(120 * time.Second)
	assert.Equal(t, "2m0s", duration.String())
}

func TestCanMarshalDuration(t *testing.T) {
	duration := core.Duration(120 * time.Second)
	payload, err := json.Marshal(duration)
	require.Nil(t, err, "Failed to marshal duration")
	assert.Equal(t, "120000", string(payload))
}

func TestCanUnmarshalDurationFromInt(t *testing.T) {
	payload := `{"duration":120000}`
	result := struct {
		Duration core.Duration `json:"duration"`
	}{}
	err := json.Unmarshal([]byte(payload), &result)
	require.Nil(t, err, "Failed to unmarshal duration, %s", err)
	assert.Equal(t, 120*time.Second, time.Duration(result.Duration))
}

func TestCanUnmarshalDurationFromFloat(t *testing.T) {
	payload := `{"duration":120000.0}`
	result := struct {
		Duration core.Duration `json:"duration"`
	}{}
	err := json.Unmarshal([]byte(payload), &result)
	require.Nil(t, err, "Failed to unmarshal duration, %s", err)
	assert.Equal(t, 120*time.Second, time.Duration(result.Duration))
}

func TestCanUnmarshalDurationFromISO(t *testing.T) {
	result := struct {
		Duration core.Duration `json:"duration"`
	}{}

	err := json.Unmarshal([]byte(`{"duration":"PT2H30M15S"}`), &result)
	require.Nil(t, err, "Failed to unmarshal duration, %s", err)
	assert.Equal(t, 2*time.Hour+30*time.Minute+15*time.Second, time.Duration(result.Duration))

	err = json.Unmarshal([]byte(`{"duration":"PT0.005S"}`), &result)
	require.Nil(t, err, "Failed to unmarshal duration, %s", err)
	assert.Equal(t, 5*time.Millisecond, time.Duration(result.Duration))

	err = json.Unmarshal([]byte(`{"duration":"P2D"}`), &result)
	require.Nil(t, err, "Failed to unmarshal duration, %s", err)
	assert.Equal(t, 2*24*time.Hour, time.Duration(result.Duration))

	err = json.Unmarshal([]byte(`{"duration":"P2W"}`), &result)
	require.Nil(t, err, "Failed to unmarshal duration, %s", err)
	assert.Equal(t, 2*7*24*time.Hour, time.Duration(result.Duration))

	err = json.Unmarshal([]byte(`{"duration":"P2Y"}`), &result)
	require.Nil(t, err, "Failed to unmarshal duration, %s", err)
	assert.Equal(t, 2*365*24*time.Hour, time.Duration(result.Duration))

	err = json.Unmarshal([]byte(`{"duration":"P1Y2M3DT4H5M6S"}`), &result)
	require.Nil(t, err, "Failed to unmarshal duration, %s", err)
	assert.Equal(t, 10276*time.Hour+5*time.Minute+6*time.Second, time.Duration(result.Duration))
}

func TestCanUnmarshalDurationFromGO(t *testing.T) {
	payload := `{"duration":"48h"}`
	result := struct {
		Duration core.Duration `json:"duration"`
	}{}
	err := json.Unmarshal([]byte(payload), &result)
	require.Nil(t, err, "Failed to unmarshal duration, %s", err)
	assert.Equal(t, 2*24*time.Hour, time.Duration(result.Duration))
}

func TestShouldFailUnmarshalDurationWithInvalidPayload(t *testing.T) {
	var err error
	var duration core.Duration

	err = json.Unmarshal([]byte(`"P5"`), &duration)
	require.Error(t, err, "Should have failed to unmarshal")
	assert.Equal(t, `"P5" is not an ISO8601 duration`, err.Error())
	err = json.Unmarshal([]byte(`"5000ts"`), &duration)
	require.Error(t, err, "Should have failed to unmarshal")
	assert.Equal(t, `time: unknown unit "ts" in duration "5000ts"`, err.Error())
	err = json.Unmarshal([]byte(`[]`), &duration)
	require.Error(t, err, "Should have failed to unmarshal")
	assert.Equal(t, "Invalid Duration", err.Error())
}

func TestDurationShouldFailParseWithInvalidData(t *testing.T) {
	var err error

	_, err = core.ParseDuration("hello")
	require.Error(t, err, "Should have failed to parse")
	assert.Equal(t, `time: invalid duration "hello"`, err.Error())

	_, err = core.ParseDuration("P5")
	require.Error(t, err, "Should have failed to parse")
	assert.Equal(t, `"P5" is not an ISO8601 duration`, err.Error())
}
