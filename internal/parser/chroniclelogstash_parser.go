// Code generated from ChronicleLogstashParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

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
		"", "", "", "'if'", "'else if'", "'else'", "'for'", "", "'not'", "",
		"'}'", "'['", "']'", "'('", "')'", "", "','", "'!'", "", "'=='", "'!='",
		"'<'", "'>'", "'<='", "'>='", "'=~'", "'!~'",
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
		"filterblock", "conditionalblock", "for_var", "for_iterable", "statement",
		"expression", "unary_expression", "binary_expression", "expression_val",
		"signed_number", "paren_list", "paren_value", "math_statement", "math_expression",
		"boolean_op", "boolean_eval", "plugin", "keyvalue", "kv_lvalue", "kv_rvalue",
		"hash", "list", "if_list", "list_value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 39, 282, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 53,
		8, 0, 10, 0, 12, 0, 56, 9, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 5, 1, 66, 8, 1, 10, 1, 12, 1, 69, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 5, 1, 77, 8, 1, 10, 1, 12, 1, 80, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 3, 1, 87, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 94, 8, 1, 10,
		1, 12, 1, 97, 9, 1, 1, 1, 1, 1, 3, 1, 101, 8, 1, 1, 2, 1, 2, 1, 3, 4, 3,
		106, 8, 3, 11, 3, 12, 3, 107, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 124, 8, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 5, 4, 130, 8, 4, 10, 4, 12, 4, 133, 9, 4, 1, 5, 1, 5, 3, 5,
		137, 8, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 4, 7, 145, 8, 7, 11, 7,
		12, 7, 146, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		3, 8, 159, 8, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 168,
		8, 10, 10, 10, 12, 10, 171, 9, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 183, 8, 11, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 3, 12, 191, 8, 12, 1, 12, 1, 12, 1, 12, 5, 12, 196,
		8, 12, 10, 12, 12, 12, 199, 9, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 3, 15, 218, 8, 15, 1, 16, 1, 16, 1, 16, 5, 16, 223, 8, 16, 10, 16,
		12, 16, 226, 9, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 234,
		8, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3,
		19, 245, 8, 19, 1, 20, 1, 20, 5, 20, 249, 8, 20, 10, 20, 12, 20, 252, 9,
		20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 5, 21, 259, 8, 21, 10, 21, 12, 21,
		262, 9, 21, 3, 21, 264, 8, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 5, 22,
		271, 8, 22, 10, 22, 12, 22, 274, 9, 22, 3, 22, 276, 8, 22, 1, 22, 1, 22,
		1, 23, 1, 23, 1, 23, 0, 2, 8, 24, 24, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 0, 9, 1, 0, 3,
		4, 1, 0, 37, 38, 2, 0, 36, 36, 38, 38, 2, 0, 29, 29, 32, 32, 2, 0, 32,
		32, 34, 34, 1, 0, 27, 28, 2, 0, 30, 30, 33, 33, 2, 0, 29, 30, 34, 34, 3,
		0, 16, 16, 29, 30, 32, 34, 317, 0, 48, 1, 0, 0, 0, 2, 100, 1, 0, 0, 0,
		4, 102, 1, 0, 0, 0, 6, 105, 1, 0, 0, 0, 8, 123, 1, 0, 0, 0, 10, 136, 1,
		0, 0, 0, 12, 138, 1, 0, 0, 0, 14, 140, 1, 0, 0, 0, 16, 158, 1, 0, 0, 0,
		18, 160, 1, 0, 0, 0, 20, 163, 1, 0, 0, 0, 22, 182, 1, 0, 0, 0, 24, 190,
		1, 0, 0, 0, 26, 200, 1, 0, 0, 0, 28, 204, 1, 0, 0, 0, 30, 217, 1, 0, 0,
		0, 32, 219, 1, 0, 0, 0, 34, 229, 1, 0, 0, 0, 36, 235, 1, 0, 0, 0, 38, 244,
		1, 0, 0, 0, 40, 246, 1, 0, 0, 0, 42, 255, 1, 0, 0, 0, 44, 267, 1, 0, 0,
		0, 46, 279, 1, 0, 0, 0, 48, 49, 5, 33, 0, 0, 49, 54, 5, 9, 0, 0, 50, 53,
		3, 32, 16, 0, 51, 53, 3, 2, 1, 0, 52, 50, 1, 0, 0, 0, 52, 51, 1, 0, 0,
		0, 53, 56, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 57,
		1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 57, 58, 5, 10, 0, 0, 58, 59, 5, 0, 0, 1,
		59, 1, 1, 0, 0, 0, 60, 61, 7, 0, 0, 0, 61, 62, 3, 8, 4, 0, 62, 67, 5, 9,
		0, 0, 63, 66, 3, 32, 16, 0, 64, 66, 3, 2, 1, 0, 65, 63, 1, 0, 0, 0, 65,
		64, 1, 0, 0, 0, 66, 69, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0,
		0, 68, 70, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 70, 71, 5, 10, 0, 0, 71, 101,
		1, 0, 0, 0, 72, 73, 5, 5, 0, 0, 73, 78, 5, 9, 0, 0, 74, 77, 3, 32, 16,
		0, 75, 77, 3, 2, 1, 0, 76, 74, 1, 0, 0, 0, 76, 75, 1, 0, 0, 0, 77, 80,
		1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 81, 1, 0, 0, 0,
		80, 78, 1, 0, 0, 0, 81, 101, 5, 10, 0, 0, 82, 83, 5, 6, 0, 0, 83, 86, 3,
		4, 2, 0, 84, 85, 5, 36, 0, 0, 85, 87, 3, 4, 2, 0, 86, 84, 1, 0, 0, 0, 86,
		87, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 89, 5, 37, 0, 0, 89, 90, 3, 6,
		3, 0, 90, 95, 5, 39, 0, 0, 91, 94, 3, 32, 16, 0, 92, 94, 3, 2, 1, 0, 93,
		91, 1, 0, 0, 0, 93, 92, 1, 0, 0, 0, 94, 97, 1, 0, 0, 0, 95, 93, 1, 0, 0,
		0, 95, 96, 1, 0, 0, 0, 96, 98, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 98, 99,
		5, 10, 0, 0, 99, 101, 1, 0, 0, 0, 100, 60, 1, 0, 0, 0, 100, 72, 1, 0, 0,
		0, 100, 82, 1, 0, 0, 0, 101, 3, 1, 0, 0, 0, 102, 103, 7, 1, 0, 0, 103,
		5, 1, 0, 0, 0, 104, 106, 7, 2, 0, 0, 105, 104, 1, 0, 0, 0, 106, 107, 1,
		0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 7, 1, 0, 0, 0,
		109, 110, 6, 4, -1, 0, 110, 111, 5, 13, 0, 0, 111, 112, 3, 8, 4, 0, 112,
		113, 5, 14, 0, 0, 113, 124, 1, 0, 0, 0, 114, 115, 5, 11, 0, 0, 115, 116,
		3, 8, 4, 0, 116, 117, 5, 12, 0, 0, 117, 124, 1, 0, 0, 0, 118, 119, 5, 17,
		0, 0, 119, 124, 3, 8, 4, 3, 120, 121, 5, 8, 0, 0, 121, 124, 3, 8, 4, 2,
		122, 124, 3, 10, 5, 0, 123, 109, 1, 0, 0, 0, 123, 114, 1, 0, 0, 0, 123,
		118, 1, 0, 0, 0, 123, 120, 1, 0, 0, 0, 123, 122, 1, 0, 0, 0, 124, 131,
		1, 0, 0, 0, 125, 126, 10, 4, 0, 0, 126, 127, 3, 28, 14, 0, 127, 128, 3,
		8, 4, 5, 128, 130, 1, 0, 0, 0, 129, 125, 1, 0, 0, 0, 130, 133, 1, 0, 0,
		0, 131, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 9, 1, 0, 0, 0, 133,
		131, 1, 0, 0, 0, 134, 137, 3, 14, 7, 0, 135, 137, 3, 12, 6, 0, 136, 134,
		1, 0, 0, 0, 136, 135, 1, 0, 0, 0, 137, 11, 1, 0, 0, 0, 138, 139, 7, 3,
		0, 0, 139, 13, 1, 0, 0, 0, 140, 144, 3, 16, 8, 0, 141, 142, 3, 30, 15,
		0, 142, 143, 3, 16, 8, 0, 143, 145, 1, 0, 0, 0, 144, 141, 1, 0, 0, 0, 145,
		146, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 15, 1,
		0, 0, 0, 148, 159, 3, 24, 12, 0, 149, 159, 3, 18, 9, 0, 150, 159, 5, 34,
		0, 0, 151, 159, 3, 44, 22, 0, 152, 159, 5, 32, 0, 0, 153, 159, 5, 30, 0,
		0, 154, 159, 5, 31, 0, 0, 155, 159, 5, 29, 0, 0, 156, 159, 5, 33, 0, 0,
		157, 159, 3, 20, 10, 0, 158, 148, 1, 0, 0, 0, 158, 149, 1, 0, 0, 0, 158,
		150, 1, 0, 0, 0, 158, 151, 1, 0, 0, 0, 158, 152, 1, 0, 0, 0, 158, 153,
		1, 0, 0, 0, 158, 154, 1, 0, 0, 0, 158, 155, 1, 0, 0, 0, 158, 156, 1, 0,
		0, 0, 158, 157, 1, 0, 0, 0, 159, 17, 1, 0, 0, 0, 160, 161, 5, 18, 0, 0,
		161, 162, 5, 34, 0, 0, 162, 19, 1, 0, 0, 0, 163, 164, 5, 13, 0, 0, 164,
		169, 3, 22, 11, 0, 165, 166, 5, 16, 0, 0, 166, 168, 3, 22, 11, 0, 167,
		165, 1, 0, 0, 0, 168, 171, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 170,
		1, 0, 0, 0, 170, 172, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 172, 173, 5, 14,
		0, 0, 173, 21, 1, 0, 0, 0, 174, 183, 5, 30, 0, 0, 175, 183, 5, 31, 0, 0,
		176, 183, 5, 29, 0, 0, 177, 183, 5, 34, 0, 0, 178, 183, 3, 18, 9, 0, 179,
		183, 5, 33, 0, 0, 180, 183, 5, 32, 0, 0, 181, 183, 3, 44, 22, 0, 182, 174,
		1, 0, 0, 0, 182, 175, 1, 0, 0, 0, 182, 176, 1, 0, 0, 0, 182, 177, 1, 0,
		0, 0, 182, 178, 1, 0, 0, 0, 182, 179, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0,
		182, 181, 1, 0, 0, 0, 183, 23, 1, 0, 0, 0, 184, 185, 6, 12, -1, 0, 185,
		186, 5, 13, 0, 0, 186, 187, 3, 24, 12, 0, 187, 188, 5, 14, 0, 0, 188, 191,
		1, 0, 0, 0, 189, 191, 3, 26, 13, 0, 190, 184, 1, 0, 0, 0, 190, 189, 1,
		0, 0, 0, 191, 197, 1, 0, 0, 0, 192, 193, 10, 2, 0, 0, 193, 194, 5, 18,
		0, 0, 194, 196, 3, 24, 12, 3, 195, 192, 1, 0, 0, 0, 196, 199, 1, 0, 0,
		0, 197, 195, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 25, 1, 0, 0, 0, 199,
		197, 1, 0, 0, 0, 200, 201, 7, 4, 0, 0, 201, 202, 5, 18, 0, 0, 202, 203,
		7, 4, 0, 0, 203, 27, 1, 0, 0, 0, 204, 205, 7, 5, 0, 0, 205, 29, 1, 0, 0,
		0, 206, 218, 5, 19, 0, 0, 207, 218, 5, 20, 0, 0, 208, 218, 5, 21, 0, 0,
		209, 218, 5, 22, 0, 0, 210, 218, 5, 23, 0, 0, 211, 218, 5, 24, 0, 0, 212,
		218, 5, 25, 0, 0, 213, 218, 5, 26, 0, 0, 214, 218, 5, 7, 0, 0, 215, 216,
		5, 8, 0, 0, 216, 218, 5, 7, 0, 0, 217, 206, 1, 0, 0, 0, 217, 207, 1, 0,
		0, 0, 217, 208, 1, 0, 0, 0, 217, 209, 1, 0, 0, 0, 217, 210, 1, 0, 0, 0,
		217, 211, 1, 0, 0, 0, 217, 212, 1, 0, 0, 0, 217, 213, 1, 0, 0, 0, 217,
		214, 1, 0, 0, 0, 217, 215, 1, 0, 0, 0, 218, 31, 1, 0, 0, 0, 219, 220, 5,
		33, 0, 0, 220, 224, 5, 9, 0, 0, 221, 223, 3, 34, 17, 0, 222, 221, 1, 0,
		0, 0, 223, 226, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0,
		225, 227, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 227, 228, 5, 10, 0, 0, 228,
		33, 1, 0, 0, 0, 229, 230, 3, 36, 18, 0, 230, 231, 5, 15, 0, 0, 231, 233,
		3, 38, 19, 0, 232, 234, 5, 16, 0, 0, 233, 232, 1, 0, 0, 0, 233, 234, 1,
		0, 0, 0, 234, 35, 1, 0, 0, 0, 235, 236, 7, 6, 0, 0, 236, 37, 1, 0, 0, 0,
		237, 245, 5, 34, 0, 0, 238, 245, 3, 42, 21, 0, 239, 245, 3, 40, 20, 0,
		240, 245, 5, 30, 0, 0, 241, 245, 5, 29, 0, 0, 242, 245, 5, 33, 0, 0, 243,
		245, 5, 32, 0, 0, 244, 237, 1, 0, 0, 0, 244, 238, 1, 0, 0, 0, 244, 239,
		1, 0, 0, 0, 244, 240, 1, 0, 0, 0, 244, 241, 1, 0, 0, 0, 244, 242, 1, 0,
		0, 0, 244, 243, 1, 0, 0, 0, 245, 39, 1, 0, 0, 0, 246, 250, 5, 9, 0, 0,
		247, 249, 3, 34, 17, 0, 248, 247, 1, 0, 0, 0, 249, 252, 1, 0, 0, 0, 250,
		248, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 253, 1, 0, 0, 0, 252, 250,
		1, 0, 0, 0, 253, 254, 5, 10, 0, 0, 254, 41, 1, 0, 0, 0, 255, 263, 5, 11,
		0, 0, 256, 260, 3, 46, 23, 0, 257, 259, 3, 46, 23, 0, 258, 257, 1, 0, 0,
		0, 259, 262, 1, 0, 0, 0, 260, 258, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261,
		264, 1, 0, 0, 0, 262, 260, 1, 0, 0, 0, 263, 256, 1, 0, 0, 0, 263, 264,
		1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 266, 5, 12, 0, 0, 266, 43, 1, 0,
		0, 0, 267, 275, 5, 11, 0, 0, 268, 272, 7, 7, 0, 0, 269, 271, 3, 46, 23,
		0, 270, 269, 1, 0, 0, 0, 271, 274, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 272,
		273, 1, 0, 0, 0, 273, 276, 1, 0, 0, 0, 274, 272, 1, 0, 0, 0, 275, 268,
		1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 278, 5, 12,
		0, 0, 278, 45, 1, 0, 0, 0, 279, 280, 7, 8, 0, 0, 280, 47, 1, 0, 0, 0, 29,
		52, 54, 65, 67, 76, 78, 86, 93, 95, 100, 107, 123, 131, 136, 146, 158,
		169, 182, 190, 197, 217, 224, 233, 244, 250, 260, 263, 272, 275,
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
	ChronicleLogstashParserRULE_for_var           = 2
	ChronicleLogstashParserRULE_for_iterable      = 3
	ChronicleLogstashParserRULE_statement         = 4
	ChronicleLogstashParserRULE_expression        = 5
	ChronicleLogstashParserRULE_unary_expression  = 6
	ChronicleLogstashParserRULE_binary_expression = 7
	ChronicleLogstashParserRULE_expression_val    = 8
	ChronicleLogstashParserRULE_signed_number     = 9
	ChronicleLogstashParserRULE_paren_list        = 10
	ChronicleLogstashParserRULE_paren_value       = 11
	ChronicleLogstashParserRULE_math_statement    = 12
	ChronicleLogstashParserRULE_math_expression   = 13
	ChronicleLogstashParserRULE_boolean_op        = 14
	ChronicleLogstashParserRULE_boolean_eval      = 15
	ChronicleLogstashParserRULE_plugin            = 16
	ChronicleLogstashParserRULE_keyvalue          = 17
	ChronicleLogstashParserRULE_kv_lvalue         = 18
	ChronicleLogstashParserRULE_kv_rvalue         = 19
	ChronicleLogstashParserRULE_hash              = 20
	ChronicleLogstashParserRULE_list              = 21
	ChronicleLogstashParserRULE_if_list           = 22
	ChronicleLogstashParserRULE_list_value        = 23
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
		p.SetState(48)
		p.Match(ChronicleLogstashParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(ChronicleLogstashParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8589934712) != 0 {
		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case ChronicleLogstashParserID:
			{
				p.SetState(50)
				p.Plugin()
			}

		case ChronicleLogstashParserIF, ChronicleLogstashParserELSEIF, ChronicleLogstashParserELSE, ChronicleLogstashParserFOR:
			{
				p.SetState(51)
				p.Conditionalblock()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(57)
		p.Match(ChronicleLogstashParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
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
	AllFor_var() []IFor_varContext
	For_var(i int) IFor_varContext
	FORIN() antlr.TerminalNode
	For_iterable() IFor_iterableContext
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

func (s *ConditionalblockContext) AllFor_var() []IFor_varContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFor_varContext); ok {
			len++
		}
	}

	tst := make([]IFor_varContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFor_varContext); ok {
			tst[i] = t.(IFor_varContext)
			i++
		}
	}

	return tst
}

