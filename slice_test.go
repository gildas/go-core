package core_test

import (
	"testing"

	"github.com/gildas/go-core"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSliceCanFindValue(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result, found := core.Find(items, 5)
	assert.True(t, found)
	assert.Equal(t, 5, result)

	_, found = core.Find(items, 11)
	assert.False(t, found)
}

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

func TestSliceCanContains(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}

	assert.True(t, core.Contains(items, 1))
	assert.False(t, core.Contains(items, 6))
}

func TestSliceFindValueWithFunc(t *testing.T) {
	t.Run("Raw", func(t *testing.T) {
		item := Something1{"3"}
		items := []Something1{{"1"}, {"2"}, {"3"}, {"4"}, {"5"}}

		result, found := core.FindWithFunc(items, func(item Something1) bool {
			return item.Data == "3"
		})
		assert.True(t, found)
		assert.Equal(t, item, result)

		_, found = core.FindWithFunc(items, func(item Something1) bool {
			return item.Data == "6"
		})
		assert.False(t, found)
	})

	t.Run("Named", func(t *testing.T) {
		item := Something1{"3"}
		items := []Something1{{"1"}, {"2"}, {"3"}, {"4"}, {"5"}}

		result, found := core.FindWithFunc(items, core.MatchNamed(item))
		assert.True(t, found)
		assert.Equal(t, item, result)
	})

	t.Run("Identifiable", func(t *testing.T) {
		item := Something3{uuid.New()}
		items := []Something3{{uuid.New()}, {uuid.New()}, {uuid.New()}, item, {uuid.New()}}

		result, found := core.FindWithFunc(items, core.MatchIdentifiable(item))
		assert.True(t, found)
		assert.Equal(t, item, result)
	})

	t.Run("StringIdentifiable", func(t *testing.T) {
		item := Something4{ID: "3"}
		items := []Something4{{ID: "0"}, {ID: "1"}, {ID: "2"}, item, {ID: "4"}}

		result, found := core.FindWithFunc(items, core.MatchStringIdentifiable(item))
		assert.True(t, found)
		assert.Equal(t, item, result)
	})
}

func TestSliceCanContainsWithFunc(t *testing.T) {
	t.Run("Raw", func(t *testing.T) {
		items := []Something1{{"1"}, {"2"}, {"3"}, {"4"}, {"5"}}

		assert.True(t, core.ContainsWithFunc(items, Something1{"1"}, func(a, b Something1) bool { return a.Data == b.Data }))
		assert.True(t, core.ContainsWithFunc(items, Something1{"1"}, core.EqualNamed))
		assert.False(t, core.ContainsWithFunc(items, Something1{"6"}, func(a, b Something1) bool { return a.Data == b.Data }))
	})

	t.Run("Identifiable", func(t *testing.T) {
		item := Something3{uuid.New()}
		items := []Something3{{uuid.New()}, {uuid.New()}, {uuid.New()}, item, {uuid.New()}}

		assert.True(t, core.ContainsWithFunc(items, item, core.EqualIdentifiable))
		assert.False(t, core.ContainsWithFunc(items, Something3{uuid.New()}, core.EqualIdentifiable))
	})

	t.Run("StringIdentifiable", func(t *testing.T) {
		item := Something4{ID: "3"}
		items := []Something4{{ID: "0"}, {ID: "1"}, {ID: "2"}, item, {ID: "4"}}

		assert.True(t, core.ContainsWithFunc(items, item, core.EqualStringIdentifiable))
		assert.False(t, core.ContainsWithFunc(items, Something4{ID: "6"}, core.EqualStringIdentifiable))
	})
}

func TestSliceCanBeCompared(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}

	items := []int{1, 2, 3, 4, 5}
	assert.True(t, core.EqualSlices(items, expected))

	items = []int{1, 3, 2, 4, 5}
	assert.True(t, core.EqualSlices(items, expected))

	items = []int{1, 2, 3, 4, 5, 6}
	assert.False(t, core.EqualSlices(items, expected))

	items = []int{1, 2, 2, 4, 5}
	assert.False(t, core.EqualSlices(items, expected))

	items = []int{1, 2, 3, 4, 6}
	assert.False(t, core.EqualSlices(items, expected))
}

func TestSliceCanBeComparedWithFunc(t *testing.T) {
	expected := []Something{Something1{"1"}, Something1{"2"}, Something1{"3"}, Something1{"4"}, Something1{"5"}}

	items := []Something{Something1{"1"}, Something1{"2"}, Something1{"3"}, Something1{"4"}, Something1{"5"}}
	assert.True(t, core.EqualSlicesWithFunc(items, expected, func(a, b Something) bool { return a.GetData() == b.GetData() }))

	items = []Something{Something1{"1"}, Something1{"3"}, Something1{"2"}, Something1{"4"}, Something1{"5"}}
	assert.True(t, core.EqualSlicesWithFunc(items, expected, func(a, b Something) bool { return a.GetData() == b.GetData() }))

	items = []Something{Something1{"1"}, Something1{"2"}, Something1{"3"}, Something1{"4"}, Something1{"5"}, Something1{"6"}}
	assert.False(t, core.EqualSlicesWithFunc(items, expected, func(a, b Something) bool { return a.GetData() == b.GetData() }))

	items = []Something{Something1{"1"}, Something1{"2"}, Something1{"2"}, Something1{"4"}, Something1{"5"}}
	assert.False(t, core.EqualSlicesWithFunc(items, expected, func(a, b Something) bool { return a.GetData() == b.GetData() }))

	items = []Something{Something1{"1"}, Something1{"2"}, Something1{"3"}, Something1{"4"}, Something1{"6"}}
	assert.False(t, core.EqualSlicesWithFunc(items, expected, func(a, b Something) bool { return a.GetData() == b.GetData() }))
}

func TestSliceCanJoin(t *testing.T) {
	items := []Something2{{"1"}, {"2"}, {"3"}, {"4"}, {"5"}}
	assert.Equal(t, "1, 2, 3, 4, 5", core.Join(items, ", "))

	items = []Something2{{"1"}}
	assert.Equal(t, "1", core.Join(items, ", "))

	items = []Something2{}
	assert.Equal(t, "", core.Join(items, ", "))
}

func TestSliceCanJoinWithFunc(t *testing.T) {
	items := []Something1{{"1"}, {"2"}, {"3"}, {"4"}, {"5"}}
	assert.Equal(t, "1, 2, 3, 4, 5", core.JoinWithFunc(items, ", ", func(item Something1) string { return item.Data }))

	items = []Something1{{"1"}}
	assert.Equal(t, "1", core.JoinWithFunc(items, ", ", func(item Something1) string { return item.Data }))

	items = []Something1{}
	assert.Equal(t, "", core.JoinWithFunc(items, ", ", func(item Something1) string { return item.Data }))
}
