package parsing

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/calebryant/gostash-antlr/internal/parser"
)

// ASTBuilder builds an AST from the parse tree using the visitor pattern
type ASTBuilder struct {
	filterName string
	blocks     []ASTNode
}

// NewASTBuilder creates a new AST builder
func NewASTBuilder(filterName string) *ASTBuilder {
	return &ASTBuilder{
		filterName: filterName,
		blocks:     []ASTNode{},
	}
}

// Build traverses the parse tree and builds an AST
func (b *ASTBuilder) Build(filterBlock parser.IFilterblockContext) *AST {
	if filterBlock == nil {
		return &AST{
			FilterName: b.filterName,
			Blocks:     []ASTNode{},
		}
	}

	// Extract filter name from first child (ID token)
	if filterBlock.GetChild(0) != nil {
		if token, ok := filterBlock.GetChild(0).(antlr.TerminalNode); ok {
			b.filterName = token.GetText()
		}
	}

	// Process children (plugins and conditionals)
	for i := 0; i < filterBlock.GetChildCount(); i++ {
		child := filterBlock.GetChild(i)
		if parseTree, ok := child.(antlr.ParseTree); ok {
			node := b.visitChild(parseTree)
			if node != nil {
				b.blocks = append(b.blocks, node)
			}
		}
	}

	return &AST{
		FilterName: b.filterName,
		Blocks:     b.blocks,
	}
}

func (b *ASTBuilder) visitChild(child antlr.ParseTree) ASTNode {
	if child == nil {
		return nil
	}

	switch v := child.(type) {
	case *parser.PluginContext:
		return b.visitPlugin(v)
	case *parser.ConditionalblockContext:
		return b.visitConditional(v)
	}

	return nil
}

func (b *ASTBuilder) visitPlugin(ctx *parser.PluginContext) *Plugin {
	if ctx == nil {
		return nil
	}

	plugin := &Plugin{
		Type:   "plugin",
		Config: make(map[string]interface{}),
	}

	// Get plugin name (first ID)
	if ctx.ID() != nil {
		plugin.Name = ctx.ID().GetText()
	}

	// Get plugin line number
	if ctx.GetStart() != nil {
		plugin.LineNumber = ctx.GetStart().GetLine()
	}

	// Process key-value pairs
	for _, kvCtx := range ctx.AllKeyvalue() {
		if kvCtx != nil {
			key, value := b.visitKeyValue(kvCtx)
			if key != "" {
				plugin.Config[key] = value
			}
		}
	}

	return plugin
}

func (b *ASTBuilder) visitKeyValue(ctx parser.IKeyvalueContext) (string, interface{}) {
	if ctx == nil {
		return "", nil
	}

	var key string
	var value interface{}

	// Get key (lvalue)
	if ctx.Kv_lvalue() != nil {
		lvalue := ctx.Kv_lvalue()
		key = lvalue.GetText()
		// Remove quotes if present
		key = strings.Trim(key, "\"'")
	}

	// Get value (rvalue)
	if ctx.Kv_rvalue() != nil {
		value = b.visitRValue(ctx.Kv_rvalue())
	}

	return key, value
}

func (b *ASTBuilder) visitRValue(ctx parser.IKv_rvalueContext) interface{} {
	if ctx == nil {
		return nil
	}

	text := ctx.GetText()

	// Try to parse as number
	if strings.Contains(text, ".") {
		// Might be a float
		if f, err := parseFloat(text); err == nil {
			return f
		}
	} else {
		// Try integer
		if i, err := parseInt(text); err == nil {
			return i
		}
	}

	// Check for boolean
	if text == "true" || text == "false" {
		return text == "true"
	}

	// Handle hash
	if ctx.Hash() != nil {
		return b.visitHash(ctx.Hash())
	}

	// Handle list
	if ctx.List() != nil {
		return b.visitList(ctx.List())
	}

	// String or ID - remove quotes
	return strings.Trim(text, "\"'")
}

func (b *ASTBuilder) visitHash(ctx parser.IHashContext) map[string]interface{} {
	if ctx == nil {
		return nil
	}

	hash := make(map[string]interface{})

	for _, kvCtx := range ctx.AllKeyvalue() {
		if kvCtx != nil {
			key, value := b.visitKeyValue(kvCtx)
			if key != "" {
				hash[key] = value
			}
		}
	}

	return hash
}

func (b *ASTBuilder) visitList(ctx parser.IListContext) []interface{} {
	if ctx == nil {
		return nil
	}

	var list []interface{}

	for _, lvCtx := range ctx.AllList_value() {
		if lvCtx != nil {
			value := b.visitListValue(lvCtx)
			list = append(list, value)
		}
	}

	return list
}

func (b *ASTBuilder) visitListValue(ctx parser.IList_valueContext) interface{} {
	if ctx == nil {
		return nil
	}

	text := ctx.GetText()

	// Skip commas
	if text == "," {
		return nil
	}

	// Try to parse as number
	if i, err := parseInt(text); err == nil {
		return i
	}
	if f, err := parseFloat(text); err == nil {
		return f
	}

	// Check for boolean
	if text == "true" {
		return true
	}
	if text == "false" {
		return false
	}

	// String or ID - remove quotes
	return strings.Trim(text, "\"'")
}

