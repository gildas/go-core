package core

import (
	"encoding/binary"
	"strings"

	"github.com/google/uuid"
)

type UUID uuid.UUID

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
	binary.BigEndian.PutUint64(bytes, decodeUInt64(encoded[:11]))
	binary.BigEndian.PutUint64(bytes[8:], decodeUInt64(encoded[11:]))
	return uuid.FromBytes(bytes)
}

func encodeUInt64(number uint64) string {
	base := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
	result := ""
	for i := 0; i < 11; i++ {
		result = string(base[number&0x3f]) + result
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

// IsZero returns true if the UUID is uuid.Nil
func (id UUID) IsZero() bool {
	return uuid.UUID(id) == uuid.Nil
}

// String returns the UUID as a string
//
// If the UUID is uuid.Nil, an empty string is returned.
//
// Implements the fmt.Stringer interface
func (id UUID) String() string {
	if id.IsZero() {
		return ""
	}
	return uuid.UUID(id).String()
}

// MarshalText marshals the UUID as a string
//
// Implements the json.Marshaler interface
func (id UUID) MarshalText() ([]byte, error) {
	if id.IsZero() {
		return nil, nil
	}
	return []byte(uuid.UUID(id).String()), nil
}

// UnmarshalText unmarshals the UUID from a string
//
// If the string is empty or null, the UUID is set to uuid.Nil.
//
// Implements the json.Unmarshaler interface
func (id *UUID) UnmarshalText(payload []byte) (err error) {
	if len(payload) == 0 || string(payload) == `""` || string(payload) == `null` {
		*id = UUID(uuid.Nil)
		return nil
	}
	parsed, err := uuid.Parse(string(payload))
	if err == nil {
		*id = UUID(parsed)
	}
	return err
}
