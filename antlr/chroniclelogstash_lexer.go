// Code generated from ChronicleLogstashLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ChronicleLogstashLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ChronicleLogstashLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func chroniclelogstashlexerLexerInit() {
	staticData := &ChronicleLogstashLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE", "FORMODE",
	}
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
		"WS", "COMMENT", "IF", "ELSEIF", "ELSE", "FOR", "IN", "NOT", "TRUE",
		"FALSE", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET", "LPAREN", "RPAREN",
		"KVSEPARATOR", "COMMA", "BOOLNOT", "MATHOP", "EQUAL", "NOTEQUAL", "LESSTHAN",
		"GREATERTHAN", "LTEQUAL", "GTEQUAL", "MATCH", "NOTMATCH", "AND", "OR",
		"BOOLEAN", "STRING", "STRING_DOUBLE", "STRING_SINGLE", "ESC", "REGEX",
		"IFSTATEMENTID", "ID", "NUMBER", "DIGIT", "LETTERS", "FORWS", "FORCOMMA",
		"FORIN", "FORID", "FOROPENER",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 39, 329, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3,
		7, 3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9,
		7, 9, 2, 10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7,
		14, 2, 15, 7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19,
		2, 20, 7, 20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2,
		25, 7, 25, 2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30,
		7, 30, 2, 31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7,
		35, 2, 36, 7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40,
		2, 41, 7, 41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 1,
		0, 4, 0, 96, 8, 0, 11, 0, 12, 0, 97, 1, 0, 1, 0, 1, 1, 1, 1, 5, 1, 104,
		8, 1, 10, 1, 12, 1, 107, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 3, 16, 166, 8, 16, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26,
		1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 201, 8,
		28, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 207, 8, 29, 1, 30, 1, 30, 3, 30,
		211, 8, 30, 1, 31, 1, 31, 3, 31, 215, 8, 31, 1, 32, 1, 32, 1, 32, 5, 32,
		220, 8, 32, 10, 32, 12, 32, 223, 9, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1,
		33, 5, 33, 230, 8, 33, 10, 33, 12, 33, 233, 9, 33, 1, 33, 1, 33, 1, 34,
		1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 5, 35, 243, 8, 35, 10, 35, 12, 35, 246,
		9, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 5, 35, 254, 8, 35, 10,
		35, 12, 35, 257, 9, 35, 1, 35, 1, 35, 3, 35, 261, 8, 35, 1, 36, 1, 36,
		1, 36, 4, 36, 266, 8, 36, 11, 36, 12, 36, 267, 1, 36, 4, 36, 271, 8, 36,
		11, 36, 12, 36, 272, 1, 36, 1, 36, 4, 36, 277, 8, 36, 11, 36, 12, 36, 278,
		1, 37, 1, 37, 1, 37, 1, 37, 5, 37, 285, 8, 37, 10, 37, 12, 37, 288, 9,
		37, 1, 38, 5, 38, 291, 8, 38, 10, 38, 12, 38, 294, 9, 38, 1, 38, 1, 38,
		4, 38, 298, 8, 38, 11, 38, 12, 38, 299, 1, 38, 4, 38, 303, 8, 38, 11, 38,
		12, 38, 304, 3, 38, 307, 8, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41,
		1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 1, 44, 4, 44, 322, 8, 44, 11,
		44, 12, 44, 323, 1, 45, 1, 45, 1, 45, 1, 45, 2, 244, 255, 0, 46, 2, 1,
		4, 2, 6, 3, 8, 4, 10, 5, 12, 6, 14, 7, 16, 8, 18, 0, 20, 0, 22, 9, 24,
		10, 26, 11, 28, 12, 30, 13, 32, 14, 34, 15, 36, 16, 38, 17, 40, 18, 42,
		19, 44, 20, 46, 21, 48, 22, 50, 23, 52, 24, 54, 25, 56, 26, 58, 27, 60,
		28, 62, 29, 64, 30, 66, 0, 68, 0, 70, 0, 72, 31, 74, 32, 76, 33, 78, 34,
		80, 0, 82, 0, 84, 35, 86, 36, 88, 37, 90, 38, 92, 39, 2, 0, 1, 14, 3, 0,
		9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 2, 0, 58, 58, 61, 61, 3, 0,
		42, 43, 45, 45, 47, 47, 4, 0, 10, 10, 13, 13, 34, 34, 92, 92, 4, 0, 10,
		10, 13, 13, 39, 39, 92, 92, 3, 0, 10, 10, 13, 13, 32, 32, 3, 0, 10, 10,
		13, 13, 47, 47, 3, 0, 45, 45, 64, 64, 95, 95, 1, 0, 95, 95, 1, 0, 46, 46,
		1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 6, 0, 9, 10, 13, 13, 32, 32, 44, 44,
		123, 123, 125, 125, 349, 0, 2, 1, 0, 0, 0, 0, 4, 1, 0, 0, 0, 0, 6, 1, 0,
		0, 0, 0, 8, 1, 0, 0, 0, 0, 10, 1, 0, 0, 0, 0, 12, 1, 0, 0, 0, 0, 14, 1,
		0, 0, 0, 0, 16, 1, 0, 0, 0, 0, 22, 1, 0, 0, 0, 0, 24, 1, 0, 0, 0, 0, 26,
		1, 0, 0, 0, 0, 28, 1, 0, 0, 0, 0, 30, 1, 0, 0, 0, 0, 32, 1, 0, 0, 0, 0,
		34, 1, 0, 0, 0, 0, 36, 1, 0, 0, 0, 0, 38, 1, 0, 0, 0, 0, 40, 1, 0, 0, 0,
		0, 42, 1, 0, 0, 0, 0, 44, 1, 0, 0, 0, 0, 46, 1, 0, 0, 0, 0, 48, 1, 0, 0,
		0, 0, 50, 1, 0, 0, 0, 0, 52, 1, 0, 0, 0, 0, 54, 1, 0, 0, 0, 0, 56, 1, 0,
		0, 0, 0, 58, 1, 0, 0, 0, 0, 60, 1, 0, 0, 0, 0, 62, 1, 0, 0, 0, 0, 64, 1,
		0, 0, 0, 0, 72, 1, 0, 0, 0, 0, 74, 1, 0, 0, 0, 0, 76, 1, 0, 0, 0, 0, 78,
		1, 0, 0, 0, 1, 84, 1, 0, 0, 0, 1, 86, 1, 0, 0, 0, 1, 88, 1, 0, 0, 0, 1,
		90, 1, 0, 0, 0, 1, 92, 1, 0, 0, 0, 2, 95, 1, 0, 0, 0, 4, 101, 1, 0, 0,
		0, 6, 110, 1, 0, 0, 0, 8, 113, 1, 0, 0, 0, 10, 121, 1, 0, 0, 0, 12, 126,
		1, 0, 0, 0, 14, 132, 1, 0, 0, 0, 16, 135, 1, 0, 0, 0, 18, 139, 1, 0, 0,
		0, 20, 144, 1, 0, 0, 0, 22, 150, 1, 0, 0, 0, 24, 152, 1, 0, 0, 0, 26, 154,
		1, 0, 0, 0, 28, 156, 1, 0, 0, 0, 30, 158, 1, 0, 0, 0, 32, 160, 1, 0, 0,
		0, 34, 165, 1, 0, 0, 0, 36, 167, 1, 0, 0, 0, 38, 169, 1, 0, 0, 0, 40, 171,
		1, 0, 0, 0, 42, 173, 1, 0, 0, 0, 44, 176, 1, 0, 0, 0, 46, 179, 1, 0, 0,
		0, 48, 181, 1, 0, 0, 0, 50, 183, 1, 0, 0, 0, 52, 186, 1, 0, 0, 0, 54, 189,
		1, 0, 0, 0, 56, 192, 1, 0, 0, 0, 58, 200, 1, 0, 0, 0, 60, 206, 1, 0, 0,
		0, 62, 210, 1, 0, 0, 0, 64, 214, 1, 0, 0, 0, 66, 216, 1, 0, 0, 0, 68, 226,
		1, 0, 0, 0, 70, 236, 1, 0, 0, 0, 72, 260, 1, 0, 0, 0, 74, 276, 1, 0, 0,
		0, 76, 280, 1, 0, 0, 0, 78, 306, 1, 0, 0, 0, 80, 308, 1, 0, 0, 0, 82, 310,
		1, 0, 0, 0, 84, 312, 1, 0, 0, 0, 86, 316, 1, 0, 0, 0, 88, 318, 1, 0, 0,
		0, 90, 321, 1, 0, 0, 0, 92, 325, 1, 0, 0, 0, 94, 96, 7, 0, 0, 0, 95, 94,
		1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0,
		98, 99, 1, 0, 0, 0, 99, 100, 6, 0, 0, 0, 100, 3, 1, 0, 0, 0, 101, 105,
		5, 35, 0, 0, 102, 104, 8, 1, 0, 0, 103, 102, 1, 0, 0, 0, 104, 107, 1, 0,
		0, 0, 105, 103, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 108, 1, 0, 0, 0,
		107, 105, 1, 0, 0, 0, 108, 109, 6, 1, 0, 0, 109, 5, 1, 0, 0, 0, 110, 111,
		5, 105, 0, 0, 111, 112, 5, 102, 0, 0, 112, 7, 1, 0, 0, 0, 113, 114, 5,
		101, 0, 0, 114, 115, 5, 108, 0, 0, 115, 116, 5, 115, 0, 0, 116, 117, 5,
		101, 0, 0, 117, 118, 5, 32, 0, 0, 118, 119, 5, 105, 0, 0, 119, 120, 5,
		102, 0, 0, 120, 9, 1, 0, 0, 0, 121, 122, 5, 101, 0, 0, 122, 123, 5, 108,
		0, 0, 123, 124, 5, 115, 0, 0, 124, 125, 5, 101, 0, 0, 125, 11, 1, 0, 0,
		0, 126, 127, 5, 102, 0, 0, 127, 128, 5, 111, 0, 0, 128, 129, 5, 114, 0,
		0, 129, 130, 1, 0, 0, 0, 130, 131, 6, 5, 1, 0, 131, 13, 1, 0, 0, 0, 132,
		133, 5, 105, 0, 0, 133, 134, 5, 110, 0, 0, 134, 15, 1, 0, 0, 0, 135, 136,
		5, 110, 0, 0, 136, 137, 5, 111, 0, 0, 137, 138, 5, 116, 0, 0, 138, 17,
		1, 0, 0, 0, 139, 140, 5, 116, 0, 0, 140, 141, 5, 114, 0, 0, 141, 142, 5,
		117, 0, 0, 142, 143, 5, 101, 0, 0, 143, 19, 1, 0, 0, 0, 144, 145, 5, 102,
		0, 0, 145, 146, 5, 97, 0, 0, 146, 147, 5, 108, 0, 0, 147, 148, 5, 115,
		0, 0, 148, 149, 5, 101, 0, 0, 149, 21, 1, 0, 0, 0, 150, 151, 5, 123, 0,
		0, 151, 23, 1, 0, 0, 0, 152, 153, 5, 125, 0, 0, 153, 25, 1, 0, 0, 0, 154,
		155, 5, 91, 0, 0, 155, 27, 1, 0, 0, 0, 156, 157, 5, 93, 0, 0, 157, 29,
		1, 0, 0, 0, 158, 159, 5, 40, 0, 0, 159, 31, 1, 0, 0, 0, 160, 161, 5, 41,
		0, 0, 161, 33, 1, 0, 0, 0, 162, 163, 5, 61, 0, 0, 163, 166, 5, 62, 0, 0,
		164, 166, 7, 2, 0, 0, 165, 162, 1, 0, 0, 0, 165, 164, 1, 0, 0, 0, 166,
		35, 1, 0, 0, 0, 167, 168, 5, 44, 0, 0, 168, 37, 1, 0, 0, 0, 169, 170, 5,
		33, 0, 0, 170, 39, 1, 0, 0, 0, 171, 172, 7, 3, 0, 0, 172, 41, 1, 0, 0,
		0, 173, 174, 5, 61, 0, 0, 174, 175, 5, 61, 0, 0, 175, 43, 1, 0, 0, 0, 176,
		177, 5, 33, 0, 0, 177, 178, 5, 61, 0, 0, 178, 45, 1, 0, 0, 0, 179, 180,
		5, 60, 0, 0, 180, 47, 1, 0, 0, 0, 181, 182, 5, 62, 0, 0, 182, 49, 1, 0,
		0, 0, 183, 184, 5, 60, 0, 0, 184, 185, 5, 61, 0, 0, 185, 51, 1, 0, 0, 0,
		186, 187, 5, 62, 0, 0, 187, 188, 5, 61, 0, 0, 188, 53, 1, 0, 0, 0, 189,
		190, 5, 61, 0, 0, 190, 191, 5, 126, 0, 0, 191, 55, 1, 0, 0, 0, 192, 193,
		5, 33, 0, 0, 193, 194, 5, 126, 0, 0, 194, 57, 1, 0, 0, 0, 195, 196, 5,
		97, 0, 0, 196, 197, 5, 110, 0, 0, 197, 201, 5, 100, 0, 0, 198, 199, 5,
		38, 0, 0, 199, 201, 5, 38, 0, 0, 200, 195, 1, 0, 0, 0, 200, 198, 1, 0,
		0, 0, 201, 59, 1, 0, 0, 0, 202, 203, 5, 111, 0, 0, 203, 207, 5, 114, 0,
		0, 204, 205, 5, 124, 0, 0, 205, 207, 5, 124, 0, 0, 206, 202, 1, 0, 0, 0,
		206, 204, 1, 0, 0, 0, 207, 61, 1, 0, 0, 0, 208, 211, 3, 18, 8, 0, 209,
		211, 3, 20, 9, 0, 210, 208, 1, 0, 0, 0, 210, 209, 1, 0, 0, 0, 211, 63,
		1, 0, 0, 0, 212, 215, 3, 66, 32, 0, 213, 215, 3, 68, 33, 0, 214, 212, 1,
		0, 0, 0, 214, 213, 1, 0, 0, 0, 215, 65, 1, 0, 0, 0, 216, 221, 5, 34, 0,
		0, 217, 220, 3, 70, 34, 0, 218, 220, 8, 4, 0, 0, 219, 217, 1, 0, 0, 0,
		219, 218, 1, 0, 0, 0, 220, 223, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 221,
		222, 1, 0, 0, 0, 222, 224, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 224, 225,
		5, 34, 0, 0, 225, 67, 1, 0, 0, 0, 226, 231, 5, 39, 0, 0, 227, 230, 3, 70,
		34, 0, 228, 230, 8, 5, 0, 0, 229, 227, 1, 0, 0, 0, 229, 228, 1, 0, 0, 0,
		230, 233, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232,
		234, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0, 234, 235, 5, 39, 0, 0, 235, 69,
		1, 0, 0, 0, 236, 237, 5, 92, 0, 0, 237, 238, 8, 6, 0, 0, 238, 71, 1, 0,
		0, 0, 239, 244, 5, 47, 0, 0, 240, 243, 3, 70, 34, 0, 241, 243, 8, 7, 0,
		0, 242, 240, 1, 0, 0, 0, 242, 241, 1, 0, 0, 0, 243, 246, 1, 0, 0, 0, 244,
		245, 1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 245, 247, 1, 0, 0, 0, 246, 244,
		1, 0, 0, 0, 247, 261, 5, 47, 0, 0, 248, 249, 5, 92, 0, 0, 249, 250, 5,
		47, 0, 0, 250, 255, 1, 0, 0, 0, 251, 254, 3, 70, 34, 0, 252, 254, 8, 7,
		0, 0, 253, 251, 1, 0, 0, 0, 253, 252, 1, 0, 0, 0, 254, 257, 1, 0, 0, 0,
		255, 256, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 256, 258, 1, 0, 0, 0, 257,
		255, 1, 0, 0, 0, 258, 259, 5, 92, 0, 0, 259, 261, 5, 47, 0, 0, 260, 239,
		1, 0, 0, 0, 260, 248, 1, 0, 0, 0, 261, 73, 1, 0, 0, 0, 262, 270, 3, 26,
		12, 0, 263, 271, 3, 82, 40, 0, 264, 266, 3, 80, 39, 0, 265, 264, 1, 0,
		0, 0, 266, 267, 1, 0, 0, 0, 267, 265, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0,
		268, 271, 1, 0, 0, 0, 269, 271, 7, 8, 0, 0, 270, 263, 1, 0, 0, 0, 270,
		265, 1, 0, 0, 0, 270, 269, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 270,
		1, 0, 0, 0, 272, 273, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 275, 3, 28,
		13, 0, 275, 277, 1, 0, 0, 0, 276, 262, 1, 0, 0, 0, 277, 278, 1, 0, 0, 0,
		278, 276, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 75, 1, 0, 0, 0, 280, 286,
		3, 82, 40, 0, 281, 285, 3, 82, 40, 0, 282, 285, 3, 80, 39, 0, 283, 285,
		7, 9, 0, 0, 284, 281, 1, 0, 0, 0, 284, 282, 1, 0, 0, 0, 284, 283, 1, 0,
		0, 0, 285, 288, 1, 0, 0, 0, 286, 284, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0,
		287, 77, 1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 289, 291, 3, 80, 39, 0, 290,
		289, 1, 0, 0, 0, 291, 294, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 292, 293,
		1, 0, 0, 0, 293, 295, 1, 0, 0, 0, 294, 292, 1, 0, 0, 0, 295, 297, 7, 10,
		0, 0, 296, 298, 3, 80, 39, 0, 297, 296, 1, 0, 0, 0, 298, 299, 1, 0, 0,
		0, 299, 297, 1, 0, 0, 0, 299, 300, 1, 0, 0, 0, 300, 307, 1, 0, 0, 0, 301,
		303, 3, 80, 39, 0, 302, 301, 1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 302,
		1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 307, 1, 0, 0, 0, 306, 292, 1, 0,
		0, 0, 306, 302, 1, 0, 0, 0, 307, 79, 1, 0, 0, 0, 308, 309, 7, 11, 0, 0,
		309, 81, 1, 0, 0, 0, 310, 311, 7, 12, 0, 0, 311, 83, 1, 0, 0, 0, 312, 313,
		3, 2, 0, 0, 313, 314, 1, 0, 0, 0, 314, 315, 6, 41, 0, 0, 315, 85, 1, 0,
		0, 0, 316, 317, 3, 36, 17, 0, 317, 87, 1, 0, 0, 0, 318, 319, 3, 14, 6,
		0, 319, 89, 1, 0, 0, 0, 320, 322, 8, 13, 0, 0, 321, 320, 1, 0, 0, 0, 322,
		323, 1, 0, 0, 0, 323, 321, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 91, 1,
		0, 0, 0, 325, 326, 5, 123, 0, 0, 326, 327, 1, 0, 0, 0, 327, 328, 6, 45,
		2, 0, 328, 93, 1, 0, 0, 0, 29, 0, 1, 97, 105, 165, 200, 206, 210, 214,
		219, 221, 229, 231, 242, 244, 253, 255, 260, 267, 270, 272, 278, 284, 286,
		292, 299, 304, 306, 323, 3, 6, 0, 0, 5, 1, 0, 4, 0, 0,
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

