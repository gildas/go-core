package core_test

import (
	"net/url"
	"testing"

	"github.com/gildas/go-core"
	"github.com/stretchr/testify/assert"
)

func TestMustShouldNotPanicWithValidArgument(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Must panicked with %v", r)
		}
	}()
	value := core.Must(url.Parse("https://www.acme.com"))
	assert.NotNil(t, value)
}

func TestMustShouldPanicWithBadArgument(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Must() should have panicked")
		}
	}()
	value := core.Must(url.Parse("%%NOT_A_URL!!!"))
	t.Logf("Value is %v", value)
}
