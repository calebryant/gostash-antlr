// Code generated from ChronicleLogstashParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ChronicleLogstashParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ChronicleLogstashParser struct {
	*antlr.BaseParser
}

var ChronicleLogstashParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func chroniclelogstashparserParserInit() {
	staticData := &ChronicleLogstashParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "'if'", "'else if'", "'else'", "'for'", "'in'", "'not'",
		"", "'}'", "'['", "']'", "'('", "')'", "", "','", "'!'", "", "'=='",
		"'!='", "'<'", "'>'", "'<='", "'>='", "'=~'", "'!~'",
	}
	staticData.SymbolicNames = []string{
		"", "WS", "COMMENT", "IF", "ELSEIF", "ELSE", "FOR", "IN", "NOT", "LBRACE",
		"RBRACE", "LBRACKET", "RBRACKET", "LPAREN", "RPAREN", "KVSEPARATOR",
		"COMMA", "BOOLNOT", "MATHOP", "EQUAL", "NOTEQUAL", "LESSTHAN", "GREATERTHAN",
		"LTEQUAL", "GTEQUAL", "MATCH", "NOTMATCH", "AND", "OR", "BOOLEAN", "STRING",
		"REGEX", "IFSTATEMENTID", "ID", "NUMBER", "FORWS", "FORCOMMA", "FORIN",
		"FORID", "FOROPENER",
	}
	staticData.RuleNames = []string{
		"filterblock", "conditionalblock", "statement", "expression", "unary_expression",
		"binary_expression", "expression_val", "math_statement", "math_expression",
		"boolean_op", "boolean_eval", "plugin", "keyvalue", "kv_lvalue", "kv_rvalue",
		"hash", "list", "if_list", "list_value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 39, 220, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0,
		43, 8, 0, 10, 0, 12, 0, 46, 9, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 5, 1, 56, 8, 1, 10, 1, 12, 1, 59, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 5, 1, 67, 8, 1, 10, 1, 12, 1, 70, 9, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 3, 1, 76, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 84, 8, 1,
		10, 1, 12, 1, 87, 9, 1, 1, 1, 3, 1, 90, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 104, 8, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 5, 2, 110, 8, 2, 10, 2, 12, 2, 113, 9, 2, 1, 3, 1, 3, 3,
		3, 117, 8, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 133, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 3, 7, 141, 8, 7, 1, 7, 1, 7, 1, 7, 5, 7, 146, 8, 7, 10, 7, 12,
		7, 149, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 11, 5, 11, 162, 8, 11, 10, 11, 12, 11, 165, 9, 11, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 173, 8, 12, 1, 13, 1, 13, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 183, 8, 14, 1, 15, 1, 15, 5,
		15, 187, 8, 15, 10, 15, 12, 15, 190, 9, 15, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 16, 5, 16, 197, 8, 16, 10, 16, 12, 16, 200, 9, 16, 3, 16, 202, 8, 16,
		1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 5, 17, 209, 8, 17, 10, 17, 12, 17, 212,
		9, 17, 3, 17, 214, 8, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 0, 2, 4, 14,
		19, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 0, 8, 1, 0, 3, 4, 2, 0, 29, 29, 32, 32, 2, 0, 32, 32, 34, 34, 1, 0,
		27, 28, 2, 0, 7, 7, 19, 26, 2, 0, 30, 30, 33, 33, 2, 0, 29, 30, 34, 34,
		3, 0, 16, 16, 29, 30, 33, 34, 237, 0, 38, 1, 0, 0, 0, 2, 89, 1, 0, 0, 0,
		4, 103, 1, 0, 0, 0, 6, 116, 1, 0, 0, 0, 8, 118, 1, 0, 0, 0, 10, 120, 1,
		0, 0, 0, 12, 132, 1, 0, 0, 0, 14, 140, 1, 0, 0, 0, 16, 150, 1, 0, 0, 0,
		18, 154, 1, 0, 0, 0, 20, 156, 1, 0, 0, 0, 22, 158, 1, 0, 0, 0, 24, 168,
		1, 0, 0, 0, 26, 174, 1, 0, 0, 0, 28, 182, 1, 0, 0, 0, 30, 184, 1, 0, 0,
		0, 32, 193, 1, 0, 0, 0, 34, 205, 1, 0, 0, 0, 36, 217, 1, 0, 0, 0, 38, 39,
		5, 33, 0, 0, 39, 44, 5, 9, 0, 0, 40, 43, 3, 22, 11, 0, 41, 43, 3, 2, 1,
		0, 42, 40, 1, 0, 0, 0, 42, 41, 1, 0, 0, 0, 43, 46, 1, 0, 0, 0, 44, 42,
		1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 47, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0,
		47, 48, 5, 10, 0, 0, 48, 49, 5, 0, 0, 1, 49, 1, 1, 0, 0, 0, 50, 51, 7,
		0, 0, 0, 51, 52, 3, 4, 2, 0, 52, 57, 5, 9, 0, 0, 53, 56, 3, 22, 11, 0,
		54, 56, 3, 2, 1, 0, 55, 53, 1, 0, 0, 0, 55, 54, 1, 0, 0, 0, 56, 59, 1,
		0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 60, 1, 0, 0, 0, 59,
		57, 1, 0, 0, 0, 60, 61, 5, 10, 0, 0, 61, 90, 1, 0, 0, 0, 62, 63, 5, 5,
		0, 0, 63, 68, 5, 9, 0, 0, 64, 67, 3, 22, 11, 0, 65, 67, 3, 2, 1, 0, 66,
		64, 1, 0, 0, 0, 66, 65, 1, 0, 0, 0, 67, 70, 1, 0, 0, 0, 68, 66, 1, 0, 0,
		0, 68, 69, 1, 0, 0, 0, 69, 71, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 71, 90,
		5, 10, 0, 0, 72, 75, 5, 6, 0, 0, 73, 74, 5, 38, 0, 0, 74, 76, 5, 36, 0,
		0, 75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 78,
		5, 38, 0, 0, 78, 79, 5, 37, 0, 0, 79, 80, 5, 38, 0, 0, 80, 85, 5, 39, 0,
		0, 81, 84, 3, 22, 11, 0, 82, 84, 3, 2, 1, 0, 83, 81, 1, 0, 0, 0, 83, 82,
		1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0,
		86, 88, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 88, 90, 5, 10, 0, 0, 89, 50, 1,
		0, 0, 0, 89, 62, 1, 0, 0, 0, 89, 72, 1, 0, 0, 0, 90, 3, 1, 0, 0, 0, 91,
		92, 6, 2, -1, 0, 92, 93, 5, 13, 0, 0, 93, 94, 3, 4, 2, 0, 94, 95, 5, 14,
		0, 0, 95, 104, 1, 0, 0, 0, 96, 97, 5, 11, 0, 0, 97, 98, 3, 4, 2, 0, 98,
		99, 5, 12, 0, 0, 99, 104, 1, 0, 0, 0, 100, 101, 5, 17, 0, 0, 101, 104,
		3, 4, 2, 2, 102, 104, 3, 6, 3, 0, 103, 91, 1, 0, 0, 0, 103, 96, 1, 0, 0,
		0, 103, 100, 1, 0, 0, 0, 103, 102, 1, 0, 0, 0, 104, 111, 1, 0, 0, 0, 105,
		106, 10, 3, 0, 0, 106, 107, 3, 18, 9, 0, 107, 108, 3, 4, 2, 4, 108, 110,
		1, 0, 0, 0, 109, 105, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111, 109, 1, 0,
		0, 0, 111, 112, 1, 0, 0, 0, 112, 5, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 114,
		117, 3, 10, 5, 0, 115, 117, 3, 8, 4, 0, 116, 114, 1, 0, 0, 0, 116, 115,
		1, 0, 0, 0, 117, 7, 1, 0, 0, 0, 118, 119, 7, 1, 0, 0, 119, 9, 1, 0, 0,
		0, 120, 121, 3, 12, 6, 0, 121, 122, 3, 20, 10, 0, 122, 123, 3, 12, 6, 0,
		123, 11, 1, 0, 0, 0, 124, 133, 3, 14, 7, 0, 125, 133, 5, 34, 0, 0, 126,
		133, 3, 34, 17, 0, 127, 133, 5, 32, 0, 0, 128, 133, 5, 30, 0, 0, 129, 133,
		5, 31, 0, 0, 130, 133, 5, 29, 0, 0, 131, 133, 5, 33, 0, 0, 132, 124, 1,
		0, 0, 0, 132, 125, 1, 0, 0, 0, 132, 126, 1, 0, 0, 0, 132, 127, 1, 0, 0,
		0, 132, 128, 1, 0, 0, 0, 132, 129, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 132,
		131, 1, 0, 0, 0, 133, 13, 1, 0, 0, 0, 134, 135, 6, 7, -1, 0, 135, 136,
		5, 13, 0, 0, 136, 137, 3, 14, 7, 0, 137, 138, 5, 14, 0, 0, 138, 141, 1,
		0, 0, 0, 139, 141, 3, 16, 8, 0, 140, 134, 1, 0, 0, 0, 140, 139, 1, 0, 0,
		0, 141, 147, 1, 0, 0, 0, 142, 143, 10, 2, 0, 0, 143, 144, 5, 18, 0, 0,
		144, 146, 3, 14, 7, 3, 145, 142, 1, 0, 0, 0, 146, 149, 1, 0, 0, 0, 147,
		145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 15, 1, 0, 0, 0, 149, 147, 1,
		0, 0, 0, 150, 151, 7, 2, 0, 0, 151, 152, 5, 18, 0, 0, 152, 153, 7, 2, 0,
		0, 153, 17, 1, 0, 0, 0, 154, 155, 7, 3, 0, 0, 155, 19, 1, 0, 0, 0, 156,
		157, 7, 4, 0, 0, 157, 21, 1, 0, 0, 0, 158, 159, 5, 33, 0, 0, 159, 163,
		5, 9, 0, 0, 160, 162, 3, 24, 12, 0, 161, 160, 1, 0, 0, 0, 162, 165, 1,
		0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 166, 1, 0, 0,
		0, 165, 163, 1, 0, 0, 0, 166, 167, 5, 10, 0, 0, 167, 23, 1, 0, 0, 0, 168,
		169, 3, 26, 13, 0, 169, 170, 5, 15, 0, 0, 170, 172, 3, 28, 14, 0, 171,
		173, 5, 16, 0, 0, 172, 171, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 25,
		1, 0, 0, 0, 174, 175, 7, 5, 0, 0, 175, 27, 1, 0, 0, 0, 176, 183, 5, 34,
		0, 0, 177, 183, 3, 32, 16, 0, 178, 183, 3, 30, 15, 0, 179, 183, 5, 30,
		0, 0, 180, 183, 5, 29, 0, 0, 181, 183, 5, 33, 0, 0, 182, 176, 1, 0, 0,
		0, 182, 177, 1, 0, 0, 0, 182, 178, 1, 0, 0, 0, 182, 179, 1, 0, 0, 0, 182,
		180, 1, 0, 0, 0, 182, 181, 1, 0, 0, 0, 183, 29, 1, 0, 0, 0, 184, 188, 5,
		9, 0, 0, 185, 187, 3, 24, 12, 0, 186, 185, 1, 0, 0, 0, 187, 190, 1, 0,
		0, 0, 188, 186, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 191, 1, 0, 0, 0,
		190, 188, 1, 0, 0, 0, 191, 192, 5, 10, 0, 0, 192, 31, 1, 0, 0, 0, 193,
		201, 5, 11, 0, 0, 194, 198, 3, 36, 18, 0, 195, 197, 3, 36, 18, 0, 196,
		195, 1, 0, 0, 0, 197, 200, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 198, 199,
		1, 0, 0, 0, 199, 202, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 201, 194, 1, 0,
		0, 0, 201, 202, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 204, 5, 12, 0, 0,
		204, 33, 1, 0, 0, 0, 205, 213, 5, 11, 0, 0, 206, 210, 7, 6, 0, 0, 207,
		209, 3, 36, 18, 0, 208, 207, 1, 0, 0, 0, 209, 212, 1, 0, 0, 0, 210, 208,
		1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 214, 1, 0, 0, 0, 212, 210, 1, 0,
		0, 0, 213, 206, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0,
		215, 216, 5, 12, 0, 0, 216, 35, 1, 0, 0, 0, 217, 218, 7, 7, 0, 0, 218,
		37, 1, 0, 0, 0, 24, 42, 44, 55, 57, 66, 68, 75, 83, 85, 89, 103, 111, 116,
		132, 140, 147, 163, 172, 182, 188, 198, 201, 210, 213,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ChronicleLogstashParserInit initializes any static state used to implement ChronicleLogstashParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewChronicleLogstashParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ChronicleLogstashParserInit() {
	staticData := &ChronicleLogstashParserParserStaticData
	staticData.once.Do(chroniclelogstashparserParserInit)
}

// NewChronicleLogstashParser produces a new parser instance for the optional input antlr.TokenStream.
func NewChronicleLogstashParser(input antlr.TokenStream) *ChronicleLogstashParser {
	ChronicleLogstashParserInit()
	this := new(ChronicleLogstashParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ChronicleLogstashParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ChronicleLogstashParser.g4"

	return this
}

// ChronicleLogstashParser tokens.
const (
	ChronicleLogstashParserEOF           = antlr.TokenEOF
	ChronicleLogstashParserWS            = 1
	ChronicleLogstashParserCOMMENT       = 2
	ChronicleLogstashParserIF            = 3
	ChronicleLogstashParserELSEIF        = 4
	ChronicleLogstashParserELSE          = 5
	ChronicleLogstashParserFOR           = 6
	ChronicleLogstashParserIN            = 7
	ChronicleLogstashParserNOT           = 8
	ChronicleLogstashParserLBRACE        = 9
	ChronicleLogstashParserRBRACE        = 10
	ChronicleLogstashParserLBRACKET      = 11
	ChronicleLogstashParserRBRACKET      = 12
	ChronicleLogstashParserLPAREN        = 13
	ChronicleLogstashParserRPAREN        = 14
	ChronicleLogstashParserKVSEPARATOR   = 15
	ChronicleLogstashParserCOMMA         = 16
	ChronicleLogstashParserBOOLNOT       = 17
	ChronicleLogstashParserMATHOP        = 18
	ChronicleLogstashParserEQUAL         = 19
	ChronicleLogstashParserNOTEQUAL      = 20
	ChronicleLogstashParserLESSTHAN      = 21
	ChronicleLogstashParserGREATERTHAN   = 22
	ChronicleLogstashParserLTEQUAL       = 23
	ChronicleLogstashParserGTEQUAL       = 24
	ChronicleLogstashParserMATCH         = 25
	ChronicleLogstashParserNOTMATCH      = 26
	ChronicleLogstashParserAND           = 27
	ChronicleLogstashParserOR            = 28
	ChronicleLogstashParserBOOLEAN       = 29
	ChronicleLogstashParserSTRING        = 30
	ChronicleLogstashParserREGEX         = 31
	ChronicleLogstashParserIFSTATEMENTID = 32
	ChronicleLogstashParserID            = 33
	ChronicleLogstashParserNUMBER        = 34
	ChronicleLogstashParserFORWS         = 35
	ChronicleLogstashParserFORCOMMA      = 36
	ChronicleLogstashParserFORIN         = 37
	ChronicleLogstashParserFORID         = 38
	ChronicleLogstashParserFOROPENER     = 39
)

// ChronicleLogstashParser rules.
const (
	ChronicleLogstashParserRULE_filterblock       = 0
	ChronicleLogstashParserRULE_conditionalblock  = 1
	ChronicleLogstashParserRULE_statement         = 2
	ChronicleLogstashParserRULE_expression        = 3
	ChronicleLogstashParserRULE_unary_expression  = 4
	ChronicleLogstashParserRULE_binary_expression = 5
	ChronicleLogstashParserRULE_expression_val    = 6
	ChronicleLogstashParserRULE_math_statement    = 7
	ChronicleLogstashParserRULE_math_expression   = 8
	ChronicleLogstashParserRULE_boolean_op        = 9
	ChronicleLogstashParserRULE_boolean_eval      = 10
	ChronicleLogstashParserRULE_plugin            = 11
	ChronicleLogstashParserRULE_keyvalue          = 12
	ChronicleLogstashParserRULE_kv_lvalue         = 13
	ChronicleLogstashParserRULE_kv_rvalue         = 14
	ChronicleLogstashParserRULE_hash              = 15
	ChronicleLogstashParserRULE_list              = 16
	ChronicleLogstashParserRULE_if_list           = 17
	ChronicleLogstashParserRULE_list_value        = 18
)

// IFilterblockContext is an interface to support dynamic dispatch.
type IFilterblockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	EOF() antlr.TerminalNode
	AllPlugin() []IPluginContext
	Plugin(i int) IPluginContext
	AllConditionalblock() []IConditionalblockContext
	Conditionalblock(i int) IConditionalblockContext

	// IsFilterblockContext differentiates from other interfaces.
	IsFilterblockContext()
}

type FilterblockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterblockContext() *FilterblockContext {
	var p = new(FilterblockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_filterblock
	return p
}

func InitEmptyFilterblockContext(p *FilterblockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_filterblock
}

func (*FilterblockContext) IsFilterblockContext() {}

func NewFilterblockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterblockContext {
	var p = new(FilterblockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_filterblock

	return p
}

func (s *FilterblockContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterblockContext) ID() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserID, 0)
}

func (s *FilterblockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserLBRACE, 0)
}

