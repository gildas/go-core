package core

import (
	"io"
)

// ContentReader defines a content Reader (like data from an HTTP request)
type ContentReader struct {
	Type   string        `json:"contentType"`
	Length int           `json:"contentLength"`
	Reader io.ReadCloser `json:"-"`
}