func (s *ConditionalblockContext) For_var(i int) IFor_varContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_varContext); ok {
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

	return t.(IFor_varContext)
}

func (s *ConditionalblockContext) FORIN() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserFORIN, 0)
}

func (s *ConditionalblockContext) For_iterable() IFor_iterableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_iterableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_iterableContext)
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

	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ChronicleLogstashParserIF, ChronicleLogstashParserELSEIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ChronicleLogstashParserIF || _la == ChronicleLogstashParserELSEIF) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(61)
			p.statement(0)
		}
		{
			p.SetState(62)
			p.Match(ChronicleLogstashParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8589934712) != 0 {
			p.SetState(65)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case ChronicleLogstashParserID:
				{
					p.SetState(63)
					p.Plugin()
				}

			case ChronicleLogstashParserIF, ChronicleLogstashParserELSEIF, ChronicleLogstashParserELSE, ChronicleLogstashParserFOR:
				{
					p.SetState(64)
					p.Conditionalblock()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(69)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(70)
			p.Match(ChronicleLogstashParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserELSE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.Match(ChronicleLogstashParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.Match(ChronicleLogstashParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8589934712) != 0 {
			p.SetState(76)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case ChronicleLogstashParserID:
				{
					p.SetState(74)
					p.Plugin()
				}

			case ChronicleLogstashParserIF, ChronicleLogstashParserELSEIF, ChronicleLogstashParserELSE, ChronicleLogstashParserFOR:
				{
					p.SetState(75)
					p.Conditionalblock()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(80)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(81)
			p.Match(ChronicleLogstashParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserFOR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(82)
			p.Match(ChronicleLogstashParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.For_var()
		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ChronicleLogstashParserFORCOMMA {
			{
				p.SetState(84)
				p.Match(ChronicleLogstashParserFORCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(85)
				p.For_var()
			}

		}
		{
			p.SetState(88)
			p.Match(ChronicleLogstashParserFORIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)
			p.For_iterable()
		}
		{
			p.SetState(90)
			p.Match(ChronicleLogstashParserFOROPENER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8589934712) != 0 {
			p.SetState(93)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case ChronicleLogstashParserID:
				{
					p.SetState(91)
					p.Plugin()
				}

			case ChronicleLogstashParserIF, ChronicleLogstashParserELSEIF, ChronicleLogstashParserELSE, ChronicleLogstashParserFOR:
				{
					p.SetState(92)
					p.Conditionalblock()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(97)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(98)
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

// IFor_varContext is an interface to support dynamic dispatch.
type IFor_varContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FORID() antlr.TerminalNode
	FORIN() antlr.TerminalNode

	// IsFor_varContext differentiates from other interfaces.
	IsFor_varContext()
}

type For_varContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_varContext() *For_varContext {
	var p = new(For_varContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_for_var
	return p
}

func InitEmptyFor_varContext(p *For_varContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_for_var
}

func (*For_varContext) IsFor_varContext() {}

func NewFor_varContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_varContext {
	var p = new(For_varContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_for_var

	return p
}

func (s *For_varContext) GetParser() antlr.Parser { return s.parser }

func (s *For_varContext) FORID() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserFORID, 0)
}

func (s *For_varContext) FORIN() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserFORIN, 0)
}

func (s *For_varContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_varContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_varContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterFor_var(s)
	}
}

func (s *For_varContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitFor_var(s)
	}
}

func (p *ChronicleLogstashParser) For_var() (localctx IFor_varContext) {
	localctx = NewFor_varContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ChronicleLogstashParserRULE_for_var)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ChronicleLogstashParserFORIN || _la == ChronicleLogstashParserFORID) {
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

// IFor_iterableContext is an interface to support dynamic dispatch.
type IFor_iterableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFORID() []antlr.TerminalNode
	FORID(i int) antlr.TerminalNode
	AllFORCOMMA() []antlr.TerminalNode
	FORCOMMA(i int) antlr.TerminalNode

	// IsFor_iterableContext differentiates from other interfaces.
	IsFor_iterableContext()
}

type For_iterableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_iterableContext() *For_iterableContext {
	var p = new(For_iterableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_for_iterable
	return p
}

func InitEmptyFor_iterableContext(p *For_iterableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_for_iterable
}

func (*For_iterableContext) IsFor_iterableContext() {}

func NewFor_iterableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_iterableContext {
	var p = new(For_iterableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_for_iterable

	return p
}

func (s *For_iterableContext) GetParser() antlr.Parser { return s.parser }

func (s *For_iterableContext) AllFORID() []antlr.TerminalNode {
	return s.GetTokens(ChronicleLogstashParserFORID)
}

func (s *For_iterableContext) FORID(i int) antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserFORID, i)
}

func (s *For_iterableContext) AllFORCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ChronicleLogstashParserFORCOMMA)
}

func (s *For_iterableContext) FORCOMMA(i int) antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserFORCOMMA, i)
}

func (s *For_iterableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_iterableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_iterableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterFor_iterable(s)
	}
}

func (s *For_iterableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitFor_iterable(s)
	}
}

func (p *ChronicleLogstashParser) For_iterable() (localctx IFor_iterableContext) {
	localctx = NewFor_iterableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ChronicleLogstashParserRULE_for_iterable)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ChronicleLogstashParserFORCOMMA || _la == ChronicleLogstashParserFORID {
		{
			p.SetState(104)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ChronicleLogstashParserFORCOMMA || _la == ChronicleLogstashParserFORID) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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
	NOT() antlr.TerminalNode
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

func (s *StatementContext) NOT() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserNOT, 0)
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
	_startState := 8
	p.EnterRecursionRule(localctx, 8, ChronicleLogstashParserRULE_statement, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(110)
			p.Match(ChronicleLogstashParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.statement(0)
		}
		{
			p.SetState(112)
			p.Match(ChronicleLogstashParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(114)
			p.Match(ChronicleLogstashParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)
			p.statement(0)
		}
		{
			p.SetState(116)
			p.Match(ChronicleLogstashParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(118)
			p.Match(ChronicleLogstashParserBOOLNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)
			p.statement(3)
		}

	case 4:
		{
			p.SetState(120)
			p.Match(ChronicleLogstashParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.statement(2)
		}

	case 5:
		{
			p.SetState(122)
			p.Expression()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
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
			p.SetState(125)

			if !(p.Precpred(p.GetParserRuleContext(), 4)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				goto errorExit
			}
			{
				p.SetState(126)
				p.Boolean_op()
			}
			{
				p.SetState(127)
				p.statement(5)
			}

		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 10, ChronicleLogstashParserRULE_expression)
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.Binary_expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
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
	p.EnterRule(localctx, 12, ChronicleLogstashParserRULE_unary_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
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
	AllBoolean_eval() []IBoolean_evalContext
	Boolean_eval(i int) IBoolean_evalContext

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

func (s *Binary_expressionContext) AllBoolean_eval() []IBoolean_evalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoolean_evalContext); ok {
			len++
		}
	}

	tst := make([]IBoolean_evalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoolean_evalContext); ok {
			tst[i] = t.(IBoolean_evalContext)
			i++
		}
	}

	return tst
}

func (s *Binary_expressionContext) Boolean_eval(i int) IBoolean_evalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_evalContext); ok {
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
	p.EnterRule(localctx, 14, ChronicleLogstashParserRULE_binary_expression)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Expression_val()
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(141)
				p.Boolean_eval()
			}
			{
				p.SetState(142)
				p.Expression_val()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
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
	Signed_number() ISigned_numberContext
	NUMBER() antlr.TerminalNode
	If_list() IIf_listContext
	IFSTATEMENTID() antlr.TerminalNode
	STRING() antlr.TerminalNode
	REGEX() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	ID() antlr.TerminalNode
	Paren_list() IParen_listContext

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

func (s *Expression_valContext) Signed_number() ISigned_numberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISigned_numberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISigned_numberContext)
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

func (s *Expression_valContext) Paren_list() IParen_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParen_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParen_listContext)
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
	p.EnterRule(localctx, 16, ChronicleLogstashParserRULE_expression_val)
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)
			p.math_statement(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(149)
			p.Signed_number()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(150)
			p.Match(ChronicleLogstashParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(151)
			p.If_list()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(152)
			p.Match(ChronicleLogstashParserIFSTATEMENTID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(153)
			p.Match(ChronicleLogstashParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(154)
			p.Match(ChronicleLogstashParserREGEX)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(155)
			p.Match(ChronicleLogstashParserBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(156)
			p.Match(ChronicleLogstashParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(157)
			p.Paren_list()
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

// ISigned_numberContext is an interface to support dynamic dispatch.
type ISigned_numberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MATHOP() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsSigned_numberContext differentiates from other interfaces.
	IsSigned_numberContext()
}

type Signed_numberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySigned_numberContext() *Signed_numberContext {
	var p = new(Signed_numberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_signed_number
	return p
}

func InitEmptySigned_numberContext(p *Signed_numberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_signed_number
}

func (*Signed_numberContext) IsSigned_numberContext() {}

func NewSigned_numberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Signed_numberContext {
	var p = new(Signed_numberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_signed_number

	return p
}

func (s *Signed_numberContext) GetParser() antlr.Parser { return s.parser }

func (s *Signed_numberContext) MATHOP() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserMATHOP, 0)
}

func (s *Signed_numberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserNUMBER, 0)
}

func (s *Signed_numberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Signed_numberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Signed_numberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterSigned_number(s)
	}
}

func (s *Signed_numberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitSigned_number(s)
	}
}

func (p *ChronicleLogstashParser) Signed_number() (localctx ISigned_numberContext) {
	localctx = NewSigned_numberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ChronicleLogstashParserRULE_signed_number)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(ChronicleLogstashParserMATHOP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		p.Match(ChronicleLogstashParserNUMBER)
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

// IParen_listContext is an interface to support dynamic dispatch.
type IParen_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	AllParen_value() []IParen_valueContext
	Paren_value(i int) IParen_valueContext
	RPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParen_listContext differentiates from other interfaces.
	IsParen_listContext()
}

type Paren_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParen_listContext() *Paren_listContext {
	var p = new(Paren_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_paren_list
	return p
}

func InitEmptyParen_listContext(p *Paren_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_paren_list
}

func (*Paren_listContext) IsParen_listContext() {}

func NewParen_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Paren_listContext {
	var p = new(Paren_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_paren_list

	return p
}

func (s *Paren_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Paren_listContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserLPAREN, 0)
}

func (s *Paren_listContext) AllParen_value() []IParen_valueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParen_valueContext); ok {
			len++
		}
	}

	tst := make([]IParen_valueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParen_valueContext); ok {
			tst[i] = t.(IParen_valueContext)
			i++
		}
	}

	return tst
}