func (s *FilterblockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserRBRACE, 0)
}

func (s *FilterblockContext) EOF() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserEOF, 0)
}

func (s *FilterblockContext) AllPlugin() []IPluginContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPluginContext); ok {
			len++
		}
	}

	tst := make([]IPluginContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPluginContext); ok {
			tst[i] = t.(IPluginContext)
			i++
		}
	}

	return tst
}

func (s *FilterblockContext) Plugin(i int) IPluginContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPluginContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPluginContext)
}

func (s *FilterblockContext) AllConditionalblock() []IConditionalblockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionalblockContext); ok {
			len++
		}
	}

	tst := make([]IConditionalblockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionalblockContext); ok {
			tst[i] = t.(IConditionalblockContext)
			i++
		}
	}

	return tst
}

func (s *FilterblockContext) Conditionalblock(i int) IConditionalblockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalblockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalblockContext)
}

func (s *FilterblockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterblockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterblockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterFilterblock(s)
	}
}

func (s *FilterblockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitFilterblock(s)
	}
}

func (p *ChronicleLogstashParser) Filterblock() (localctx IFilterblockContext) {
	localctx = NewFilterblockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ChronicleLogstashParserRULE_filterblock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(ChronicleLogstashParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.Match(ChronicleLogstashParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8589934712) != 0 {
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case ChronicleLogstashParserID:
			{
				p.SetState(40)
				p.Plugin()
			}

		case ChronicleLogstashParserIF, ChronicleLogstashParserELSEIF, ChronicleLogstashParserELSE, ChronicleLogstashParserFOR:
			{
				p.SetState(41)
				p.Conditionalblock()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(47)
		p.Match(ChronicleLogstashParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		p.Match(ChronicleLogstashParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionalblockContext is an interface to support dynamic dispatch.
type IConditionalblockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Statement() IStatementContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	IF() antlr.TerminalNode
	ELSEIF() antlr.TerminalNode
	AllPlugin() []IPluginContext
	Plugin(i int) IPluginContext
	AllConditionalblock() []IConditionalblockContext
	Conditionalblock(i int) IConditionalblockContext
	ELSE() antlr.TerminalNode
	FOR() antlr.TerminalNode
	AllFORID() []antlr.TerminalNode
	FORID(i int) antlr.TerminalNode
	FORIN() antlr.TerminalNode
	FOROPENER() antlr.TerminalNode
	FORCOMMA() antlr.TerminalNode

	// IsConditionalblockContext differentiates from other interfaces.
	IsConditionalblockContext()
}

type ConditionalblockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalblockContext() *ConditionalblockContext {
	var p = new(ConditionalblockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_conditionalblock
	return p
}

func InitEmptyConditionalblockContext(p *ConditionalblockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_conditionalblock
}

func (*ConditionalblockContext) IsConditionalblockContext() {}

func NewConditionalblockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalblockContext {
	var p = new(ConditionalblockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_conditionalblock

	return p
}

func (s *ConditionalblockContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalblockContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ConditionalblockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserLBRACE, 0)
}

func (s *ConditionalblockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserRBRACE, 0)
}

func (s *ConditionalblockContext) IF() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserIF, 0)
}

func (s *ConditionalblockContext) ELSEIF() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserELSEIF, 0)
}

func (s *ConditionalblockContext) AllPlugin() []IPluginContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPluginContext); ok {
			len++
		}
	}

	tst := make([]IPluginContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPluginContext); ok {
			tst[i] = t.(IPluginContext)
			i++
		}
	}

	return tst
}

