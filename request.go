package core

import (
	"context"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gildas/go-logger"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// RequestOptions defines options of an HTTP request
type RequestOptions struct {
	Method         string
	URL            *url.URL
	Proxy          *url.URL
	Headers        map[string]string
	Parameters     map[string]string
	Accept         string
	Payload        interface{}
	Content        ContentReader
	Authentication string
	RequestID      string
	UserAgent      string
	Logger         *logger.Logger
}

// RequestError is returned when an HTTP Status is >= 400
type RequestError struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
}

// SendRequest sends an HTTP request
func SendRequest(ctx context.Context, options *RequestOptions, results interface{}) (*ContentReader, error) {
	if options.URL == nil {
		return nil, errors.WithStack(fmt.Errorf("error.url.empty"))
	}

	log := logger.CreateWithSink(nil)
	if options.Logger != nil {
		log = options.Logger.Scope("request")
	}
	if len(options.RequestID) == 0 {
		options.RequestID = uuid.Must(uuid.NewRandom()).String()
	}
	log = log.Record("reqid", options.RequestID)

	reqContent, err := buildRequestContent(log, options)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(options.Method) == 0 {
		if reqContent.Length > 0 {
			options.Method = "POST"
		} else {
			options.Method = "GET"
		}
	}

	req, err := http.NewRequest(options.Method, options.URL.String(), reqContent.Reader)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Setting request headers
	req.Header.Set("User-Agent",   options.UserAgent)
	req.Header.Set("Accept",       options.Accept)
	req.Header.Set("X-Request-Id", options.RequestID)
	if len(reqContent.Type) > 0 {
		req.Header.Set("Content-Type", reqContent.Type)
	}
	if reqContent.Length > 0 {
		req.Header.Set("Content-Length", strconv.FormatInt(reqContent.Length, 10))
	}
	for key, value := range options.Headers {
		req.Header.Set(key, value)
	}

	// Sending the request...
	log.Debugf("HTTP %s %s", req.Method, req.URL.String())
	log.Tracef("Request Headers: %#v", req.Header)
	httpclient := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	if options.Proxy != nil {
		httpclient.Transport = &http.Transport{Proxy: http.ProxyURL(options.Proxy)}
	}
	start    := time.Now()
	res, err := httpclient.Do(req)
	duration := time.Since(start)
	log      = log.Record("duration", duration)
	if err != nil {
		log.Errorf("Failed to send request", err)
		return nil, errors.WithStack(err)
	}
	defer res.Body.Close()
	log.Debugf("Response %s in %s", res.Status, duration)
	log.Tracef("Response Headers: %#v", res.Header)

	// Reading the response body
	resContent, err := ContentFromReader(res.Body, res.Header.Get("Content-Type"), Atoi(res.Header.Get("Content-Length"), 0))
	if err != nil {
		log.Errorf("Failed to read response body", err)
		return nil, errors.WithStack(err)
	}
	// some servers give the wrong mime type for JPEG files
	if resContent.Type == "image/jpg" {
		resContent.Type = "image/jpeg"
	}
	if len(resContent.Type) == 0 || resContent.Type == "application/octet-stream" {
		if len(options.Accept) > 0 {
			// TODO: well... Accept is not always a simple mime type...
			resContent.Type = options.Accept
		}
		if resContent.Type == "application/octet-stream" {
			if restype := mime.TypeByExtension(filepath.Ext(options.URL.Path)); len(restype) > 0 {
				resContent.Type = restype
			}
		}
	}
	log.Tracef("Response body (%s, %d bytes): %s", resContent.Type, resContent.Length, string(resContent.Data[:int(math.Min(1024,float64(resContent.Length)))]))

	// Processing the status
	if res.StatusCode == http.StatusFound {
		follow, err := res.Location()
		if err == nil {
			log.Warnf("TODO: we should get stuff from %s", follow.String())
		}
	}
	if res.StatusCode >= 400 {
		return resContent.Reader(), errors.WithStack(RequestError{res.StatusCode, res.Status})
	}

	// Unmarshaling the content if requested
	if results != nil {
		err = json.Unmarshal(resContent.Data, results)
		if err != nil {
			log.Errorf("Failed to decode response", err)
			return resContent.Reader(), errors.WithStack(err)
		}
	}
	return resContent.Reader(), nil
}

func (err RequestError) Error() string {
	return err.Status
}

// buildRequestContent builds a Content for the request
func buildRequestContent(log *logger.Logger, options *RequestOptions) (*ContentReader, error) {
	content := &Content{}
	if len(options.Parameters) > 0 {
		if options.Content.Type == "application/x-www-form-urlencoded" {
			query := url.Values{}
			for param, value := range options.Parameters {
				query.Set(param, value)
			}
			content = ContentWithData([]byte(query.Encode()), options.Content.Type)
		} else { // Create a multipart data form
			body   := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			for param, value := range options.Parameters {
				if strings.HasPrefix(param, ">") {
					param = strings.TrimPrefix(param, ">")
					if len(value) == 0 {
						return nil, errors.WithStack(fmt.Errorf("Empty value for field %s", param))
					}
					part, err := writer.CreateFormFile(param, value)
					if err != nil {
						return nil, errors.Wrap(err, fmt.Sprintf("Failed to create multipart for field %s", param))
					}
					if options.Content.Length == 0 {
						return nil, errors.WithStack(fmt.Errorf("Missing/Empty Content for Parameter %s", param))
					}
					written, err := io.Copy(part, options.Content)
					if err != nil {
						return nil, errors.Wrap(err, fmt.Sprintf("Failed to write payload to multipart field %s", param))
					}
					log.Tracef("Wrote %d bytes to field %s", written, param)
				} else {
					if err := writer.WriteField(param, value); err != nil {
						return nil, errors.Wrap(err, fmt.Sprintf("Failed to create field %s", param))
					}
					log.Tracef("Added field %s = %s", param, value)
				}
			}
			if err := writer.Close(); err != nil {
				return nil, errors.Wrap(err, "Failed to create multipart data")
			}
			content, _ = ContentFromReader(body, writer.FormDataContentType())
		}
	} else if options.Payload != nil {
		// Create a JSON payload
		// TODO: Add other payload types like XML, etc
		payload, err := json.Marshal(options.Payload)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to encode payload into JSON")
		}
		content = ContentWithData(payload, options.Content.Type)
		if len(content.Type) == 0 {
			content.Type = "application/json"
		}
	} else if options.Content.Length > 0 {
		content, _ = ContentFromReader(options.Content)
		content.Type = options.Content.Type
		if len(content.Type) == 0 {
			content.Type = "application/octet-stream"
		}
	}
	return content.Reader(), nil
} 