func (s *Paren_listContext) Paren_value(i int) IParen_valueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParen_valueContext); ok {
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

	return t.(IParen_valueContext)
}

func (s *Paren_listContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserRPAREN, 0)
}

func (s *Paren_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ChronicleLogstashParserCOMMA)
}

func (s *Paren_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserCOMMA, i)
}

func (s *Paren_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Paren_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Paren_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterParen_list(s)
	}
}

func (s *Paren_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitParen_list(s)
	}
}

func (p *ChronicleLogstashParser) Paren_list() (localctx IParen_listContext) {
	localctx = NewParen_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ChronicleLogstashParserRULE_paren_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(ChronicleLogstashParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(164)
		p.Paren_value()
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ChronicleLogstashParserCOMMA {
		{
			p.SetState(165)
			p.Match(ChronicleLogstashParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.Paren_value()
		}

		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(172)
		p.Match(ChronicleLogstashParserRPAREN)
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

// IParen_valueContext is an interface to support dynamic dispatch.
type IParen_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	REGEX() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	Signed_number() ISigned_numberContext
	ID() antlr.TerminalNode
	IFSTATEMENTID() antlr.TerminalNode
	If_list() IIf_listContext

	// IsParen_valueContext differentiates from other interfaces.
	IsParen_valueContext()
}

type Paren_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParen_valueContext() *Paren_valueContext {
	var p = new(Paren_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_paren_value
	return p
}

func InitEmptyParen_valueContext(p *Paren_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ChronicleLogstashParserRULE_paren_value
}

func (*Paren_valueContext) IsParen_valueContext() {}

func NewParen_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Paren_valueContext {
	var p = new(Paren_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChronicleLogstashParserRULE_paren_value

	return p
}

func (s *Paren_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Paren_valueContext) STRING() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserSTRING, 0)
}

func (s *Paren_valueContext) REGEX() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserREGEX, 0)
}

func (s *Paren_valueContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserBOOLEAN, 0)
}