func (s *ConditionalblockContext) Plugin(i int) IPluginContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPluginContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPluginContext)
}

func (s *ConditionalblockContext) AllConditionalblock() []IConditionalblockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionalblockContext); ok {
			len++
		}
	}

	tst := make([]IConditionalblockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionalblockContext); ok {
			tst[i] = t.(IConditionalblockContext)
			i++
		}
	}

	return tst
}

func (s *ConditionalblockContext) Conditionalblock(i int) IConditionalblockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalblockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalblockContext)
}

func (s *ConditionalblockContext) ELSE() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserELSE, 0)
}

func (s *ConditionalblockContext) FOR() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserFOR, 0)
}

func (s *ConditionalblockContext) AllFORID() []antlr.TerminalNode {
	return s.GetTokens(ChronicleLogstashParserFORID)
}

func (s *ConditionalblockContext) FORID(i int) antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserFORID, i)
}

func (s *ConditionalblockContext) FORIN() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserFORIN, 0)
}

func (s *ConditionalblockContext) FOROPENER() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserFOROPENER, 0)
}

func (s *ConditionalblockContext) FORCOMMA() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserFORCOMMA, 0)
}

func (s *ConditionalblockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalblockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalblockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterConditionalblock(s)
	}
}

func (s *ConditionalblockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitConditionalblock(s)
	}
}

