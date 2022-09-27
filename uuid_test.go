package core_test

import (
	"testing"

	. "github.com/gildas/go-core"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCanEncodeUUID(t *testing.T) {
	reqid, _ := uuid.Parse("b934e02c-96cb-4de3-ba7c-3b7daf593131")
	encoded := EncodeUUID(reqid)
	assert.NotEmpty(t, encoded, "encoded value is empty")
	assert.True(t, len(encoded) < 26, "encoded value is too long")
	t.Logf("Encoded reqid: %s (%d bytes)", encoded, len(encoded))

	decoded, err := DecodeUUID(encoded)
	assert.Nil(t, err, "Not a valid compressed UUID: %s", encoded)
	assert.Equal(t, reqid, decoded)
}
