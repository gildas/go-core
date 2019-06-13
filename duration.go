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
// We store duration in milliseconds
func (duration Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(int64(time.Duration(duration)) / int64(1000000))
}

// UnmarshalJSON decodes JSON
// an int (int64) will be assumed to be milli-seconds
// a string will be parsed as an ISO 8601 then as a GO time.Duration
func (duration *Duration) UnmarshalJSON(payload []byte) (err error) {
	var inner interface{}
	if err := json.Unmarshal(payload, &inner); err != nil {
		return err
	}
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
func ParseDuration(iso8601 string) (duration time.Duration, err error) {
	parser := regexp.MustCompile(`P(?P<years>\d+Y)?(?P<months>\d+M)?(?P<weeks>\d+W)?(?P<days>\d+D)?T?(?P<hours>\d+H)?(?P<minutes>\d+M)?(?P<seconds>\d+S)?`)
	matches := parser.FindStringSubmatch(iso8601)
	var parsed int

	if len(matches) == 0 {
		return time.Duration(0), fmt.Errorf(`"%s" is not an ISO8601 duration`, iso8601)
	}
	if len(matches[1]) > 0 {
		if parsed, err = strconv.Atoi(matches[1][:len(matches[1])-1]); err != nil {
			return
		}
		duration = time.Duration(parsed*24*365) * time.Hour // years
	}

	if len(matches[2]) > 0 {
		if parsed, err = strconv.Atoi(matches[2][:len(matches[2])-1]); err != nil {
			return
		}
		duration += time.Duration(parsed*24*30) * time.Hour // months
	}

	if len(matches[3]) > 0 {
		if parsed, err = strconv.Atoi(matches[3][:len(matches[3])-1]); err != nil {
			return
		}
		duration += time.Duration(parsed*24*7) * time.Hour // weeks
	}

	if len(matches[4]) > 0 {
		if parsed, err = strconv.Atoi(matches[4][:len(matches[4])-1]); err != nil {
			return
		}
		duration += time.Duration(parsed*24) * time.Hour // days
	}

	if len(matches[5]) > 0 {
		if parsed, err = strconv.Atoi(matches[5][:len(matches[5])-1]); err != nil {
			return
		}
		duration += time.Duration(parsed) * time.Hour // hours
	}

	if len(matches[6]) > 0 {
		if parsed, err = strconv.Atoi(matches[6][:len(matches[6])-1]); err != nil {
			return
		}
		duration += time.Duration(parsed) * time.Minute // minutes
	}

	if len(matches[7]) > 0 {
		if parsed, err = strconv.Atoi(matches[7][:len(matches[7])-1]); err != nil {
			return
		}
		duration += time.Duration(parsed) * time.Second // seconds
	}

	return
}

func (duration Duration) String() string {
	return time.Duration(duration).String()
}