func (p *ChronicleLogstashParser) Conditionalblock() (localctx IConditionalblockContext) {
	localctx = NewConditionalblockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ChronicleLogstashParserRULE_conditionalblock)
	var _la int

	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ChronicleLogstashParserIF, ChronicleLogstashParserELSEIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ChronicleLogstashParserIF || _la == ChronicleLogstashParserELSEIF) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(51)
			p.statement(0)
		}
		{
			p.SetState(52)
			p.Match(ChronicleLogstashParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8589934712) != 0 {
			p.SetState(55)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case ChronicleLogstashParserID:
				{
					p.SetState(53)
					p.Plugin()
				}

			case ChronicleLogstashParserIF, ChronicleLogstashParserELSEIF, ChronicleLogstashParserELSE, ChronicleLogstashParserFOR:
				{
					p.SetState(54)
					p.Conditionalblock()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(59)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(60)
			p.Match(ChronicleLogstashParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserELSE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)
			p.Match(ChronicleLogstashParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.Match(ChronicleLogstashParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8589934712) != 0 {
			p.SetState(66)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case ChronicleLogstashParserID:
				{
					p.SetState(64)
					p.Plugin()
				}

			case ChronicleLogstashParserIF, ChronicleLogstashParserELSEIF, ChronicleLogstashParserELSE, ChronicleLogstashParserFOR:
				{
					p.SetState(65)
					p.Conditionalblock()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(70)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(71)
			p.Match(ChronicleLogstashParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserFOR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(72)
			p.Match(ChronicleLogstashParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(73)
				p.Match(ChronicleLogstashParserFORID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(74)
				p.Match(ChronicleLogstashParserFORCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(77)
			p.Match(ChronicleLogstashParserFORID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.Match(ChronicleLogstashParserFORIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.Match(ChronicleLogstashParserFORID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)
			p.Match(ChronicleLogstashParserFOROPENER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8589934712) != 0 {
			p.SetState(83)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case ChronicleLogstashParserID:
				{
					p.SetState(81)
					p.Plugin()
				}

			case ChronicleLogstashParserIF, ChronicleLogstashParserELSEIF, ChronicleLogstashParserELSE, ChronicleLogstashParserFOR:
				{
					p.SetState(82)
					p.Conditionalblock()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(87)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(88)
			p.Match(ChronicleLogstashParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	RPAREN() antlr.TerminalNode
	LBRACKET() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	BOOLNOT() antlr.TerminalNode
	Expression() IExpressionContext
	Boolean_op() IBoolean_opContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserLPAREN, 0)
}

func (s *StatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserRPAREN, 0)
}

func (s *StatementContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserLBRACKET, 0)
}

func (s *StatementContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserRBRACKET, 0)
}

func (s *StatementContext) BOOLNOT() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserBOOLNOT, 0)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) Boolean_op() IBoolean_opContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_opContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolean_opContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *ChronicleLogstashParser) Statement() (localctx IStatementContext) {
	return p.statement(0)
}

func (p *ChronicleLogstashParser) statement(_p int) (localctx IStatementContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewStatementContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IStatementContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, ChronicleLogstashParserRULE_statement, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(92)
			p.Match(ChronicleLogstashParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(93)
			p.statement(0)
		}
		{
			p.SetState(94)
			p.Match(ChronicleLogstashParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(96)
			p.Match(ChronicleLogstashParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(97)
			p.statement(0)
		}
		{
			p.SetState(98)
			p.Match(ChronicleLogstashParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(100)
			p.Match(ChronicleLogstashParserBOOLNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.statement(2)
		}

	case 4:
		{
			p.SetState(102)
			p.Expression()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewStatementContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ChronicleLogstashParserRULE_statement)
			p.SetState(105)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(106)
				p.Boolean_op()
			}
			{
				p.SetState(107)
				p.statement(4)
			}

		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Binary_expression() IBinary_expressionContext
	Unary_expression() IUnary_expressionContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Binary_expression() IBinary_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinary_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinary_expressionContext)
}

func (s *ExpressionContext) Unary_expression() IUnary_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnary_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnary_expressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ChronicleLogstashParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ChronicleLogstashParserRULE_expression)
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)
			p.Binary_expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(115)
			p.Unary_expression()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnary_expressionContext is an interface to support dynamic dispatch.
type IUnary_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IFSTATEMENTID() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode

	// IsUnary_expressionContext differentiates from other interfaces.
	IsUnary_expressionContext()
}

type Unary_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnary_expressionContext() *Unary_expressionContext {
	var p = new(Unary_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_unary_expression
	return p
}

func InitEmptyUnary_expressionContext(p *Unary_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_unary_expression
}

func (*Unary_expressionContext) IsUnary_expressionContext() {}

func NewUnary_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unary_expressionContext {
	var p = new(Unary_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_unary_expression

	return p
}

func (s *Unary_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Unary_expressionContext) IFSTATEMENTID() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserIFSTATEMENTID, 0)
}

func (s *Unary_expressionContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserBOOLEAN, 0)
}

func (s *Unary_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unary_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unary_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterUnary_expression(s)
	}
}

func (s *Unary_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitUnary_expression(s)
	}
}

func (p *ChronicleLogstashParser) Unary_expression() (localctx IUnary_expressionContext) {
	localctx = NewUnary_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ChronicleLogstashParserRULE_unary_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ChronicleLogstashParserBOOLEAN || _la == ChronicleLogstashParserIFSTATEMENTID) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinary_expressionContext is an interface to support dynamic dispatch.
type IBinary_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression_val() []IExpression_valContext
	Expression_val(i int) IExpression_valContext
	Boolean_eval() IBoolean_evalContext

	// IsBinary_expressionContext differentiates from other interfaces.
	IsBinary_expressionContext()
}

type Binary_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinary_expressionContext() *Binary_expressionContext {
	var p = new(Binary_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_binary_expression
	return p
}

func InitEmptyBinary_expressionContext(p *Binary_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_binary_expression
}

func (*Binary_expressionContext) IsBinary_expressionContext() {}

func NewBinary_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Binary_expressionContext {
	var p = new(Binary_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_binary_expression

	return p
}

func (s *Binary_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Binary_expressionContext) AllExpression_val() []IExpression_valContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpression_valContext); ok {
			len++
		}
	}

	tst := make([]IExpression_valContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpression_valContext); ok {
			tst[i] = t.(IExpression_valContext)
			i++
		}
	}

	return tst
}

