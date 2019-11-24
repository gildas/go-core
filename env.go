package core

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// GetEnvAsString returns the string value of an environment variable by its name
// or of the fallback if it is not present
func GetEnvAsString(name, fallback string) string {
	if value, ok := os.LookupEnv(name); ok && len(value) > 0 {
		return value
	}
	return fallback
}

// GetEnvAsBool returns the bool value of an environment variable by its name
// or of the fallback if it is not present
func GetEnvAsBool(name string, fallback bool) bool {
	if value, ok := os.LookupEnv(name); ok && len(value) > 0 {
		if strings.Contains(strings.ToLower(value), "1onyestrue") {
			return true
		}
		return false
	}
	return fallback
}

// GetEnvAsInt returns the int value of an environment variable by its name
// or of the fallback if it is not present
func GetEnvAsInt(name string, fallback int) int {
	if value, ok := os.LookupEnv(name); ok {
		if intvalue, err := strconv.Atoi(value); err == nil {
			return intvalue
		}
	}
	return fallback
}

// GetEnvAsTime returns the time value of an environment variable by its name
// or of the fallback if it is not present
func GetEnvAsTime(name string, fallback time.Time) time.Time {
	if value, ok := os.LookupEnv(name); ok {
		if timevalue, err := time.Parse(time.RFC3339, value); err == nil {
			return timevalue
		}
	}
	return fallback
}

// GetEnvAsDuration returns the time value of an environment variable by its name
// or of the fallback if it is not present
func GetEnvAsDuration(name string, fallback time.Duration) time.Duration {
	if value, ok := os.LookupEnv(name); ok {
		if duration, err := ParseDuration(value); err == nil {
			return duration
		}
	}
	return fallback
}