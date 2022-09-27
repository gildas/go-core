package core

import (
	"errors"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// GetEnvAsString returns the string value of an environment variable by its name
//
// if not present, the fallback value is used
func GetEnvAsString(name, fallback string) string {
	if value, ok := os.LookupEnv(name); ok && len(value) > 0 {
		return value
	}
	return fallback
}

// GetEnvAsBool returns the bool value of an environment variable by its name
//
// if not present, the fallback value is used
func GetEnvAsBool(name string, fallback bool) bool {
	if value, ok := os.LookupEnv(name); ok && len(value) > 0 {
		return strings.Contains("1onyestrue", strings.ToLower(value))
	}
	return fallback
}

// GetEnvAsInt returns the int value of an environment variable by its name
//
// if not present, the fallback value is used
func GetEnvAsInt(name string, fallback int) int {
	if value, ok := os.LookupEnv(name); ok && len(value) > 0 {
		if intvalue, err := strconv.Atoi(value); err == nil {
			return intvalue
		}
	}
	return fallback
}

// GetEnvAsTime returns the time value of an environment variable by its name
//
// if not present, the fallback value is used
func GetEnvAsTime(name string, fallback time.Time) time.Time {
	if value, ok := os.LookupEnv(name); ok && len(value) > 0 {
		if timevalue, err := time.Parse(time.RFC3339, value); err == nil {
			return timevalue
		}
	}
	return fallback
}

// GetEnvAsDuration returns the time value of an environment variable by its name
//
// if not present, the fallback value is used
func GetEnvAsDuration(name string, fallback time.Duration) time.Duration {
	if value, ok := os.LookupEnv(name); ok && len(value) > 0 {
		if duration, err := ParseDuration(value); err == nil {
			return duration
		}
	}
	return fallback
}

// GetEnvAsURL returns the URL value of an environment variable by its name
//
// if not present, the fallback value is used
func GetEnvAsURL(name string, fallback interface{}) *url.URL {
	if value, ok := os.LookupEnv(name); ok && len(value) > 0 {
		if address, err := url.Parse(value); err == nil {
			return address
		}
	}
	if address, ok := fallback.(*url.URL); ok {
		return address
	}
	if address, ok := fallback.(url.URL); ok {
		return &address
	}
	if value, ok := fallback.(string); ok {
		if address, err := url.Parse(value); err == nil {
			return address
		}
	}
	panic(errors.New("Invalid fallback type in core.GetEnvAsURL"))
}

func GetEnvAsUUID(name string, fallback uuid.UUID) uuid.UUID {
	if value, ok := os.LookupEnv(name); ok && len(value) > 0 {
		if uuid, err := uuid.Parse(value); err == nil {
			return uuid
		}
	}
	return fallback
}
