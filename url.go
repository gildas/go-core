package core

import (
	"encoding/json"
	"net/url"
)

// URL is a placeholder so we can add new funcs to the type
type URL url.URL

// MarshalJSON marshals this into JSON
func (u URL) MarshalJSON() ([]byte, error) {
	uu := url.URL(u)
	return json.Marshal((&uu).String())
}

// UnmarshalJSON decodes JSON
func (u *URL) UnmarshalJSON(payload []byte) (err error) {
	var inner string
	if err = json.Unmarshal(payload, &inner); err != nil {
		return
	}
	if len(inner) == 0 {
		return nil
	}
	parsed, err := url.Parse(inner)
	if err != nil {
		return
	}
	(*u) = *(*URL)(parsed)
	return
}

func (u *URL) String() string {
	return (*url.URL)(u).String()
}
