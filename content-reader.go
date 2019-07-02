package core

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/url"
)

// Content defines some content
type Content struct {
	Type   string        `json:"contentType"`
	URL    *url.URL      `json:"contentUrl"`
	Length int           `json:"contentLength"`
	Data   []byte        `json:"contentData"`
}

// ContentReader defines a content Reader (like data from an HTTP request)
type ContentReader struct {
	Type   string        `json:"contentType"`
	Length int           `json:"contentLength"`
	Reader io.ReadCloser `json:"-"`
}

// Reader gets a ContentReader from this Content
func (content *Content) Reader() *ContentReader {
	return &ContentReader{
		Type:   content.Type,
		Length: content.Length,
		Reader: ioutil.NopCloser(bytes.NewReader(content.Data)),
	}
}