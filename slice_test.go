package core_test

import (
	"testing"

	"github.com/gildas/go-core"
	"github.com/stretchr/testify/assert"
)

func TestSliceCanBeFiltered(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 4, 6, 8, 10}

	result := core.Filter(items, func(item int) bool {
		return item%2 == 0
	})
	assert.Equal(t, expected, result)
}

func TestSliceCanBeMapped(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}

	result := core.Map(items, func(item int) int {
		return item * item
	})
	assert.Equal(t, expected, result)
}

func TestSliceCanBeReduced(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := 55

	result := core.Reduce(items, 0, func(sum, item int) int {
		return sum + item
	})
	assert.Equal(t, expected, result)
}

func TestSliceCanBeSorted(t *testing.T) {
	items := []int{10, 2, 8, 3, 6, 5, 4, 7, 9, 1}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	core.Sort(items, func(a, b int) bool {
		return a < b
	})
	assert.Equal(t, expected, items)
}