func (b *ASTBuilder) visitConditional(ctx *parser.ConditionalblockContext) ASTNode {
	if ctx == nil {
		return nil
	}

	// Determine the type of conditional (if, else if, else, for)
	if ctx.FOR() != nil {
		return b.visitForLoop(ctx)
	}

	if ctx.IF() != nil {
		return b.visitIfBlock(ctx)
	}

	if ctx.ELSEIF() != nil {
		// This shouldn't happen at top level, handled by parent
		return nil
	}

	if ctx.ELSE() != nil {
		// This shouldn't happen at top level, handled by parent
		return nil
	}

	return nil
}

func (b *ASTBuilder) visitIfBlock(ctx *parser.ConditionalblockContext) *Conditional {
	cond := &Conditional{
		Type: "conditional",
	}

	if ctx.GetStart() != nil {
		cond.LineNumber = ctx.GetStart().GetLine()
	}

	// Get condition
	if ctx.Statement() != nil {
		cond.Condition = b.visitStatement(ctx.Statement())
	}

	// Process blocks inside the if
	for _, pluginCtx := range ctx.AllPlugin() {
		if pluginCtxCast, ok := pluginCtx.(*parser.PluginContext); ok {
			if plugin := b.visitPlugin(pluginCtxCast); plugin != nil {
				cond.Blocks = append(cond.Blocks, plugin)
			}
		}
	}

	for _, condCtx := range ctx.AllConditionalblock() {
		if condCtxCast, ok := condCtx.(*parser.ConditionalblockContext); ok {
			if node := b.visitConditional(condCtxCast); node != nil {
				cond.Blocks = append(cond.Blocks, node)
			}
		}
	}

	return cond
}

func (b *ASTBuilder) visitForLoop(ctx *parser.ConditionalblockContext) *ForLoop {
	loop := &ForLoop{
		Type: "for",
	}

	if ctx.GetStart() != nil {
		loop.LineNumber = ctx.GetStart().GetLine()
	}

	// Extract for loop components. With two loop variables the first is the
	// index/key and the second is the value: `for key, value in thing map { }`.
	forVars := ctx.AllFor_var()
	switch len(forVars) {
	case 1:
		loop.Variable = forVars[0].GetText()
	case 2:
		loop.Index = forVars[0].GetText()
		loop.Variable = forVars[1].GetText()
	}

	if ctx.For_iterable() != nil {
		loop.Iterable = sourceText(ctx.For_iterable())
	}

	// Process blocks inside the for loop
	for _, pluginCtx := range ctx.AllPlugin() {
		if pluginCtxCast, ok := pluginCtx.(*parser.PluginContext); ok {
			if plugin := b.visitPlugin(pluginCtxCast); plugin != nil {
				loop.Blocks = append(loop.Blocks, plugin)
			}
		}
	}

	for _, condCtx := range ctx.AllConditionalblock() {
		if condCtxCast, ok := condCtx.(*parser.ConditionalblockContext); ok {
			if node := b.visitConditional(condCtxCast); node != nil {
				loop.Blocks = append(loop.Blocks, node)
			}
		}
	}

	return loop
}

func (b *ASTBuilder) visitStatement(ctx parser.IStatementContext) string {
	if ctx == nil {
		return ""
	}

	return ctx.GetText()
}

// sourceText returns the original source text of a rule, including the
// whitespace between its tokens. ctx.GetText() concatenates token text and so
// would render `addition map` as `additionmap`.
func sourceText(ctx antlr.ParserRuleContext) string {
	start, stop := ctx.GetStart(), ctx.GetStop()
	if start == nil || stop == nil {
		return ctx.GetText()
	}
	stream := start.GetInputStream()
	if stream == nil {
		return ctx.GetText()
	}
	return stream.GetTextFromInterval(antlr.NewInterval(start.GetStart(), stop.GetStop()))
}

// Helper functions
func parseInt(s string) (int64, error) {
	s = strings.TrimSpace(s)
	i := 0
	_, err := fmt.Sscanf(s, "%d", &i)
	return int64(i), err
}

func parseFloat(s string) (float64, error) {
	s = strings.TrimSpace(s)
	f := 0.0
	_, err := fmt.Sscanf(s, "%f", &f)
	return f, err
}

// ExtractStringValue removes quotes from a string if present
func ExtractStringValue(text string) string {
	text = strings.TrimSpace(text)

	// Remove surrounding quotes
	if (strings.HasPrefix(text, "\"") && strings.HasSuffix(text, "\"")) ||
		(strings.HasPrefix(text, "'") && strings.HasSuffix(text, "'")) {
		return text[1 : len(text)-1]
	}

	// Handle regex patterns
	if strings.HasPrefix(text, "/") && strings.HasSuffix(text, "/") {
		return text[1 : len(text)-1]
	}

	return text
}

// ExtractRegexPattern extracts regex pattern and flags
func ExtractRegexPattern(text string) string {
	text = strings.TrimSpace(text)

	// Remove /, \/ patterns
	if strings.HasPrefix(text, "\\/") && strings.HasSuffix(text, "\\/") {
		return text[2 : len(text)-2]
	}
	if strings.HasPrefix(text, "/") && strings.HasSuffix(text, "/") {
		return text[1 : len(text)-1]
	}

	return text
}

// NormalizeCondition normalizes condition text for readability
func NormalizeCondition(condition string) string {
	// Replace bracket notation with readable format
	re := regexp.MustCompile(`\[([^\]]+)\]`)
	condition = re.ReplaceAllString(condition, "[$1]")
	return condition
}
