package core

import (
	"encoding/json"
	"strings"
	"time"
)

// Time is a placeholder so we can add new funcs to the type
type Time time.Time

// AsTime converts a core.Time into a time.Time
func (t Time) AsTime() time.Time {
	return (time.Time)(t)
}

// Now returns the current local time
func Now() Time {
	return (Time)(time.Now())
}

// NowIn returns the current time in the given location
func NowIn(loc *time.Location) Time {
	return (Time)(time.Now().In(loc))
}

// NowUTC returns the current UTC time
func NowUTC() Time {
	return (Time)(time.Now().UTC())
}

// Date returns a new Date
func Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	return (Time)(time.Date(year, month, day, hour, min, sec, nsec, loc))
}

// Date returns the year, month, and day in which t occurs
func (t Time) Date() (year int, month time.Month, day int) {
	return t.AsTime().Date()
}

// UTC returns t with the location set to UTC
func (t Time) UTC() Time {
	return (Time)(t.AsTime().UTC())
}

// IsZero reports whether t represents the zero time instant, January 1, year 1, 00:00:00 UTC
func (t Time) IsZero() bool {
	return t.AsTime().IsZero()
}

// Location returns the time zone information associated with t
func (t Time) Location() *time.Location {
	return t.AsTime().Location()
}

// BeginOfDay returns the Beginning of the Day (i.e. midnight)
func (t Time) BeginOfDay() Time {
	tt := t.AsTime()
	year, month, day := t.AsTime().Date()
	return Date(year, month, day, 0, 0, 0, 0, tt.Location())
}

// EndOfDay returns the End of the Day (i.e. 1 second before midnight)
func (t Time) EndOfDay() Time {
	tt := t.AsTime()
	year, month, day := t.AsTime().Date()
	return (Time)(time.Date(year, month, day, 0, 0, 0, 0, tt.Location()))
}

// Tomorrow returns t shifted to tomorrow
func (t Time) Tomorrow() Time {
	return (Time)(t.AsTime().Add(24 *time.Hour))
}

// Yesterday returns t shifted to yesterday
func (t Time) Yesterday() Time {
	return (Time)(t.AsTime().Add(-24 *time.Hour))
}

// After reports whether the time instant t is after u
func (t Time) After(u Time) bool {
	return t.AsTime().After(u.AsTime())
}

// Before reports whether the time instant t is before u
func (t Time) Before(u Time) bool {
	return t.AsTime().Before(u.AsTime())
}

// Equal reports whether t and u represent the same time instant.
// Two times can be equal even if they are in different locations.
// For example, 6:00 +0200 and 4:00 UTC are Equal.
// See the documentation on the Time type for the pitfalls of using == with Time values; most code should use Equal instead.
func (t Time) Equal(u Time) bool {
	return t.AsTime().Equal(u.AsTime())
}

// ParseTime parses the given string for a Time, if the Time is not UTC it is set in the current location
func ParseTime(value string) (Time, error) {
	return ParseTimeIn(value, time.Now().Location())
}

// ParseTimeIn parses the given string for a Time, if the Time is not UTC it is set in the given location
func ParseTimeIn(value string, loc *time.Location) (Time, error) {
	now := NowIn(loc)
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "now":
		return now, nil
	case "today":
		return now.BeginOfDay(), nil
	case "tomorrow":
		return now.BeginOfDay().Tomorrow(), nil
	case "yesterday":
		return now.BeginOfDay().Yesterday(), nil
	}

	var parsed time.Time
	var err error

	if strings.HasPrefix(value, "T") { // value contains only time information
		if strings.HasSuffix(value, "Z") {
			parsed, err = time.Parse("T15:04:05Z", value)
		} else if strings.ContainsAny(value, "+-") { // value contains a time zone
			parsed, err = time.Parse("T15:04:05-07:00", value)
		} else {
			parsed, err = time.ParseInLocation("T15:04:05", value, loc)
		}
	} else if !strings.Contains(value, "T") { // value contains only date information
		parsed, err = time.ParseInLocation("2006-01-02", value, loc)
	} else {
		parsed, err = time.Parse(time.RFC3339, value)
	}
	return (Time)(parsed), err
}

// Format returns a textual representation of the time value formatted according to layout,
// which defines the format by showing how the reference time, defined to be
//
//    Mon Jan 2 15:04:05 -0700 MST 2006
//
// would be displayed if it were the value; it serves as an example of the desired output.
// The same display rules will then be applied to the time value.
//
// A fractional second is represented by adding a period and zeros to the end of the seconds section
// of layout string, as in "15:04:05.000" to format a time stamp with millisecond precision.
//
// Predefined layouts ANSIC, UnixDate, RFC3339 and others describe standard and convenient representations
// of the reference time. For more information about the formats and the definition of the reference time,
// see the documentation for ANSIC and the other constants defined by this package.
func (t Time) Format(layout string) string {
	return t.AsTime().Format(layout)
}

// String returns the time formatted using the format string
//
//    "2006-01-02 15:04:05.999999999 -0700 MST"
//
// If the time has a monotonic clock reading, the returned string includes a final field "m=±<value>",
// where value is the monotonic clock reading formatted as a decimal number of seconds.
//
// The returned string is meant for debugging; for a stable serialized representation,
// use t.MarshalText, t.MarshalBinary, or t.Format with an explicit format string.
func (t Time) String() string {
	return t.AsTime().String()
}

// MarshalJSON marshals this into JSON
//   implements json.Marshaler interface
//
//   We store time as RFC3339 UTC
func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).UTC().Format(time.RFC3339))
}

// UnmarshalJSON decodes JSON
//   implements json.Unmarshaler interface
//
//   We read time as RFC3339 UTC
func (t *Time) UnmarshalJSON(payload []byte) (err error) {
	var inner string
	if err = json.Unmarshal(payload, &inner); err != nil {
		return
	}
	if len(inner) == 0 {
		(*t) = Time(time.Time{})
		return
	}
	parsed, err := time.Parse(time.RFC3339, inner)
	if err != nil {
		return
	}
	(*t) = Time(parsed)
	return
}