func (s *Binary_expressionContext) Expression_val(i int) IExpression_valContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_valContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpression_valContext)
}

func (s *Binary_expressionContext) Boolean_eval() IBoolean_evalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_evalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolean_evalContext)
}

func (s *Binary_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Binary_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Binary_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterBinary_expression(s)
	}
}

func (s *Binary_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitBinary_expression(s)
	}
}

func (p *ChronicleLogstashParser) Binary_expression() (localctx IBinary_expressionContext) {
	localctx = NewBinary_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ChronicleLogstashParserRULE_binary_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Expression_val()
	}
	{
		p.SetState(121)
		p.Boolean_eval()
	}
	{
		p.SetState(122)
		p.Expression_val()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpression_valContext is an interface to support dynamic dispatch.
type IExpression_valContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Math_statement() IMath_statementContext
	NUMBER() antlr.TerminalNode
	If_list() IIf_listContext
	IFSTATEMENTID() antlr.TerminalNode
	STRING() antlr.TerminalNode
	REGEX() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsExpression_valContext differentiates from other interfaces.
	IsExpression_valContext()
}

type Expression_valContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_valContext() *Expression_valContext {
	var p = new(Expression_valContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_expression_val
	return p
}

func InitEmptyExpression_valContext(p *Expression_valContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_expression_val
}

func (*Expression_valContext) IsExpression_valContext() {}

func NewExpression_valContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_valContext {
	var p = new(Expression_valContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_expression_val

	return p
}

func (s *Expression_valContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_valContext) Math_statement() IMath_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMath_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMath_statementContext)
}

func (s *Expression_valContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserNUMBER, 0)
}

func (s *Expression_valContext) If_list() IIf_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_listContext)
}

func (s *Expression_valContext) IFSTATEMENTID() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserIFSTATEMENTID, 0)
}

func (s *Expression_valContext) STRING() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserSTRING, 0)
}

func (s *Expression_valContext) REGEX() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserREGEX, 0)
}

func (s *Expression_valContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserBOOLEAN, 0)
}

func (s *Expression_valContext) ID() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserID, 0)
}

func (s *Expression_valContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_valContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_valContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterExpression_val(s)
	}
}

func (s *Expression_valContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitExpression_val(s)
	}
}

func (p *ChronicleLogstashParser) Expression_val() (localctx IExpression_valContext) {
	localctx = NewExpression_valContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ChronicleLogstashParserRULE_expression_val)
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(124)
			p.math_statement(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.Match(ChronicleLogstashParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(126)
			p.If_list()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(127)
			p.Match(ChronicleLogstashParserIFSTATEMENTID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(128)
			p.Match(ChronicleLogstashParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(129)
			p.Match(ChronicleLogstashParserREGEX)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(130)
			p.Match(ChronicleLogstashParserBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(131)
			p.Match(ChronicleLogstashParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMath_statementContext is an interface to support dynamic dispatch.
type IMath_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	AllMath_statement() []IMath_statementContext
	Math_statement(i int) IMath_statementContext
	RPAREN() antlr.TerminalNode
	Math_expression() IMath_expressionContext
	MATHOP() antlr.TerminalNode

	// IsMath_statementContext differentiates from other interfaces.
	IsMath_statementContext()
}

type Math_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMath_statementContext() *Math_statementContext {
	var p = new(Math_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_math_statement
	return p
}

func InitEmptyMath_statementContext(p *Math_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_math_statement
}

func (*Math_statementContext) IsMath_statementContext() {}

func NewMath_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Math_statementContext {
	var p = new(Math_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_math_statement

	return p
}

func (s *Math_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Math_statementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserLPAREN, 0)
}

func (s *Math_statementContext) AllMath_statement() []IMath_statementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMath_statementContext); ok {
			len++
		}
	}

	tst := make([]IMath_statementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMath_statementContext); ok {
			tst[i] = t.(IMath_statementContext)
			i++
		}
	}

	return tst
}

func (s *Math_statementContext) Math_statement(i int) IMath_statementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMath_statementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMath_statementContext)
}

