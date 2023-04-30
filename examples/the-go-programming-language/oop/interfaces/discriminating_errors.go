package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("File path not passed")
	}

	filepath := os.Args[1]
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("File %s not found", filepath)
		}
		if os.IsPermission(err) {
			log.Fatalf("Permission denied to read file %s", filepath)
		}
		log.Fatal(err)
	}

	fmt.Printf("%s\n", fileContent)
}
