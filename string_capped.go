package core

// CappedString returns a string capped to the given length.
//
// If the string is longer than the given length, it is truncated and "..." is appended.
//
// Note: the original string is not modified, a truncated copy is returned.
//
// Example:
//
//	// Cap a string to 10 characters
//	str := "This is a long string"
//	cappedStr := CappedString(str, 10) // "This is a ..."
func CappedString(s string, length int) string {
	return CappedStringWith(s, length, "...")
}

// CappedStringWith returns a string capped to the given length with a custom suffix.
//
// If the string is longer than the given length, it is truncated and the suffix is appended.
//
// Note: the original string is not modified, a truncated copy is returned.
//
// Example:
//
//	// Cap a string to 10 characters with " (more)" suffix
//	str := "This is a long string"
//	cappedStr := CappedStringWith(str, 10, " (more)") // "Thi (more)"
func CappedStringWith(s string, length int, suffix string) string {
	if len(s) <= length {
		return s
	}
	suffixLength := len(suffix)
	if length <= suffixLength {
		return s[:length]
	}
	return s[:length-suffixLength] + suffix
}
