package core_test

import (
	"testing"

	"github.com/gildas/go-core"
	"github.com/stretchr/testify/assert"
)

func TestMapCanJoin(t *testing.T) {
	expected := map[string]int{"1": 1, "2": 4, "3": 6, "4": 8, "5": 10}
	map1 := map[string]int{"1": 1, "2": 2, "3": 3}
	map2 := map[string]int{"2": 4, "3": 6}
	map3 := map[string]int{"4": 8, "5": 10}
	assert.Equal(t, expected, core.MapJoin(map1, map2, map3))
}
