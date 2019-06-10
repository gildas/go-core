package core_test

import (
	"gotest.tools/assert"
	"github.com/stretchr/testify/require"
	"encoding/json"
	"testing"
	"time"
	. "github.com/gildas/go-core"
)

func TestCanMarshalDuration(t *testing.T) {
	duration := Duration(120*time.Second)
	payload, err := json.Marshal(duration)
	require.Nil(t, err, "Failed to marshal duration")
	assert.Equal(t, "120000", string(payload))
}

func TestCanUnmarshalDurationFromInt(t *testing.T) {
	payload := `{"duration":120000}`
	result  := struct{Duration Duration `json:"duration"`}{}
	err     := json.Unmarshal([]byte(payload), &result)
	require.Nil(t, err, "Failed to unmarshal duration, %s", err)
	assert.Equal(t, 120*time.Second, time.Duration(result.Duration))
}

func TestCanUnmarshalDurationFromISO(t *testing.T) {
	payload := `{"duration":"P2H30M15S"}`
	result  := struct{Duration Duration `json:"duration"`}{}
	err     := json.Unmarshal([]byte(payload), &result)
	require.Nil(t, err, "Failed to unmarshal duration P2H30M15S, %s", err)
	assert.Equal(t, 2*time.Hour+30*time.Minute+15*time.Second, time.Duration(result.Duration))

	payload = `{"duration":"P2D"}`
	result  = struct{Duration Duration `json:"duration"`}{}
	err     = json.Unmarshal([]byte(payload), &result)
	require.Nil(t, err, "Failed to unmarshal duration P2D, %s", err)
	assert.Equal(t, 2*24*time.Hour, time.Duration(result.Duration))

	payload = `{"duration":"P2W"}`
	result  = struct{Duration Duration `json:"duration"`}{}
	err     = json.Unmarshal([]byte(payload), &result)
	require.Nil(t, err, "Failed to unmarshal duration P2W, %s", err)
	assert.Equal(t, 2*7*24*time.Hour, time.Duration(result.Duration))

	payload = `{"duration":"P2Y"}`
	result  = struct{Duration Duration `json:"duration"`}{}
	err     = json.Unmarshal([]byte(payload), &result)
	require.Nil(t, err, "Failed to unmarshal duration P2Y, %s", err)
	assert.Equal(t, 2*365*24*time.Hour, time.Duration(result.Duration))
}

func TestCanUnmarshalDurationFromGO(t *testing.T) {
	payload := `{"duration":"48h"}`
	result  := struct{Duration Duration `json:"duration"`}{}
	err     := json.Unmarshal([]byte(payload), &result)
	require.Nil(t, err, "Failed to unmarshal duration, %s", err)
	assert.Equal(t, 2*24*time.Hour, time.Duration(result.Duration))
}
