package main

import (
	"fmt"
	"log"

	"github.com/calebryant/gostash-antlr/parsing"
)

func main() {
	filePath := "examples/parser.conf"

	fmt.Printf("Parsing file: %s\n\n", filePath)

	// Parse and get AST
	ast, err := parsing.ParseToAST(filePath)
	if err != nil {
		log.Fatalf("Error parsing file: %v", err)
	}

	// Convert to JSON
	jsonStr, err := ast.ToJSON()
	if err != nil {
		log.Fatalf("Error converting to JSON: %v", err)
	}

	// Print JSON output
	fmt.Println(jsonStr)
}