func (s *Math_statementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserRPAREN, 0)
}

func (s *Math_statementContext) Math_expression() IMath_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMath_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMath_expressionContext)
}

func (s *Math_statementContext) MATHOP() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserMATHOP, 0)
}

func (s *Math_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Math_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Math_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterMath_statement(s)
	}
}

func (s *Math_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitMath_statement(s)
	}
}

func (p *ChronicleLogstashParser) Math_statement() (localctx IMath_statementContext) {
	return p.math_statement(0)
}

func (p *ChronicleLogstashParser) math_statement(_p int) (localctx IMath_statementContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewMath_statementContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMath_statementContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, ChronicleLogstashParserRULE_math_statement, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ChronicleLogstashParserLPAREN:
		{
			p.SetState(135)
			p.Match(ChronicleLogstashParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.math_statement(0)
		}
		{
			p.SetState(137)
			p.Match(ChronicleLogstashParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserIFSTATEMENTID, ChronicleLogstashParserNUMBER:
		{
			p.SetState(139)
			p.Math_expression()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMath_statementContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ChronicleLogstashParserRULE_math_statement)
			p.SetState(142)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(143)
				p.Match(ChronicleLogstashParserMATHOP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(144)
				p.math_statement(3)
			}

		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMath_expressionContext is an interface to support dynamic dispatch.
type IMath_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MATHOP() antlr.TerminalNode
	AllIFSTATEMENTID() []antlr.TerminalNode
	IFSTATEMENTID(i int) antlr.TerminalNode
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode

	// IsMath_expressionContext differentiates from other interfaces.
	IsMath_expressionContext()
}

type Math_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMath_expressionContext() *Math_expressionContext {
	var p = new(Math_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_math_expression
	return p
}

func InitEmptyMath_expressionContext(p *Math_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_math_expression
}

func (*Math_expressionContext) IsMath_expressionContext() {}

func NewMath_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Math_expressionContext {
	var p = new(Math_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_math_expression

	return p
}

func (s *Math_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Math_expressionContext) MATHOP() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserMATHOP, 0)
}

func (s *Math_expressionContext) AllIFSTATEMENTID() []antlr.TerminalNode {
	return s.GetTokens(ChronicleLogstashParserIFSTATEMENTID)
}

func (s *Math_expressionContext) IFSTATEMENTID(i int) antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserIFSTATEMENTID, i)
}

func (s *Math_expressionContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(ChronicleLogstashParserNUMBER)
}

func (s *Math_expressionContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserNUMBER, i)
}

func (s *Math_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Math_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Math_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterMath_expression(s)
	}
}

func (s *Math_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitMath_expression(s)
	}
}

func (p *ChronicleLogstashParser) Math_expression() (localctx IMath_expressionContext) {
	localctx = NewMath_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ChronicleLogstashParserRULE_math_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ChronicleLogstashParserIFSTATEMENTID || _la == ChronicleLogstashParserNUMBER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(151)
		p.Match(ChronicleLogstashParserMATHOP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ChronicleLogstashParserIFSTATEMENTID || _la == ChronicleLogstashParserNUMBER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoolean_opContext is an interface to support dynamic dispatch.
type IBoolean_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsBoolean_opContext differentiates from other interfaces.
	IsBoolean_opContext()
}

type Boolean_opContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolean_opContext() *Boolean_opContext {
	var p = new(Boolean_opContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_boolean_op
	return p
}

func InitEmptyBoolean_opContext(p *Boolean_opContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_boolean_op
}

func (*Boolean_opContext) IsBoolean_opContext() {}

func NewBoolean_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_opContext {
	var p = new(Boolean_opContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_boolean_op

	return p
}

func (s *Boolean_opContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_opContext) AND() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserAND, 0)
}

func (s *Boolean_opContext) OR() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserOR, 0)
}

func (s *Boolean_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Boolean_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterBoolean_op(s)
	}
}

func (s *Boolean_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitBoolean_op(s)
	}
}

func (p *ChronicleLogstashParser) Boolean_op() (localctx IBoolean_opContext) {
	localctx = NewBoolean_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ChronicleLogstashParserRULE_boolean_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ChronicleLogstashParserAND || _la == ChronicleLogstashParserOR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoolean_evalContext is an interface to support dynamic dispatch.
type IBoolean_evalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQUAL() antlr.TerminalNode
	NOTEQUAL() antlr.TerminalNode
	LESSTHAN() antlr.TerminalNode
	GREATERTHAN() antlr.TerminalNode
	LTEQUAL() antlr.TerminalNode
	GTEQUAL() antlr.TerminalNode
	MATCH() antlr.TerminalNode
	NOTMATCH() antlr.TerminalNode
	IN() antlr.TerminalNode

	// IsBoolean_evalContext differentiates from other interfaces.
	IsBoolean_evalContext()
}

type Boolean_evalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolean_evalContext() *Boolean_evalContext {
	var p = new(Boolean_evalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_boolean_eval
	return p
}

func InitEmptyBoolean_evalContext(p *Boolean_evalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_boolean_eval
}

func (*Boolean_evalContext) IsBoolean_evalContext() {}

func NewBoolean_evalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_evalContext {
	var p = new(Boolean_evalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_boolean_eval

	return p
}

func (s *Boolean_evalContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_evalContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserEQUAL, 0)
}

func (s *Boolean_evalContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserNOTEQUAL, 0)
}

func (s *Boolean_evalContext) LESSTHAN() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserLESSTHAN, 0)
}

func (s *Boolean_evalContext) GREATERTHAN() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserGREATERTHAN, 0)
}

func (s *Boolean_evalContext) LTEQUAL() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserLTEQUAL, 0)
}

func (s *Boolean_evalContext) GTEQUAL() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserGTEQUAL, 0)
}

func (s *Boolean_evalContext) MATCH() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserMATCH, 0)
}

func (s *Boolean_evalContext) NOTMATCH() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserNOTMATCH, 0)
}

func (s *Boolean_evalContext) IN() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserIN, 0)
}

func (s *Boolean_evalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_evalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Boolean_evalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterBoolean_eval(s)
	}
}

func (s *Boolean_evalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitBoolean_eval(s)
	}
}

