package core

import (
	"strconv"
	"strings"
)

// Got the technique from:
// https://engineering.bitnami.com/articles/dealing-with-json-with-non-homogeneous-types-in-go.html

// FlexInt is an int that can be unmashaled from an int or a string (1234 or "1234")
type FlexInt int

// FlexInt8 is an int that can be unmashaled from an int or a string (1234 or "1234")
type FlexInt8 int8

// FlexInt16 is an int that can be unmashaled from an int or a string (1234 or "1234")
type FlexInt16 int16

// FlexInt32 is an int that can be unmashaled from an int or a string (1234 or "1234")
type FlexInt32 int32

// FlexInt64 is an int that can be unmashaled from an int or a string (1234 or "1234")
type FlexInt64 int64

// UnmarshalJSON decodes JSON
//
//	implements json.Unmarshaler interface
func (i *FlexInt) UnmarshalJSON(payload []byte) error {
	// First get rid of the surrounding double quotes
	unquoted := strings.Replace(string(payload), `"`, ``, -1)
	value, err := strconv.ParseInt(unquoted, 10, 64)
	if err != nil {
		return err
	}
	*i = FlexInt(value)
	return nil
}

// UnmarshalJSON decodes JSON
//
//	implements json.Unmarshaler interface
func (i *FlexInt8) UnmarshalJSON(payload []byte) error {
	// First get rid of the surrounding double quotes
	unquoted := strings.Replace(string(payload), `"`, ``, -1)
	value, err := strconv.ParseInt(unquoted, 10, 64)
	if err != nil {
		return err
	}
	*i = FlexInt8(value)
	return nil
}

// UnmarshalJSON decodes JSON
//
//	implements json.Unmarshaler interface
func (i *FlexInt16) UnmarshalJSON(payload []byte) error {
	// First get rid of the surrounding double quotes
	unquoted := strings.Replace(string(payload), `"`, ``, -1)
	value, err := strconv.ParseInt(unquoted, 10, 64)
	if err != nil {
		return err
	}
	*i = FlexInt16(value)
	return nil
}

// UnmarshalJSON decodes JSON
//
//	implements json.Unmarshaler interface
func (i *FlexInt32) UnmarshalJSON(payload []byte) error {
	// First get rid of the surrounding double quotes
	unquoted := strings.Replace(string(payload), `"`, ``, -1)
	value, err := strconv.ParseInt(unquoted, 10, 64)
	if err != nil {
		return err
	}
	*i = FlexInt32(value)
	return nil
}

// UnmarshalJSON decodes JSON
//
//	implements json.Unmarshaler interface
func (i *FlexInt64) UnmarshalJSON(payload []byte) error {
	// First get rid of the surrounding double quotes
	unquoted := strings.Replace(string(payload), `"`, ``, -1)
	value, err := strconv.ParseInt(unquoted, 10, 64)
	if err != nil {
		return err
	}
	*i = FlexInt64(value)
	return nil
}
