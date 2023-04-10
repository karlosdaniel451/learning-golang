package main

import (
	"fmt"
	"os"
	"strings"
)

func variadicJoin(sep string, elems ...string) string {
	// Inefficient version:
	// resultString := ""

	// for _, elem := range elems[:len(elems)-1] {
	// 	resultString += elem + sep
	// }

	// resultString += elems[len(elems)-1]
	// return resultString

	var resultString strings.Builder

	for _, elem := range elems[:len(elems)-1] {
		resultString.WriteString(elem + sep)
	}

	resultString.WriteString(elems[len(elems)-1])
	return resultString.String()
}

func main() {
	fmt.Println(strings.Join(os.Args[1:], ", "))
	fmt.Println(variadicJoin(", ", os.Args[1:]...))
}