func (p *ChronicleLogstashParser) Boolean_eval() (localctx IBoolean_evalContext) {
	localctx = NewBoolean_evalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ChronicleLogstashParserRULE_boolean_eval)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&133693568) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPluginContext is an interface to support dynamic dispatch.
type IPluginContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllKeyvalue() []IKeyvalueContext
	Keyvalue(i int) IKeyvalueContext

	// IsPluginContext differentiates from other interfaces.
	IsPluginContext()
}

type PluginContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPluginContext() *PluginContext {
	var p = new(PluginContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_plugin
	return p
}

func InitEmptyPluginContext(p *PluginContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_plugin
}

func (*PluginContext) IsPluginContext() {}

func NewPluginContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PluginContext {
	var p = new(PluginContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_plugin

	return p
}

func (s *PluginContext) GetParser() antlr.Parser { return s.parser }

func (s *PluginContext) ID() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserID, 0)
}

func (s *PluginContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserLBRACE, 0)
}

func (s *PluginContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserRBRACE, 0)
}

func (s *PluginContext) AllKeyvalue() []IKeyvalueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKeyvalueContext); ok {
			len++
		}
	}

	tst := make([]IKeyvalueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKeyvalueContext); ok {
			tst[i] = t.(IKeyvalueContext)
			i++
		}
	}

	return tst
}

func (s *PluginContext) Keyvalue(i int) IKeyvalueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyvalueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyvalueContext)
}

func (s *PluginContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PluginContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PluginContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterPlugin(s)
	}
}

func (s *PluginContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitPlugin(s)
	}
}

func (p *ChronicleLogstashParser) Plugin() (localctx IPluginContext) {
	localctx = NewPluginContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ChronicleLogstashParserRULE_plugin)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Match(ChronicleLogstashParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.Match(ChronicleLogstashParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ChronicleLogstashParserSTRING || _la == ChronicleLogstashParserID {
		{
			p.SetState(160)
			p.Keyvalue()
		}

		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(166)
		p.Match(ChronicleLogstashParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeyvalueContext is an interface to support dynamic dispatch.
type IKeyvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Kv_lvalue() IKv_lvalueContext
	KVSEPARATOR() antlr.TerminalNode
	Kv_rvalue() IKv_rvalueContext
	COMMA() antlr.TerminalNode

	// IsKeyvalueContext differentiates from other interfaces.
	IsKeyvalueContext()
}

type KeyvalueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyvalueContext() *KeyvalueContext {
	var p = new(KeyvalueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_keyvalue
	return p
}

func InitEmptyKeyvalueContext(p *KeyvalueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_keyvalue
}

func (*KeyvalueContext) IsKeyvalueContext() {}

func NewKeyvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyvalueContext {
	var p = new(KeyvalueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_keyvalue

	return p
}

func (s *KeyvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyvalueContext) Kv_lvalue() IKv_lvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKv_lvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKv_lvalueContext)
}

func (s *KeyvalueContext) KVSEPARATOR() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserKVSEPARATOR, 0)
}

func (s *KeyvalueContext) Kv_rvalue() IKv_rvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKv_rvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKv_rvalueContext)
}

func (s *KeyvalueContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserCOMMA, 0)
}

func (s *KeyvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterKeyvalue(s)
	}
}

func (s *KeyvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitKeyvalue(s)
	}
}

func (p *ChronicleLogstashParser) Keyvalue() (localctx IKeyvalueContext) {
	localctx = NewKeyvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ChronicleLogstashParserRULE_keyvalue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Kv_lvalue()
	}
	{
		p.SetState(169)
		p.Match(ChronicleLogstashParserKVSEPARATOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Kv_rvalue()
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ChronicleLogstashParserCOMMA {
		{
			p.SetState(171)
			p.Match(ChronicleLogstashParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKv_lvalueContext is an interface to support dynamic dispatch.
type IKv_lvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsKv_lvalueContext differentiates from other interfaces.
	IsKv_lvalueContext()
}

type Kv_lvalueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKv_lvalueContext() *Kv_lvalueContext {
	var p = new(Kv_lvalueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_kv_lvalue
	return p
}

func InitEmptyKv_lvalueContext(p *Kv_lvalueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_kv_lvalue
}

func (*Kv_lvalueContext) IsKv_lvalueContext() {}

func NewKv_lvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Kv_lvalueContext {
	var p = new(Kv_lvalueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_kv_lvalue

	return p
}

func (s *Kv_lvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *Kv_lvalueContext) ID() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserID, 0)
}

func (s *Kv_lvalueContext) STRING() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserSTRING, 0)
}

func (s *Kv_lvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Kv_lvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Kv_lvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterKv_lvalue(s)
	}
}

func (s *Kv_lvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitKv_lvalue(s)
	}
}

func (p *ChronicleLogstashParser) Kv_lvalue() (localctx IKv_lvalueContext) {
	localctx = NewKv_lvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ChronicleLogstashParserRULE_kv_lvalue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ChronicleLogstashParserSTRING || _la == ChronicleLogstashParserID) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKv_rvalueContext is an interface to support dynamic dispatch.
type IKv_rvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	List() IListContext
	Hash() IHashContext
	STRING() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsKv_rvalueContext differentiates from other interfaces.
	IsKv_rvalueContext()
}

type Kv_rvalueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKv_rvalueContext() *Kv_rvalueContext {
	var p = new(Kv_rvalueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_kv_rvalue
	return p
}

func InitEmptyKv_rvalueContext(p *Kv_rvalueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_kv_rvalue
}

func (*Kv_rvalueContext) IsKv_rvalueContext() {}

func NewKv_rvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Kv_rvalueContext {
	var p = new(Kv_rvalueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_kv_rvalue

	return p
}

func (s *Kv_rvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *Kv_rvalueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserNUMBER, 0)
}

func (s *Kv_rvalueContext) List() IListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *Kv_rvalueContext) Hash() IHashContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHashContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHashContext)
}

func (s *Kv_rvalueContext) STRING() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserSTRING, 0)
}

func (s *Kv_rvalueContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserBOOLEAN, 0)
}

func (s *Kv_rvalueContext) ID() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserID, 0)
}

func (s *Kv_rvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Kv_rvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Kv_rvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterKv_rvalue(s)
	}
}

func (s *Kv_rvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitKv_rvalue(s)
	}
}

