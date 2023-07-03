package core

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Duration is a placeholder so we can add new funcs to the type
type Duration time.Duration

// MarshalJSON marshals this into JSON
//
// # The duration is serialized in milliseconds
//
// implements json.Marshaler interface
func (duration Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(int64(time.Duration(duration)) / int64(1000000))
}

// UnmarshalJSON decodes JSON
//
//	implements json.Unmarshaler interface
//
//	an int (int64) will be assumed to be milli-seconds
//	a string will be parsed as an ISO 8601 then as a GO time.Duration
func (duration *Duration) UnmarshalJSON(payload []byte) (err error) {
	var inner interface{}
	_ = json.Unmarshal(payload, &inner)
	switch value := inner.(type) {
	case float64:
		(*duration) = Duration(value * float64(1000000))
	case string:
		var d time.Duration
		if strings.HasPrefix(strings.TrimSpace(value), "P") {
			if d, err = ParseDuration(value); err != nil {
				return
			}
			(*duration) = Duration(d)
			return
		}
		if d, err = time.ParseDuration(value); err != nil {
			return err
		}
		(*duration) = Duration(d)
	default:
		return fmt.Errorf("Invalid Duration")
	}
	return
}

// ParseDuration parses an ISO8601 duration
//
//	If the given value is not an ISO8601 duration, returns time.ParseDuration
func ParseDuration(value string) (duration time.Duration, err error) {
	parser := regexp.MustCompile(`P(?P<years>\d+Y)?(?P<months>\d+M)?(?P<weeks>\d+W)?(?P<days>\d+D)?T?(?P<hours>\d+H)?(?P<minutes>\d+M)?(?P<seconds>\d+(?:\.\d+)?S)?`)
	matches := parser.FindStringSubmatch(value)
	var parsed int

	if len(matches) == 0 {
		return time.ParseDuration(value)
	}

	ok := false
	for _, match := range matches[1:] {
		if len(match) > 0 {
			ok = true
		}
	}
	if !ok {
		return time.Duration(0), fmt.Errorf(`"%s" is not an ISO8601 duration`, value)
	}
	if len(matches[1]) > 0 {
		parsed, _ = strconv.Atoi(matches[1][:len(matches[1])-1]) // remove the Y
		duration = time.Duration(parsed*24*365) * time.Hour      // years
	}
	if len(matches[2]) > 0 {
		parsed, _ = strconv.Atoi(matches[2][:len(matches[2])-1]) // remove the M
		duration += time.Duration(parsed*24*30) * time.Hour      // months
	}
	if len(matches[3]) > 0 {
		parsed, _ = strconv.Atoi(matches[3][:len(matches[3])-1]) // remove the W
		duration += time.Duration(parsed*24*7) * time.Hour       // weeks
	}
	if len(matches[4]) > 0 {
		parsed, _ = strconv.Atoi(matches[4][:len(matches[4])-1]) // remove the D
		duration += time.Duration(parsed*24) * time.Hour         // days
	}
	if len(matches[5]) > 0 {
		parsed, _ = strconv.Atoi(matches[5][:len(matches[5])-1]) // remove the H
		duration += time.Duration(parsed) * time.Hour            // hours
	}
	if len(matches[6]) > 0 {
		parsed, _ = strconv.Atoi(matches[6][:len(matches[6])-1]) // remove the M
		duration += time.Duration(parsed) * time.Minute          // minutes
	}
	if len(matches[7]) > 0 {
		fraction, _ := strconv.ParseFloat(matches[7][:len(matches[7])-1], 64) // remove the S
		duration += time.Duration(fraction * float64(time.Second))
	}
	return
}

func (duration Duration) String() string {
	return time.Duration(duration).String()
}
