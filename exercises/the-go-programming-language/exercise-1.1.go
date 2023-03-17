/*
	Exercis e 1.1: Mo dif y the echo prog ram to als o pr int os.Args[0], the

name of the command that invo ked it.
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
