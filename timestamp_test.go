package core_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/gildas/go-core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimestampCanInitialize(t *testing.T) {
	epoch := int64(1534318964318) // milliseconds
	timestamp := core.Timestamp(time.Unix(0, epoch*1000000))
	assert.Equal(t, epoch, timestamp.JSEpoch())

	timestamp = core.TimestampNow()
	assert.NotEmpty(t, timestamp.String())
}

func TestCanMarshalTimestamp(t *testing.T) {
	epoch := int64(1534318964318) // milliseconds
	expected := fmt.Sprintf(`{"timestamp":%d}`, epoch)
	data := struct {
		Timestamp core.Timestamp `json:"timestamp"`
	}{
		Timestamp: core.Timestamp(time.Unix(0, epoch*1000000)),
	}
	payload, err := json.Marshal(data)
	assert.NoError(t, err, "Failed Unmarshaling: %s", err)
	assert.JSONEq(t, expected, string(payload))
}

func TestCanUnmarshalTimestampWithInt(t *testing.T) {
	epoch := int64(1534318964318) // milliseconds
	payload := fmt.Sprintf(`{"timestamp":%d}`, epoch)
	result := &struct {
		Timestamp core.Timestamp `json:"timestamp"`
	}{}
	err := json.Unmarshal([]byte(payload), &result)
	assert.NoError(t, err, "Failed Unmarshaling: %s", err)
	assert.Equal(t, epoch, result.Timestamp.JSEpoch())
}

func TestCanUnmarshalTimestampWithString(t *testing.T) {
	epoch := int64(1534318964318) // milliseconds
	payload := fmt.Sprintf(`{"timestamp":"%d"}`, epoch)
	result := &struct {
		Timestamp core.Timestamp `json:"timestamp"`
	}{}
	err := json.Unmarshal([]byte(payload), &result)
	assert.NoError(t, err, "Failed Unmarshaling: %s", err)
	assert.Equal(t, epoch, result.Timestamp.JSEpoch())
}

func TestShouldFailUnmarshalTimestampWithInvalidPayload(t *testing.T) {
	payload := `{"timestamp":"hello"}`
	result := &struct {
		Timestamp core.Timestamp `json:"timestamp"`
	}{}
	err := json.Unmarshal([]byte(payload), &result)
	require.Error(t, err, "should fail to unmarshal")
	assert.Equal(t, `strconv.ParseInt: parsing "hello": invalid syntax`, err.Error())
}
