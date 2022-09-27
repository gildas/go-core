package core_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/gildas/go-core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAtoi(t *testing.T) {
	assert.Equal(t, 12, core.Atoi("12", 0))
	assert.Equal(t, 5, core.Atoi("hello", 5))
}

func TestCanUnmarshalFlextInt(t *testing.T) {
	var err error
	var value core.FlexInt

	err = json.Unmarshal([]byte("12"), &value)
	require.NoError(t, err, "should not have failed to unmarshal")
	assert.Equal(t, 12, int(value))

	err = json.Unmarshal([]byte(`"12"`), &value)
	require.NoError(t, err, "should not have failed to unmarshal")
	assert.Equal(t, 12, int(value))

	err = json.Unmarshal([]byte(`"hello"`), &value)
	require.Error(t, err, "should have failed to unmarshal")
	assert.Equal(t, `strconv.ParseInt: parsing "hello": invalid syntax`, err.Error())
}

func TestCanUnmarshalFlextInt8(t *testing.T) {
	var err error
	var value core.FlexInt8

	err = json.Unmarshal([]byte("12"), &value)
	require.NoError(t, err, "should not have failed to unmarshal")
	assert.Equal(t, int8(12), int8(value))

	err = json.Unmarshal([]byte(`"12"`), &value)
	require.NoError(t, err, "should not have failed to unmarshal")
	assert.Equal(t, int8(12), int8(value))

	err = json.Unmarshal([]byte(`"hello"`), &value)
	require.Error(t, err, "should have failed to unmarshal")
	assert.Equal(t, `strconv.ParseInt: parsing "hello": invalid syntax`, err.Error())
}

func TestCanUnmarshalFlextInt16(t *testing.T) {
	var err error
	var value core.FlexInt16

	err = json.Unmarshal([]byte("12"), &value)
	require.NoError(t, err, "should not have failed to unmarshal")
	assert.Equal(t, int16(12), int16(value))

	err = json.Unmarshal([]byte(`"12"`), &value)
	require.NoError(t, err, "should not have failed to unmarshal")
	assert.Equal(t, int16(12), int16(value))

	err = json.Unmarshal([]byte(`"hello"`), &value)
	require.Error(t, err, "should have failed to unmarshal")
	assert.Equal(t, `strconv.ParseInt: parsing "hello": invalid syntax`, err.Error())
}

func TestCanUnmarshalFlextInt32(t *testing.T) {
	var err error
	var value core.FlexInt32

	err = json.Unmarshal([]byte("12"), &value)
	require.NoError(t, err, "should not have failed to unmarshal")
	assert.Equal(t, int32(12), int32(value))

	err = json.Unmarshal([]byte(`"12"`), &value)
	require.NoError(t, err, "should not have failed to unmarshal")
	assert.Equal(t, int32(12), int32(value))

	err = json.Unmarshal([]byte(`"hello"`), &value)
	require.Error(t, err, "should have failed to unmarshal")
	assert.Equal(t, `strconv.ParseInt: parsing "hello": invalid syntax`, err.Error())
}

func TestCanUnmarshalFlextInt64(t *testing.T) {
	var err error
	var value core.FlexInt64

	err = json.Unmarshal([]byte("12"), &value)
	require.NoError(t, err, "should not have failed to unmarshal")
	assert.Equal(t, int64(12), int64(value))

	err = json.Unmarshal([]byte(`"12"`), &value)
	require.NoError(t, err, "should not have failed to unmarshal")
	assert.Equal(t, int64(12), int64(value))

	err = json.Unmarshal([]byte(`"hello"`), &value)
	require.Error(t, err, "should have failed to unmarshal")
	assert.Equal(t, `strconv.ParseInt: parsing "hello": invalid syntax`, err.Error())
}

func TestCanExecEvery(t *testing.T) {
	var count int
	stopme, pingme, changeme := core.ExecEvery(func(tick int64, at time.Time, changeme chan time.Duration) {
		count++
		t.Logf("Tick #%d at %s, count is now: %d", tick, at, count)
	}, 200*time.Millisecond)
	defer close(stopme)
	time.Sleep(350 * time.Millisecond)
	assert.Equal(t, 2, count)
	pingme <- true
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, 3, count)
	changeme <- 300 * time.Millisecond
	time.Sleep(350 * time.Millisecond)
	assert.Equal(t, 4, count)
}