func (p *ChronicleLogstashParser) Kv_rvalue() (localctx IKv_rvalueContext) {
	localctx = NewKv_rvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ChronicleLogstashParserRULE_kv_rvalue)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ChronicleLogstashParserNUMBER:
		{
			p.SetState(176)
			p.Match(ChronicleLogstashParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserLBRACKET:
		{
			p.SetState(177)
			p.List()
		}

	case ChronicleLogstashParserLBRACE:
		{
			p.SetState(178)
			p.Hash()
		}

	case ChronicleLogstashParserSTRING:
		{
			p.SetState(179)
			p.Match(ChronicleLogstashParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserBOOLEAN:
		{
			p.SetState(180)
			p.Match(ChronicleLogstashParserBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserID:
		{
			p.SetState(181)
			p.Match(ChronicleLogstashParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHashContext is an interface to support dynamic dispatch.
type IHashContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllKeyvalue() []IKeyvalueContext
	Keyvalue(i int) IKeyvalueContext

	// IsHashContext differentiates from other interfaces.
	IsHashContext()
}

type HashContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHashContext() *HashContext {
	var p = new(HashContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_hash
	return p
}

func InitEmptyHashContext(p *HashContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_hash
}

func (*HashContext) IsHashContext() {}

func NewHashContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HashContext {
	var p = new(HashContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_hash

	return p
}

func (s *HashContext) GetParser() antlr.Parser { return s.parser }

func (s *HashContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserLBRACE, 0)
}

func (s *HashContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserRBRACE, 0)
}

func (s *HashContext) AllKeyvalue() []IKeyvalueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKeyvalueContext); ok {
			len++
		}
	}

	tst := make([]IKeyvalueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKeyvalueContext); ok {
			tst[i] = t.(IKeyvalueContext)
			i++
		}
	}

	return tst
}

func (s *HashContext) Keyvalue(i int) IKeyvalueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyvalueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyvalueContext)
}

func (s *HashContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HashContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HashContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterHash(s)
	}
}

func (s *HashContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitHash(s)
	}
}

func (p *ChronicleLogstashParser) Hash() (localctx IHashContext) {
	localctx = NewHashContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ChronicleLogstashParserRULE_hash)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Match(ChronicleLogstashParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ChronicleLogstashParserSTRING || _la == ChronicleLogstashParserID {
		{
			p.SetState(185)
			p.Keyvalue()
		}

		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(191)
		p.Match(ChronicleLogstashParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	AllList_value() []IList_valueContext
	List_value(i int) IList_valueContext

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_list
	return p
}

func InitEmptyListContext(p *ListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_list
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserLBRACKET, 0)
}

func (s *ListContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserRBRACKET, 0)
}

func (s *ListContext) AllList_value() []IList_valueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IList_valueContext); ok {
			len++
		}
	}

	tst := make([]IList_valueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IList_valueContext); ok {
			tst[i] = t.(IList_valueContext)
			i++
		}
	}

	return tst
}

func (s *ListContext) List_value(i int) IList_valueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_valueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_valueContext)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitList(s)
	}
}

func (p *ChronicleLogstashParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ChronicleLogstashParserRULE_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.Match(ChronicleLogstashParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&27380482048) != 0 {
		{
			p.SetState(194)
			p.List_value()
		}
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&27380482048) != 0 {
			{
				p.SetState(195)
				p.List_value()
			}

			p.SetState(200)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(203)
		p.Match(ChronicleLogstashParserRBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_listContext is an interface to support dynamic dispatch.
type IIf_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	STRING() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	AllList_value() []IList_valueContext
	List_value(i int) IList_valueContext

	// IsIf_listContext differentiates from other interfaces.
	IsIf_listContext()
}

type If_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_listContext() *If_listContext {
	var p = new(If_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_if_list
	return p
}

func InitEmptyIf_listContext(p *If_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_if_list
}

func (*If_listContext) IsIf_listContext() {}

func NewIf_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_listContext {
	var p = new(If_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_if_list

	return p
}

func (s *If_listContext) GetParser() antlr.Parser { return s.parser }

func (s *If_listContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserLBRACKET, 0)
}

func (s *If_listContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserRBRACKET, 0)
}

func (s *If_listContext) STRING() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserSTRING, 0)
}

func (s *If_listContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserBOOLEAN, 0)
}

func (s *If_listContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserNUMBER, 0)
}

func (s *If_listContext) AllList_value() []IList_valueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IList_valueContext); ok {
			len++
		}
	}

	tst := make([]IList_valueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IList_valueContext); ok {
			tst[i] = t.(IList_valueContext)
			i++
		}
	}

	return tst
}

func (s *If_listContext) List_value(i int) IList_valueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_valueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_valueContext)
}

func (s *If_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterIf_list(s)
	}
}

func (s *If_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitIf_list(s)
	}
}

func (p *ChronicleLogstashParser) If_list() (localctx IIf_listContext) {
	localctx = NewIf_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ChronicleLogstashParserRULE_if_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(ChronicleLogstashParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18790481920) != 0 {
		{
			p.SetState(206)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18790481920) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&27380482048) != 0 {
			{
				p.SetState(207)
				p.List_value()
			}

			p.SetState(212)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(215)
		p.Match(ChronicleLogstashParserRBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IList_valueContext is an interface to support dynamic dispatch.
type IList_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	ID() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	COMMA() antlr.TerminalNode

	// IsList_valueContext differentiates from other interfaces.
	IsList_valueContext()
}

type List_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_valueContext() *List_valueContext {
	var p = new(List_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_list_value
	return p
}

func InitEmptyList_valueContext(p *List_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_list_value
}

func (*List_valueContext) IsList_valueContext() {}

func NewList_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_valueContext {
	var p = new(List_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_list_value

	return p
}

func (s *List_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *List_valueContext) STRING() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserSTRING, 0)
}

func (s *List_valueContext) ID() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserID, 0)
}

func (s *List_valueContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserBOOLEAN, 0)
}

func (s *List_valueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserNUMBER, 0)
}

func (s *List_valueContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserCOMMA, 0)
}

func (s *List_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterList_value(s)
	}
}

func (s *List_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitList_value(s)
	}
}

func (p *ChronicleLogstashParser) List_value() (localctx IList_valueContext) {
	localctx = NewList_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ChronicleLogstashParserRULE_list_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&27380482048) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *ChronicleLogstashParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *StatementContext = nil
		if localctx != nil {
			t = localctx.(*StatementContext)
		}
		return p.Statement_Sempred(t, predIndex)

	case 7:
		var t *Math_statementContext = nil
		if localctx != nil {
			t = localctx.(*Math_statementContext)
		}
		return p.Math_statement_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ChronicleLogstashParser) Statement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ChronicleLogstashParser) Math_statement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
