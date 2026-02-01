package core_test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gildas/go-core"
)

func ExampleDuration_String() {
	duration := core.Duration(90 * time.Second)
	fmt.Println(duration.String())
	// Output: 1m30s
}

func ExampleDuration_ToISO8601() {
	duration := core.Duration(4000 * time.Hour)
	fmt.Println(duration.ToISO8601())
	// Output: P5M16DT16H
}

func ExampleDuration_MarshalJSON() {
	duration := core.Duration(90 * time.Second)
	payload, _ := json.Marshal(duration)
	fmt.Println(string(payload))
	// Output: 90000
}

func ExampleDuration_UnmarshalJSON_fromString() {
	payload := `{"duration":"1h15m30s"}`
	result := struct {
		Duration core.Duration `json:"duration"`
	}{}
	_ = json.Unmarshal([]byte(payload), &result)
	fmt.Println(time.Duration(result.Duration).String())
	// Output: 1h15m30s
}

func ExampleDuration_UnmarshalJSON_fromInt() {
	payload := `{"duration":4500000}`
	result := struct {
		Duration core.Duration `json:"duration"`
	}{}
	_ = json.Unmarshal([]byte(payload), &result)
	fmt.Println(time.Duration(result.Duration).String())
	// Output: 1h15m0s
}

func ExampleDuration_UnmarshalJSON_fromFloat() {
	payload := `{"duration":4500000.0}`
	result := struct {
		Duration core.Duration `json:"duration"`
	}{}
	_ = json.Unmarshal([]byte(payload), &result)
	fmt.Println(time.Duration(result.Duration).String())
	// Output: 1h15m0s
}

func ExampleDuration_UnmarshalJSON_fromISO() {
	payload := `{"duration":"P2DT3H4M5.006S"}`
	result := struct {
		Duration core.Duration `json:"duration"`
	}{}
	_ = json.Unmarshal([]byte(payload), &result)
	fmt.Println(time.Duration(result.Duration).String())
	// Output: 51h4m5.006s
}
