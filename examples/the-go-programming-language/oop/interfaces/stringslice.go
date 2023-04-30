package main

import (
	"fmt"
	"sort"
)

type StringSlice struct {
	Strings []string
}

func (stringSlice StringSlice) Len() int {
	return len(stringSlice.Strings)
}

func (stringSlice StringSlice) Less(i, j int) bool {
	return stringSlice.Strings[i] < stringSlice.Strings[j]
}

func (stringSlice StringSlice) Swap(i, j int) {
	stringSlice.Strings[i], stringSlice.Strings[j] = stringSlice.Strings[j], stringSlice.Strings[i]
}

func (stringSlice StringSlice) String() string {
	return fmt.Sprintf("%s", stringSlice.Strings)
}

func main() {
	// Assert at compile time that a value of `StringSlice` satisfies `sort.Interface`.
	var _ sort.Interface = StringSlice{}

	// Assert at compile time that a value of `StringSlice` satisfies `fmt.Stringer`.
	var _ fmt.Stringer = StringSlice{}

	stringSlice := StringSlice{Strings: []string{"hello", "aa", "za", "a", "b", "k", "world"}}

	sort.Sort(stringSlice)
	fmt.Println(stringSlice)
}
