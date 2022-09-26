package core_test

import (
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/gildas/go-core"
	"github.com/stretchr/testify/assert"
)

func TestTimestampIntUnmarshal(t *testing.T) {
	epoch := int64(1534318964318)
	payload := fmt.Sprintf(`{"timestamp":%d}`, epoch)
	result := &struct {
		Timestamp Timestamp `json:"timestamp"`
	}{}
	err := json.Unmarshal([]byte(payload), &result)

	assert.Nil(t, err, "Failed Unmarshaling: %s", err)
	assert.Equal(t, epoch, result.Timestamp.JSEpoch())
}

func TestTimestampStringUnmarshal(t *testing.T) {
	epoch := int64(1534318964318)
	payload := fmt.Sprintf(`{"timestamp":"%d"}`, epoch)
	result := &struct {
		Timestamp Timestamp `json:"timestamp"`
	}{}
	err := json.Unmarshal([]byte(payload), &result)

	assert.Nil(t, err, "Failed Unmarshaling: %s", err)
	assert.Equal(t, epoch, result.Timestamp.JSEpoch())
}
