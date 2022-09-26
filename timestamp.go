package core

import (
	"strconv"
	"strings"
	"time"
)

// Timestamp converts Unix Epoch to/from time.Time
type Timestamp time.Time

// TimestampNow returns a Timestamp at the time of its call
func TimestampNow() Timestamp {
	return Timestamp(time.Now())
}

// TimestampFromJSEpoch returns a Timestamp from a JS Epoch
func TimestampFromJSEpoch(epoch int64) Timestamp {
	return Timestamp(time.Unix(epoch/1000, (epoch%1000)*1000000))
}

// MarshalJSON encodes a TimeStamp to its JSON Epoch
//   implements json.Marshaler interface
func (t Timestamp) MarshalJSON() ([]byte, error) {
	// Per Node.js epochs, we need milliseconds
	return []byte(strconv.FormatInt(t.JSEpoch(), 10)), nil
}

// UnmarshalJSON decodes an Epoch from JSON and gives a Timestamp
//   implements json.Unmarshaler interface
//
//   The Epoch can be "12345" or 12345
func (t *Timestamp) UnmarshalJSON(payload []byte) (err error) {
	// First get rid of the surrounding double quotes
	unquoted := strings.Replace(string(payload), `"`, ``, -1)
	value, err := strconv.ParseInt(unquoted, 10, 64)

	if err != nil {
		return err
	}
	// Per Node.js epochs, value will be in milliseconds
	*t = TimestampFromJSEpoch(value)
	return
}

// JSEpoch returns the Unix Epoch like Javascript (i.e. in ms)
func (t Timestamp) JSEpoch() int64 {
	return time.Time(t).UnixNano() / int64(1000000)
}

// String gives the string representation of this
func (t Timestamp) String() string { return time.Time(t).String() }
