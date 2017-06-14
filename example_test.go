package radix_test

import (
	"fmt"
	"github.com/yourbasic/radix"
)

func ExampleSortSlice() {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}
	radix.SortSlice(people, func(i int) string { return people[i].Name })
	fmt.Println(people)
	// Output: [{Alice 55} {Bob 75} {Gopher 7} {Vera 24}]
}