// ChronicleLogstashLexerInit initializes any static state used to implement ChronicleLogstashLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewChronicleLogstashLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ChronicleLogstashLexerInit() {
	staticData := &ChronicleLogstashLexerLexerStaticData
	staticData.once.Do(chroniclelogstashlexerLexerInit)
}

// NewChronicleLogstashLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewChronicleLogstashLexer(input antlr.CharStream) *ChronicleLogstashLexer {
	ChronicleLogstashLexerInit()
	l := new(ChronicleLogstashLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ChronicleLogstashLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "ChronicleLogstashLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ChronicleLogstashLexer tokens.
const (
	ChronicleLogstashLexerWS            = 1
	ChronicleLogstashLexerCOMMENT       = 2
	ChronicleLogstashLexerIF            = 3
	ChronicleLogstashLexerELSEIF        = 4
	ChronicleLogstashLexerELSE          = 5
	ChronicleLogstashLexerFOR           = 6
	ChronicleLogstashLexerIN            = 7
	ChronicleLogstashLexerNOT           = 8
	ChronicleLogstashLexerLBRACE        = 9
	ChronicleLogstashLexerRBRACE        = 10
	ChronicleLogstashLexerLBRACKET      = 11
	ChronicleLogstashLexerRBRACKET      = 12
	ChronicleLogstashLexerLPAREN        = 13
	ChronicleLogstashLexerRPAREN        = 14
	ChronicleLogstashLexerKVSEPARATOR   = 15
	ChronicleLogstashLexerCOMMA         = 16
	ChronicleLogstashLexerBOOLNOT       = 17
	ChronicleLogstashLexerMATHOP        = 18
	ChronicleLogstashLexerEQUAL         = 19
	ChronicleLogstashLexerNOTEQUAL      = 20
	ChronicleLogstashLexerLESSTHAN      = 21
	ChronicleLogstashLexerGREATERTHAN   = 22
	ChronicleLogstashLexerLTEQUAL       = 23
	ChronicleLogstashLexerGTEQUAL       = 24
	ChronicleLogstashLexerMATCH         = 25
	ChronicleLogstashLexerNOTMATCH      = 26
	ChronicleLogstashLexerAND           = 27
	ChronicleLogstashLexerOR            = 28
	ChronicleLogstashLexerBOOLEAN       = 29
	ChronicleLogstashLexerSTRING        = 30
	ChronicleLogstashLexerREGEX         = 31
	ChronicleLogstashLexerIFSTATEMENTID = 32
	ChronicleLogstashLexerID            = 33
	ChronicleLogstashLexerNUMBER        = 34
	ChronicleLogstashLexerFORWS         = 35
	ChronicleLogstashLexerFORCOMMA      = 36
	ChronicleLogstashLexerFORIN         = 37
	ChronicleLogstashLexerFORID         = 38
	ChronicleLogstashLexerFOROPENER     = 39
)

// ChronicleLogstashLexerFORMODE is the ChronicleLogstashLexer mode.
const ChronicleLogstashLexerFORMODE = 1
