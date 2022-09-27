package core_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/gildas/go-core"
)

func TestTimeCanInitialize(t *testing.T) {
	now := core.NowUTC()
	assert.Contains(t, now.String(), "UTC")
	loc, _ := time.LoadLocation("Asia/Tokyo")
	now = core.NowIn(loc)
	assert.Contains(t, now.String(), "JST")
	now = core.Date(2006, time.January, 2, 15, 4, 5, 0, loc)
	assert.Equal(t, "2006-01-02 15:04:05 +0900 JST", now.String())

	year, month, day := now.Date()
	assert.Equal(t, 2006, year)
	assert.Equal(t, time.January, month)
	assert.Equal(t, 2, day)
	assert.Equal(t, loc, now.Location())
	assert.Equal(t, "2006-01-02T15:04:05+09:00", now.Format(time.RFC3339))
	assert.Equal(t, "2006-01-02 00:00:00 +0900 JST", now.BeginOfDay().String())
	assert.Equal(t, "2006-01-02 23:59:59 +0900 JST", now.EndOfDay().String())
	assert.Equal(t, "2006-01-03 15:04:05 +0900 JST", now.Tomorrow().String())
	assert.Equal(t, "2006-01-01 15:04:05 +0900 JST", now.Yesterday().String())
	assert.True(t, now.Before(now.Tomorrow()))
	assert.True(t, now.After(now.Yesterday()))
	assert.True(t, now.Equal(core.DateUTC(2006, time.January, 2, 6, 4, 5, 0)))
}

func TestTimeCanConvertToStandardTime(t *testing.T) {
	coreTime := core.Now()
	assert.False(t, coreTime.IsZero(), "Core Time from Now should not be zero")
}

func TestTimeCanConvertToString(t *testing.T) {
	loc, err := time.LoadLocation("America/Denver")
	assert.NoError(t, err, "Failed to load location")
	coreTime := core.Date(2006, time.January, 2, 15, 4, 5, 0, loc)
	assert.Equal(t, "2006-01-02 15:04:05 -0700 MST", coreTime.String())
}

func TestCanMarshalTime(t *testing.T) {
	jsonString := `{"time":"2006-01-02T22:04:05Z"}`
	loc, _ := time.LoadLocation("America/Denver")
	coreTime := core.Date(2006, time.January, 2, 15, 4, 5, 0, loc)
	data := struct {
		Time core.Time `json:"time"`
	}{Time: coreTime}
	payload, err := json.Marshal(data)
	require.NoError(t, err, "Failed to marshal Time")
	assert.Equal(t, jsonString, string(payload))
}

func TestCanUnmarshalTime(t *testing.T) {
	jsonString := `{"time":"2006-01-02T22:04:05Z"}`
	loc, _ := time.LoadLocation("America/Denver")
	coreTime := core.Date(2006, time.January, 2, 15, 4, 5, 0, loc)
	data := struct {
		Time core.Time `json:"time"`
	}{}
	err := json.Unmarshal([]byte(jsonString), &data)
	require.NoError(t, err, "Failed to unmarshal Time")
	assert.Equal(t, coreTime.UTC().String(), data.Time.String())
}

func TestCanUnmarshalTimeWithEmptyString(t *testing.T) {
	jsonString := `{"time":""}`
	coreTime := core.Time(time.Time{})
	data := struct {
		Time core.Time `json:"time"`
	}{}
	err := json.Unmarshal([]byte(jsonString), &data)
	require.NoError(t, err, "Failed to unmarshal Time")
	assert.Equal(t, coreTime.UTC().String(), data.Time.String())
}

func TestShouldFailUnmarshalTimeWithInvalidPayload(t *testing.T) {
	var err error
	data := struct {
		Time core.Time `json:"time"`
	}{}

	err = json.Unmarshal([]byte(`{"time":"hello"}`), &data)
	require.Error(t, err, "should fail to unmarshal Time")
	assert.Equal(t, `parsing time "hello" as "2006-01-02T15:04:05Z07:00": cannot parse "hello" as "2006"`, err.Error())

	err = json.Unmarshal([]byte(`{"time":12}`), &data)
	require.Error(t, err, "should fail to unmarshal Time")
	assert.Equal(t, "json: cannot unmarshal number into Go struct field .time of type string", err.Error())
}

func TestCanParseTime(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Tokyo")

	now, err := core.ParseTime("2006-01-02T15:04:05Z")
	require.NoError(t, err)
	assert.Equal(t, "2006-01-02 15:04:05 +0000 UTC", now.String())

	_, err = core.ParseTimeIn("now", loc)
	require.NoError(t, err)

	_, err = core.ParseTimeIn("today", loc)
	require.NoError(t, err)

	_, err = core.ParseTimeIn("tomorrow", loc)
	require.NoError(t, err)

	_, err = core.ParseTimeIn("yesterday", loc)
	require.NoError(t, err)

	_, err = core.ParseTimeIn("2006-01-02T15:04:05Z", loc)
	require.NoError(t, err)

	_, err = core.ParseTimeIn("2006-01-02T15:04:05+07:00", loc)
	require.NoError(t, err)

	_, err = core.ParseTimeIn("T15:04:05+07:00", loc)
	require.NoError(t, err)

	_, err = core.ParseTimeIn("T15:04:05Z", loc)
	require.NoError(t, err)

	_, err = core.ParseTimeIn("T15:04:05", loc)
	require.NoError(t, err)

	_, err = core.ParseTimeIn("2006-01-02", loc)
	require.NoError(t, err)
}
