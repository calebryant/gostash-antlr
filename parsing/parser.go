package parsing

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/calebryant/gostash-antlr/antlr"
)

// CustomErrorListener implements antlr.ErrorListener
type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	errors []string
}

func NewCustomErrorListener() *CustomErrorListener {
	return &CustomErrorListener{
		errors: []string{},
	}
}

func (l *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errors = append(l.errors, fmt.Sprintf("Line %d:%d: %s", line, column, msg))
}

func (l *CustomErrorListener) HasErrors() bool {
	return len(l.errors) > 0
}

func (l *CustomErrorListener) GetErrors() []string {
	return l.errors
}

// ParseFile reads and parses a Logstash configuration file
func ParseFile(filePath string) (parser.IFilterblockContext, error) {
	// Read the file
	input, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return ParseString(string(input))
}

// ParseString parses a Logstash configuration string
func ParseString(config string) (parser.IFilterblockContext, error) {
	// Create input stream
	inputStream := antlr.NewInputStream(config)

	// Create lexer
	lexer := parser.NewChronicleLogstashLexer(inputStream)

	// Create token stream
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create parser
	p := parser.NewChronicleLogstashParser(tokenStream)

	// Add custom error listener
	errorListener := NewCustomErrorListener()
	p.RemoveErrorListeners() // Remove default console listener
	p.AddErrorListener(errorListener)

	// Parse
	filterBlock := p.Filterblock()

	// Check for errors
	if errorListener.HasErrors() {
		for _, err := range errorListener.GetErrors() {
			fmt.Printf("Parse error: %s\n", err)
		}
	}

	return filterBlock, nil
}

// ParseAndPrint parses a file and prints the parse tree
func ParseAndPrint(filePath string) error {
	// Parse the file
	filterBlock, err := ParseFile(filePath)
	if err != nil {
		return err
	}

	if filterBlock == nil {
		return fmt.Errorf("parse tree is nil")
	}

	// Print the parse tree
	fmt.Printf("Successfully parsed: %s\n", filePath)
	fmt.Printf("Parse tree root rule: filterblock\n")
	fmt.Printf("Number of children in parse tree: %d\n", filterBlock.GetChildCount())

	return nil
}
