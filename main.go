package main

import (
	"fmt"
	"log"

	"github.com/calebryant/gostash-antlr/parsing"
)

func main() {
	filePath := "examples/parser.conf"

	fmt.Printf("Parsing file: %s\n\n", filePath)

	err := parsing.ParseAndPrint(filePath)
	if err != nil {
		log.Fatalf("Error parsing file: %v", err)
	}
}
