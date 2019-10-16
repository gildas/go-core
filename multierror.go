package core

import (
	"strings"
	"fmt"
)

// MultiError is used to collect errors during a loop, e.g.
type MultiError struct {
	Errors []error
}

func (me *MultiError) Error() string {
	text := strings.Builder{}
	for _, err := range me.Errors {
		text.WriteString(err.Error())
		text.WriteString("\n")
	}
	return fmt.Sprintf("%d Errors: %s", len(me.Errors), text.String())
}

// HasErrors tells if this contains errors
func (me MultiError) HasErrors() bool {
	return len(me.Errors) > 0
}

// Append appends a new error if any
func (me *MultiError) Append(err error) (*MultiError) {
	if err != nil {
		me.Errors = append(me.Errors, err)
	}
	return me
}

// AsError returns this if it contains errors, nil otherwise
func (me *MultiError) AsError() error {
	if me == nil || len(me.Errors) == 0 {
		return nil
	}
	return me
}

// GoString returns a GO representation of this
func (me MultiError) GoString() string {
	return fmt.Sprintf("%#v", me)
}