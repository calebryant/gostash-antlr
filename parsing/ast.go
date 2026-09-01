package parsing

import (
	"encoding/json"
)

// AST represents the Abstract Syntax Tree of a Logstash filter configuration
type AST struct {
	FilterName string                 `json:"filter_name"`
	Blocks     []ASTNode              `json:"blocks"`
	Source     map[string]interface{} `json:"source,omitempty"`
}

// ASTNode represents a node in the AST (plugin, conditional, etc.)
type ASTNode interface {
	NodeType() string
}

// Plugin represents a Logstash plugin (mutate, grok, etc.)
type Plugin struct {
	Name       string                 `json:"name"`
	Type       string                 `json:"type"` // "plugin"
	Config     map[string]interface{} `json:"config"`
	LineNumber int                    `json:"line_number,omitempty"`
}

func (p *Plugin) NodeType() string {
	return "plugin"
}

// Conditional represents an if/else if/else block
type Conditional struct {
	Type       string           `json:"type"` // "conditional"
	Condition  string           `json:"condition"`
	Blocks     []ASTNode        `json:"blocks"`
	Else       *ConditionalElse `json:"else,omitempty"`
	LineNumber int              `json:"line_number,omitempty"`
}

func (c *Conditional) NodeType() string {
	return "conditional"
}

// ConditionalElse represents an else or else if block
type ConditionalElse struct {
	Type       string           `json:"type"`                // "else" or "else if"
	Condition  string           `json:"condition,omitempty"` // only for else if
	Blocks     []ASTNode        `json:"blocks"`
	Else       *ConditionalElse `json:"else,omitempty"`
	LineNumber int              `json:"line_number,omitempty"`
}

// ForLoop represents a for loop
type ForLoop struct {
	Type       string    `json:"type"` // "for"
	Variable   string    `json:"variable"`
	Index      string    `json:"index,omitempty"`
	Iterable   string    `json:"iterable"`
	Blocks     []ASTNode `json:"blocks"`
	LineNumber int       `json:"line_number,omitempty"`
}

func (f *ForLoop) NodeType() string {
	return "for"
}

// ToJSON converts the AST to a JSON string
func (ast *AST) ToJSON() (string, error) {
	data, err := json.MarshalIndent(ast, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ToJSONBytes converts the AST to JSON bytes
func (ast *AST) ToJSONBytes() ([]byte, error) {
	return json.MarshalIndent(ast, "", "  ")
}

// Custom JSON marshaling to handle interface{} types
func (ast *AST) MarshalJSON() ([]byte, error) {
	type Alias AST
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(ast),
	})
}
