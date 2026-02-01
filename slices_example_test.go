package core_test

import (
	"fmt"

	"github.com/gildas/go-core"
)

func ExampleContains() {
	items := []string{"apple", "banana", "cherry"}

	found := core.Contains(items, "banana")
	fmt.Println("Found banana:", found)

	notFound := core.Contains(items, "date")
	fmt.Println("Found date:", notFound)

	// Output:
	// Found banana: true
	// Found date: false
}

func ExampleContainsWithFunc() {
	items := []Something1{{"apple"}, {"banana"}, {"cherry"}}

	found := core.ContainsWithFunc(items, Something1{"banana"}, func(a, b Something1) bool {
		return a.Data == b.Data
	})
	fmt.Println("Found banana:", found)

	notFound := core.ContainsWithFunc(items, Something1{"date"}, func(a, b Something1) bool {
		return a.Data == b.Data
	})
	fmt.Println("Found date:", notFound)

	// Output:
	// Found banana: true
	// Found date: false
}

func ExampleConvertFromAnySlice() {
	anySlice := []any{"one", "two", "three"}
	stringSlice := core.ConvertFromAnySlice[string](anySlice)
	fmt.Println(stringSlice)

	anySlice = []any{"one", 2, "three", 4.0, "five"}
	stringSlice = core.ConvertFromAnySlice[string](anySlice)
	fmt.Println(stringSlice)

	// Output:
	// [one two three]
	// [one three five]
}

func ExampleConvertToAnySlice() {
	stringSlice := []string{"one", "two", "three"}
	anySlice := core.ConvertToAnySlice(stringSlice)
	fmt.Println(anySlice)

	intSlice := []int{1, 2, 3, 4, 5}
	anySlice = core.ConvertToAnySlice(intSlice)
	fmt.Println(anySlice)

	// Output:
	// [one two three]
	// [1 2 3 4 5]
}
func ExampleEqualSlices() {
	slice1 := []string{"apple", "banana", "cherry"}
	slice2 := []string{"cherry", "banana", "apple"}

	equal := core.EqualSlices(slice1, slice2)
	fmt.Println("Slices are equal:", equal)

	slice3 := []string{"apple", "banana", "date"}

	notEqual := core.EqualSlices(slice1, slice3)
	fmt.Println("Slices are equal:", notEqual)

	// Output:
	// Slices are equal: true
	// Slices are equal: false
}

func ExampleEqualSlicesWithFunc() {
	slice1 := []Something{Something1{"apple"}, Something1{"banana"}, Something1{"cherry"}}
	slice2 := []Something{Something1{"cherry"}, Something1{"banana"}, Something1{"apple"}}

	equal := core.EqualSlicesWithFunc(slice1, slice2, func(a, b Something) bool {
		return a.GetData() == b.GetData()
	})
	fmt.Println("Slices are equal:", equal)

	slice3 := []Something{Something1{"apple"}, Something1{"banana"}, Something1{"date"}}

	notEqual := core.EqualSlicesWithFunc(slice1, slice3, func(a, b Something) bool {
		return a.GetData() == b.GetData()
	})
	fmt.Println("Slices are equal:", notEqual)

	// Output:
	// Slices are equal: true
	// Slices are equal: false
}

func ExampleFind() {
	items := []string{"apple", "banana", "cherry"}

	result, found := core.Find(items, "banana")
	fmt.Println("Found banana:", found)
	if found {
		fmt.Println("Banana is:", result)
	}

	_, found = core.Find(items, "date")
	fmt.Println("Found date:", found)

	// Output:
	// Found banana: true
	// Banana is: banana
	// Found date: false
}

func ExampleFindWithFunc() {
	items := []Something1{{"apple"}, {"banana"}, {"cherry"}}

	result, found := core.FindWithFunc(items, func(item Something1) bool {
		return item.Data == "banana"
	})
	fmt.Println("Found banana:", found)
	if found {
		fmt.Println("Banana is:", result.Data)
	}

	_, found = core.FindWithFunc(items, func(item Something1) bool {
		return item.Data == "date"
	})
	fmt.Println("Found date:", found)

	// Output:
	// Found banana: true
	// Banana is: banana
	// Found date: false
}

func ExampleFilter() {
	// Filter all even numbers in a slice
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := core.Filter(numbers, func(number int) bool {
		return number%2 == 0
	})
	fmt.Println(evens)

	// Output:
	// [2 4 6 8 10]
}

func ExampleJoin() {
	items := []Something2{{"apple"}, {"banana"}, {"cherry"}}
	result := core.Join(items, "; ")
	fmt.Println(result)

	// Output:
	// apple; banana; cherry
}

func ExampleJoinWithFunc() {
	items := []Something1{{"apple"}, {"banana"}, {"cherry"}}
	result := core.JoinWithFunc(items, "; ", func(item Something1) string {
		return item.Data
	})
	fmt.Println(result)

	// Output:
	// apple; banana; cherry
}

func ExampleMap() {
	// Map all Something1 in a slice to their string representation
	slice := []Something1{{"apple"}, {"banana"}, {"cherry"}}
	strings := core.Map(slice, func(item Something1) string {
		return item.Data
	})
	fmt.Println(strings)

	// Output:
	// [apple banana cherry]
}

func ExampleReduce() {
	// Sum all numbers in a slice
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := core.Reduce(numbers, 0, func(sum, number int) int {
		return sum + number
	})
	fmt.Println(sum)

	// Output:
	// 55
}

func ExampleSort() {
	// Sort a slice of numbers
	numbers := []int{10, 2, 8, 3, 6, 5, 5, 4, 7, 9, 1}
	core.Sort(numbers, func(a, b int) bool {
		return a < b
	})
	fmt.Println(numbers)

	fruits := []string{"banana", "apple", "cherry"}
	core.Sort(fruits, func(a, b string) bool {
		return a < b
	})
	fmt.Println(fruits)

	things := []Something1{{"banana"}, {"apple"}, {"cherry"}}
	core.Sort(things, func(a, b Something1) bool {
		return a.Data < b.Data
	})
	strings := core.Map(things, func(item Something1) string {
		return item.Data
	})
	fmt.Println(strings)

	// Output:
	// [1 2 3 4 5 5 6 7 8 9 10]
	// [apple banana cherry]
	// [apple banana cherry]
}
