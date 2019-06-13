package core

import (
	"encoding/json"
	"time"
)

// Time is a placeholder so we can add new funcs to the type
type Time time.Time

// MarshalJSON marshals this into JSON
// We store time as RFC3339 UTC
func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).UTC().Format(time.RFC3339))
}

// UnmarshalJSON decodes JSON
// We read time as RFC3339 UTC
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