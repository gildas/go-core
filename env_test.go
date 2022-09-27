package core_test

import (
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/gildas/go-core"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCanGetEnvAsString(t *testing.T) {
	os.Setenv("TEST", "hello")
	defer func() {
		os.Unsetenv("TEST")
	}()

	value := core.GetEnvAsString("TEST", "world")
	assert.Equal(t, "hello", value)

	value = core.GetEnvAsString("NOT_HERE", "world")
	assert.Equal(t, "world", value)
}

func TestCanGetEnvAsBool(t *testing.T) {
	os.Setenv("TEST", "true")
	defer func() {
		os.Unsetenv("TEST")
	}()

	value := core.GetEnvAsBool("TEST", false)
	assert.Equal(t, true, value)

	os.Setenv("TEST", "TruE")
	value = core.GetEnvAsBool("TEST", false)
	assert.Equal(t, true, value)

	os.Setenv("TEST", "1")
	value = core.GetEnvAsBool("TEST", false)
	assert.Equal(t, true, value)

	os.Setenv("TEST", "on")
	value = core.GetEnvAsBool("TEST", false)
	assert.Equal(t, true, value)

	os.Setenv("TEST", "yes")
	value = core.GetEnvAsBool("TEST", false)
	assert.Equal(t, true, value)

	value = core.GetEnvAsBool("NOT_HERE", true)
	assert.Equal(t, true, value)
}

func TestCanGetEnvAsInt(t *testing.T) {
	os.Setenv("TEST", "12")
	defer func() {
		os.Unsetenv("TEST")
	}()

	value := core.GetEnvAsInt("TEST", 100)
	assert.Equal(t, 12, value)

	value = core.GetEnvAsInt("NOT_HERE", 100)
	assert.Equal(t, 100, value)
}

func TestCanGetEnvAsTime(t *testing.T) {
	now := time.Now()
	os.Setenv("TEST", now.Format(time.RFC3339))
	defer func() {
		os.Unsetenv("TEST")
	}()

	value := core.GetEnvAsTime("TEST", now.Add(24*time.Hour))
	assert.Equal(t, now.Format(time.RFC3339), value.Format(time.RFC3339))

	value = core.GetEnvAsTime("NOT_HERE", now)
	assert.Equal(t, now.Format(time.RFC3339), value.Format(time.RFC3339))
}

func TestCanGetEnvAsDuration(t *testing.T) {
	os.Setenv("TEST", "12s")
	defer func() {
		os.Unsetenv("TEST")
	}()

	value := core.GetEnvAsDuration("TEST", 100*time.Second)
	assert.Equal(t, 12*time.Second, value)

	value = core.GetEnvAsDuration("NOT_HERE", 100*time.Second)
	assert.Equal(t, 100*time.Second, value)
}

func TestCanGetEnvAsURL(t *testing.T) {
	testURL, _ := url.Parse("https://www.acme.com")
	os.Setenv("TEST", testURL.String())
	defer func() {
		os.Unsetenv("TEST")
	}()

	value := core.GetEnvAsURL("TEST", "http://localhost")
	assert.Equal(t, testURL, value)

	value = core.GetEnvAsURL("NOT_HERE", testURL)
	assert.Equal(t, testURL, value)

	value = core.GetEnvAsURL("NOT_HERE", *testURL)
	assert.Equal(t, testURL, value)

	value = core.GetEnvAsURL("NOT_HERE", testURL.String())
	assert.Equal(t, testURL, value)
}

func TestPanicGetEnvAsURLWithInvalidFallback(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("code did not panic")
		}
	}()
	_ = core.GetEnvAsURL("NOT_HERE", 1234)
}

func TestCanGetEnvAsUUID(t *testing.T) {
	expected := uuid.New()
	os.Setenv("TEST", expected.String())
	os.Setenv("WRONG", "not a UUID")
	defer func() {
		os.Unsetenv("TEST")
		os.Unsetenv("WRONG")
	}()

	value := core.GetEnvAsUUID("TEST", uuid.New())
	assert.Equal(t, expected, value)

	value = core.GetEnvAsUUID("NOT_HERE", expected)
	assert.Equal(t, expected, value)

	value = core.GetEnvAsUUID("WRONG", expected)
	assert.Equal(t, expected, value)
}
