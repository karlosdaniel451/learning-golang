// package main

// import (
// 	"flag"
// 	"fmt"
// 	"strings"
// )

// var newline = flag.Bool("n", false, "Omit trailing newline")
// var separator = flag.String("s", " ", "Separator")

// func main() {
// 	flag.Parse()

// 	fmt.Print(strings.Join(flag.Args(), *separator))
// 	if *newline {
// 		fmt.Println()
// 	}
// }
