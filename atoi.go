package core

import (
	"strconv"
)

// Atoi convert a string to an int using a fallback if the conversion fails
func Atoi(value string, fallback int) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return result
}