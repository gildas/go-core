package core

import (
	"encoding/binary"
	"strings"

	"github.com/google/uuid"
)

// EncodeUUID encodes a UUID in a 22 char long string
func EncodeUUID(uuid uuid.UUID) string {
	bytes, _ := uuid.MarshalBinary()
	msi64 := binary.BigEndian.Uint64(bytes[:8])
	lsi64 := binary.BigEndian.Uint64(bytes[8:])

	result := strings.Builder{}
	result.WriteString(encodeUInt64(msi64))
	result.WriteString(encodeUInt64(lsi64))
	return result.String()
}

// DecodeUUID decodes a UUID from a 22 char long string
func DecodeUUID(encoded string) (uuid.UUID, error) {
	bytes := make([]byte, 16)
	binary.BigEndian.PutUint64(bytes,     decodeUInt64(encoded[:11]))
	binary.BigEndian.PutUint64(bytes[8:], decodeUInt64(encoded[11:]))
	return uuid.FromBytes(bytes)
}

func encodeUInt64(number uint64) string {
	base := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
	result := ""
	for i := 0; i < 11; i++ {
		result = string(base[number & 0x3f]) + result
		number = number >> 6
	}
	return result
}

func decodeUInt64(compressed string) uint64 {
	base := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
	result := uint64(0)
	for _, c := range compressed[:] {
		index := strings.IndexRune(base, c)
		result = (result << 6) | uint64(index)
	}
	return result
}