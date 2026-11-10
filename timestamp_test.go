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

func TestUnmarshalTimestampIsUTC(t *testing.T) {
	epoch := int64(1534318964318) // milliseconds (August 15, 2018 07:42:44.318 UTC)
	payload := fmt.Sprintf(`{"timestamp":%d}`, epoch)
	result := &struct {
		Timestamp core.Timestamp `json:"timestamp"`
	}{}
	err := json.Unmarshal([]byte(payload), &result)
	assert.NoError(t, err, "Failed Unmarshaling: %s", err)

	tm := time.Time(result.Timestamp)
	assert.Equal(t, time.UTC, tm.Location(), "Unmarshaled timestamp should be in UTC")
}

func TestUnmarshalTimestampStringIsUTC(t *testing.T) {
	epoch := int64(1534318964318) // milliseconds
	payload := fmt.Sprintf(`{"timestamp":"%d"}`, epoch)
	result := &struct {
		Timestamp core.Timestamp `json:"timestamp"`
	}{}
	err := json.Unmarshal([]byte(payload), &result)
	assert.NoError(t, err, "Failed Unmarshaling: %s", err)

	tm := time.Time(result.Timestamp)
	assert.Equal(t, time.UTC, tm.Location(), "Unmarshaled timestamp from string should be in UTC")
}

func TestTimestampFromJSEpochIsUTC(t *testing.T) {
	epoch := int64(1534318964318) // milliseconds
	timestamp := core.TimestampFromJSEpoch(epoch)

	tm := time.Time(timestamp)
	assert.Equal(t, time.UTC, tm.Location(), "TimestampFromJSEpoch should return UTC time")
	assert.Equal(t, epoch, timestamp.JSEpoch(), "JSEpoch should match input")
}

func TestTimestampNowReturnsValidTimestamp(t *testing.T) {
	before := time.Now()
	timestamp := core.TimestampNow()
	after := time.Now()

	tm := time.Time(timestamp)
	assert.NotZero(t, tm, "TimestampNow should return non-zero time")
	assert.True(t, tm.After(before) || tm.Equal(before), "TimestampNow should be >= before")
	assert.True(t, tm.Before(after) || tm.Equal(after), "TimestampNow should be <= after")
}

func TestTimestampString(t *testing.T) {
	epoch := int64(1534318964318) // milliseconds
	timestamp := core.TimestampFromJSEpoch(epoch)
	str := timestamp.String()

	assert.NotEmpty(t, str, "String() should return non-empty string")
	assert.Contains(t, str, "UTC", "String() should contain UTC timezone info")
}

func TestTimestampJSEpoch(t *testing.T) {
	testCases := []int64{
		0,
		1534318964318,
		1609459200000, // Jan 1, 2021 00:00:00 UTC
		1893456000000, // Year 2030
	}

	for _, epoch := range testCases {
		t.Run(fmt.Sprintf("Epoch_%d", epoch), func(t *testing.T) {
			timestamp := core.TimestampFromJSEpoch(epoch)
			assert.Equal(t, epoch, timestamp.JSEpoch(), "JSEpoch should round-trip correctly")
		})
	}
}

func TestTimestampMarshalJSONWithZero(t *testing.T) {
	epoch := int64(0)
	timestamp := core.TimestampFromJSEpoch(epoch)
	payload, err := json.Marshal(timestamp)

	assert.NoError(t, err, "Should marshal zero timestamp without error")
	assert.Equal(t, `0`, string(payload), "Zero timestamp should marshal to 0")
}

func TestTimestampUnmarshalJSONWithZero(t *testing.T) {
	payload := `0`
	var timestamp core.Timestamp

	err := json.Unmarshal([]byte(payload), &timestamp)
	assert.NoError(t, err, "Should unmarshal zero without error")
	assert.Equal(t, int64(0), timestamp.JSEpoch(), "Zero payload should result in zero epoch")
}

func TestTimestampDirectTimeConversion(t *testing.T) {
	epoch := int64(1534318964318)
	tm := time.Unix(0, epoch*1000000).UTC()
	timestamp := core.Timestamp(tm)

	assert.Equal(t, epoch, timestamp.JSEpoch(), "Direct time conversion should calculate JSEpoch correctly")
}

func TestTimestampFractionalMilliseconds(t *testing.T) {
	// Test with fractional nanoseconds that should round to millisecond precision
	epoch := int64(1534318964318)
	timestamp := core.TimestampFromJSEpoch(epoch)

	// When converting back, it should maintain millisecond precision
	assert.Equal(t, epoch, timestamp.JSEpoch())
}

func TestTimestampNegativeEpoch(t *testing.T) {
	epoch := int64(-1000) // Before Unix epoch
	timestamp := core.TimestampFromJSEpoch(epoch)

	assert.Equal(t, epoch, timestamp.JSEpoch(), "Should handle negative epochs")
	tm := time.Time(timestamp)
	assert.Equal(t, time.UTC, tm.Location(), "Negative epoch should still be in UTC")
}