func (s *Paren_valueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserNUMBER, 0)
}

func (s *Paren_valueContext) Signed_number() ISigned_numberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISigned_numberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISigned_numberContext)
}

func (s *Paren_valueContext) ID() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserID, 0)
}

func (s *Paren_valueContext) IFSTATEMENTID() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserIFSTATEMENTID, 0)
}

func (s *Paren_valueContext) If_list() IIf_listContext {
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

func (s *Paren_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Paren_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Paren_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.EnterParen_value(s)
	}
}

func (s *Paren_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChronicleLogstashParserListener); ok {
		listenerT.ExitParen_value(s)
	}
}

func (p *ChronicleLogstashParser) Paren_value() (localctx IParen_valueContext) {
	localctx = NewParen_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ChronicleLogstashParserRULE_paren_value)
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ChronicleLogstashParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(174)
			p.Match(ChronicleLogstashParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserREGEX:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(175)
			p.Match(ChronicleLogstashParserREGEX)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserBOOLEAN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(176)
			p.Match(ChronicleLogstashParserBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserNUMBER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(177)
			p.Match(ChronicleLogstashParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserMATHOP:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(178)
			p.Signed_number()
		}

	case ChronicleLogstashParserID:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(179)
			p.Match(ChronicleLogstashParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserIFSTATEMENTID:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(180)
			p.Match(ChronicleLogstashParserIFSTATEMENTID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserLBRACKET:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(181)
			p.If_list()
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
	_startState := 24
	p.EnterRecursionRule(localctx, 24, ChronicleLogstashParserRULE_math_statement, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ChronicleLogstashParserLPAREN:
		{
			p.SetState(185)
			p.Match(ChronicleLogstashParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.math_statement(0)
		}
		{
			p.SetState(187)
			p.Match(ChronicleLogstashParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserIFSTATEMENTID, ChronicleLogstashParserNUMBER:
		{
			p.SetState(189)
			p.Math_expression()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
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
			p.SetState(192)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(193)
				p.Match(ChronicleLogstashParserMATHOP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(194)
				p.math_statement(3)
			}

		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 26, ChronicleLogstashParserRULE_math_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ChronicleLogstashParserIFSTATEMENTID || _la == ChronicleLogstashParserNUMBER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(201)
		p.Match(ChronicleLogstashParserMATHOP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
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
	p.EnterRule(localctx, 28, ChronicleLogstashParserRULE_boolean_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
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
	NOT() antlr.TerminalNode

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

func (s *Boolean_evalContext) NOT() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserNOT, 0)
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
	p.EnterRule(localctx, 30, ChronicleLogstashParserRULE_boolean_eval)
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ChronicleLogstashParserEQUAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(206)
			p.Match(ChronicleLogstashParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserNOTEQUAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(207)
			p.Match(ChronicleLogstashParserNOTEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserLESSTHAN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(208)
			p.Match(ChronicleLogstashParserLESSTHAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserGREATERTHAN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(209)
			p.Match(ChronicleLogstashParserGREATERTHAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserLTEQUAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(210)
			p.Match(ChronicleLogstashParserLTEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserGTEQUAL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(211)
			p.Match(ChronicleLogstashParserGTEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserMATCH:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(212)
			p.Match(ChronicleLogstashParserMATCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserNOTMATCH:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(213)
			p.Match(ChronicleLogstashParserNOTMATCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserIN:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(214)
			p.Match(ChronicleLogstashParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserNOT:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(215)
			p.Match(ChronicleLogstashParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
			p.Match(ChronicleLogstashParserIN)
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
	p.EnterRule(localctx, 32, ChronicleLogstashParserRULE_plugin)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(219)
		p.Match(ChronicleLogstashParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(220)
		p.Match(ChronicleLogstashParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ChronicleLogstashParserSTRING || _la == ChronicleLogstashParserID {
		{
			p.SetState(221)
			p.Keyvalue()
		}

		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(227)
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
	p.EnterRule(localctx, 34, ChronicleLogstashParserRULE_keyvalue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.Kv_lvalue()
	}
	{
		p.SetState(230)
		p.Match(ChronicleLogstashParserKVSEPARATOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(231)
		p.Kv_rvalue()
	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ChronicleLogstashParserCOMMA {
		{
			p.SetState(232)
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
	p.EnterRule(localctx, 36, ChronicleLogstashParserRULE_kv_lvalue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
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
	IFSTATEMENTID() antlr.TerminalNode

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

func (s *Kv_rvalueContext) IFSTATEMENTID() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserIFSTATEMENTID, 0)
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
	p.EnterRule(localctx, 38, ChronicleLogstashParserRULE_kv_rvalue)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ChronicleLogstashParserNUMBER:
		{
			p.SetState(237)
			p.Match(ChronicleLogstashParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserLBRACKET:
		{
			p.SetState(238)
			p.List()
		}

	case ChronicleLogstashParserLBRACE:
		{
			p.SetState(239)
			p.Hash()
		}

	case ChronicleLogstashParserSTRING:
		{
			p.SetState(240)
			p.Match(ChronicleLogstashParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserBOOLEAN:
		{
			p.SetState(241)
			p.Match(ChronicleLogstashParserBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserID:
		{
			p.SetState(242)
			p.Match(ChronicleLogstashParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ChronicleLogstashParserIFSTATEMENTID:
		{
			p.SetState(243)
			p.Match(ChronicleLogstashParserIFSTATEMENTID)
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
	p.EnterRule(localctx, 40, ChronicleLogstashParserRULE_hash)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.Match(ChronicleLogstashParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ChronicleLogstashParserSTRING || _la == ChronicleLogstashParserID {
		{
			p.SetState(247)
			p.Keyvalue()
		}

		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(253)
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
	p.EnterRule(localctx, 42, ChronicleLogstashParserRULE_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(255)
		p.Match(ChronicleLogstashParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31675449344) != 0 {
		{
			p.SetState(256)
			p.List_value()
		}
		p.SetState(260)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31675449344) != 0 {
			{
				p.SetState(257)
				p.List_value()
			}

			p.SetState(262)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(265)
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
	p.EnterRule(localctx, 44, ChronicleLogstashParserRULE_if_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		p.Match(ChronicleLogstashParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18790481920) != 0 {
		{
			p.SetState(268)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18790481920) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31675449344) != 0 {
			{
				p.SetState(269)
				p.List_value()
			}

			p.SetState(274)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(277)
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
	IFSTATEMENTID() antlr.TerminalNode
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

func (s *List_valueContext) IFSTATEMENTID() antlr.TerminalNode {
	return s.GetToken(ChronicleLogstashParserIFSTATEMENTID, 0)
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
	p.EnterRule(localctx, 46, ChronicleLogstashParserRULE_list_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31675449344) != 0) {
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
	case 4:
		var t *StatementContext = nil
		if localctx != nil {
			t = localctx.(*StatementContext)
		}
		return p.Statement_Sempred(t, predIndex)

	case 12:
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
		return p.Precpred(p.GetParserRuleContext(), 4)

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
