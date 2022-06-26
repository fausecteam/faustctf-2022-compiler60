// Code generated from Algol60V2.g4 by ANTLR 4.10.1. DO NOT EDIT.

package algol60_parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "regexp"

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type Algol60V2Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var algol60v2lexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func algol60v2lexerLexerInit() {
	staticData := &algol60v2lexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "';'", "','", "'['", "']'", "'('", "')'", "':'", "'-'", "'+'", "'\\u00D7'",
		"'\\u00F7'", "'\\u2283'", "'\\u2228'", "'\\u2227'", "'\\u00AC'", "'<'",
		"'\\u2264'", "'='", "'\\u2265'", "'>'", "'\\u2260'", "", "''BEGIN''",
		"''END''", "''IF''", "''THEN''", "''ELSE''", "''FOR''", "''DO''", "''STEP''",
		"''UNTIL''", "''WHILE''", "''PROCEDURE''", "''INTEGER''", "''ARRAY''",
		"''STRING''",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "ASSIGN", "BEGIN", "END", "IF", "THEN", "ELSE",
		"FOR", "DO", "STEP", "UNTIL", "WHILE", "PROCEDURE", "INTEGER", "ARRAY",
		"STRING", "COMMENT", "QUOTED_STRING", "NUMBER", "IDENTIFIER", "SKIP_WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "ASSIGN", "BEGIN", "END", "IF",
		"THEN", "ELSE", "FOR", "DO", "STEP", "UNTIL", "WHILE", "PROCEDURE",
		"INTEGER", "ARRAY", "STRING", "COMMENT", "QUOTED_STRING", "NUMBER",
		"IDENTIFIER", "SKIP_WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 41, 292, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 5, 21, 128,
		8, 21, 10, 21, 12, 21, 131, 9, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 5, 36, 252,
		8, 36, 10, 36, 12, 36, 255, 9, 36, 1, 37, 1, 37, 1, 37, 1, 37, 5, 37, 261,
		8, 37, 10, 37, 12, 37, 264, 9, 37, 1, 37, 1, 37, 1, 38, 1, 38, 5, 38, 270,
		8, 38, 10, 38, 12, 38, 273, 9, 38, 1, 38, 1, 38, 1, 39, 1, 39, 5, 39, 279,
		8, 39, 10, 39, 12, 39, 282, 9, 39, 1, 39, 1, 39, 1, 40, 4, 40, 287, 8,
		40, 11, 40, 12, 40, 288, 1, 40, 1, 40, 0, 0, 41, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32,
		65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41,
		1, 0, 7, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 59, 59, 2, 0, 34, 34, 92, 92,
		1, 0, 48, 57, 4, 0, 9, 10, 13, 13, 32, 32, 48, 57, 2, 0, 65, 90, 97, 122,
		6, 0, 9, 10, 13, 13, 32, 32, 48, 57, 65, 90, 97, 122, 298, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0,
		0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0,
		0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1,
		0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71,
		1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0,
		79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 1, 83, 1, 0, 0, 0, 3, 85, 1, 0, 0, 0,
		5, 87, 1, 0, 0, 0, 7, 89, 1, 0, 0, 0, 9, 91, 1, 0, 0, 0, 11, 93, 1, 0,
		0, 0, 13, 95, 1, 0, 0, 0, 15, 97, 1, 0, 0, 0, 17, 99, 1, 0, 0, 0, 19, 101,
		1, 0, 0, 0, 21, 103, 1, 0, 0, 0, 23, 105, 1, 0, 0, 0, 25, 107, 1, 0, 0,
		0, 27, 109, 1, 0, 0, 0, 29, 111, 1, 0, 0, 0, 31, 113, 1, 0, 0, 0, 33, 115,
		1, 0, 0, 0, 35, 117, 1, 0, 0, 0, 37, 119, 1, 0, 0, 0, 39, 121, 1, 0, 0,
		0, 41, 123, 1, 0, 0, 0, 43, 125, 1, 0, 0, 0, 45, 134, 1, 0, 0, 0, 47, 142,
		1, 0, 0, 0, 49, 148, 1, 0, 0, 0, 51, 153, 1, 0, 0, 0, 53, 160, 1, 0, 0,
		0, 55, 167, 1, 0, 0, 0, 57, 173, 1, 0, 0, 0, 59, 178, 1, 0, 0, 0, 61, 185,
		1, 0, 0, 0, 63, 193, 1, 0, 0, 0, 65, 201, 1, 0, 0, 0, 67, 213, 1, 0, 0,
		0, 69, 223, 1, 0, 0, 0, 71, 231, 1, 0, 0, 0, 73, 240, 1, 0, 0, 0, 75, 256,
		1, 0, 0, 0, 77, 267, 1, 0, 0, 0, 79, 276, 1, 0, 0, 0, 81, 286, 1, 0, 0,
		0, 83, 84, 5, 59, 0, 0, 84, 2, 1, 0, 0, 0, 85, 86, 5, 44, 0, 0, 86, 4,
		1, 0, 0, 0, 87, 88, 5, 91, 0, 0, 88, 6, 1, 0, 0, 0, 89, 90, 5, 93, 0, 0,
		90, 8, 1, 0, 0, 0, 91, 92, 5, 40, 0, 0, 92, 10, 1, 0, 0, 0, 93, 94, 5,
		41, 0, 0, 94, 12, 1, 0, 0, 0, 95, 96, 5, 58, 0, 0, 96, 14, 1, 0, 0, 0,
		97, 98, 5, 45, 0, 0, 98, 16, 1, 0, 0, 0, 99, 100, 5, 43, 0, 0, 100, 18,
		1, 0, 0, 0, 101, 102, 5, 215, 0, 0, 102, 20, 1, 0, 0, 0, 103, 104, 5, 247,
		0, 0, 104, 22, 1, 0, 0, 0, 105, 106, 5, 8835, 0, 0, 106, 24, 1, 0, 0, 0,
		107, 108, 5, 8744, 0, 0, 108, 26, 1, 0, 0, 0, 109, 110, 5, 8743, 0, 0,
		110, 28, 1, 0, 0, 0, 111, 112, 5, 172, 0, 0, 112, 30, 1, 0, 0, 0, 113,
		114, 5, 60, 0, 0, 114, 32, 1, 0, 0, 0, 115, 116, 5, 8804, 0, 0, 116, 34,
		1, 0, 0, 0, 117, 118, 5, 61, 0, 0, 118, 36, 1, 0, 0, 0, 119, 120, 5, 8805,
		0, 0, 120, 38, 1, 0, 0, 0, 121, 122, 5, 62, 0, 0, 122, 40, 1, 0, 0, 0,
		123, 124, 5, 8800, 0, 0, 124, 42, 1, 0, 0, 0, 125, 129, 5, 58, 0, 0, 126,
		128, 7, 0, 0, 0, 127, 126, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129, 127,
		1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 132, 1, 0, 0, 0, 131, 129, 1, 0,
		0, 0, 132, 133, 5, 61, 0, 0, 133, 44, 1, 0, 0, 0, 134, 135, 5, 39, 0, 0,
		135, 136, 5, 66, 0, 0, 136, 137, 5, 69, 0, 0, 137, 138, 5, 71, 0, 0, 138,
		139, 5, 73, 0, 0, 139, 140, 5, 78, 0, 0, 140, 141, 5, 39, 0, 0, 141, 46,
		1, 0, 0, 0, 142, 143, 5, 39, 0, 0, 143, 144, 5, 69, 0, 0, 144, 145, 5,
		78, 0, 0, 145, 146, 5, 68, 0, 0, 146, 147, 5, 39, 0, 0, 147, 48, 1, 0,
		0, 0, 148, 149, 5, 39, 0, 0, 149, 150, 5, 73, 0, 0, 150, 151, 5, 70, 0,
		0, 151, 152, 5, 39, 0, 0, 152, 50, 1, 0, 0, 0, 153, 154, 5, 39, 0, 0, 154,
		155, 5, 84, 0, 0, 155, 156, 5, 72, 0, 0, 156, 157, 5, 69, 0, 0, 157, 158,
		5, 78, 0, 0, 158, 159, 5, 39, 0, 0, 159, 52, 1, 0, 0, 0, 160, 161, 5, 39,
		0, 0, 161, 162, 5, 69, 0, 0, 162, 163, 5, 76, 0, 0, 163, 164, 5, 83, 0,
		0, 164, 165, 5, 69, 0, 0, 165, 166, 5, 39, 0, 0, 166, 54, 1, 0, 0, 0, 167,
		168, 5, 39, 0, 0, 168, 169, 5, 70, 0, 0, 169, 170, 5, 79, 0, 0, 170, 171,
		5, 82, 0, 0, 171, 172, 5, 39, 0, 0, 172, 56, 1, 0, 0, 0, 173, 174, 5, 39,
		0, 0, 174, 175, 5, 68, 0, 0, 175, 176, 5, 79, 0, 0, 176, 177, 5, 39, 0,
		0, 177, 58, 1, 0, 0, 0, 178, 179, 5, 39, 0, 0, 179, 180, 5, 83, 0, 0, 180,
		181, 5, 84, 0, 0, 181, 182, 5, 69, 0, 0, 182, 183, 5, 80, 0, 0, 183, 184,
		5, 39, 0, 0, 184, 60, 1, 0, 0, 0, 185, 186, 5, 39, 0, 0, 186, 187, 5, 85,
		0, 0, 187, 188, 5, 78, 0, 0, 188, 189, 5, 84, 0, 0, 189, 190, 5, 73, 0,
		0, 190, 191, 5, 76, 0, 0, 191, 192, 5, 39, 0, 0, 192, 62, 1, 0, 0, 0, 193,
		194, 5, 39, 0, 0, 194, 195, 5, 87, 0, 0, 195, 196, 5, 72, 0, 0, 196, 197,
		5, 73, 0, 0, 197, 198, 5, 76, 0, 0, 198, 199, 5, 69, 0, 0, 199, 200, 5,
		39, 0, 0, 200, 64, 1, 0, 0, 0, 201, 202, 5, 39, 0, 0, 202, 203, 5, 80,
		0, 0, 203, 204, 5, 82, 0, 0, 204, 205, 5, 79, 0, 0, 205, 206, 5, 67, 0,
		0, 206, 207, 5, 69, 0, 0, 207, 208, 5, 68, 0, 0, 208, 209, 5, 85, 0, 0,
		209, 210, 5, 82, 0, 0, 210, 211, 5, 69, 0, 0, 211, 212, 5, 39, 0, 0, 212,
		66, 1, 0, 0, 0, 213, 214, 5, 39, 0, 0, 214, 215, 5, 73, 0, 0, 215, 216,
		5, 78, 0, 0, 216, 217, 5, 84, 0, 0, 217, 218, 5, 69, 0, 0, 218, 219, 5,
		71, 0, 0, 219, 220, 5, 69, 0, 0, 220, 221, 5, 82, 0, 0, 221, 222, 5, 39,
		0, 0, 222, 68, 1, 0, 0, 0, 223, 224, 5, 39, 0, 0, 224, 225, 5, 65, 0, 0,
		225, 226, 5, 82, 0, 0, 226, 227, 5, 82, 0, 0, 227, 228, 5, 65, 0, 0, 228,
		229, 5, 89, 0, 0, 229, 230, 5, 39, 0, 0, 230, 70, 1, 0, 0, 0, 231, 232,
		5, 39, 0, 0, 232, 233, 5, 83, 0, 0, 233, 234, 5, 84, 0, 0, 234, 235, 5,
		82, 0, 0, 235, 236, 5, 73, 0, 0, 236, 237, 5, 78, 0, 0, 237, 238, 5, 71,
		0, 0, 238, 239, 5, 39, 0, 0, 239, 72, 1, 0, 0, 0, 240, 241, 5, 39, 0, 0,
		241, 242, 5, 67, 0, 0, 242, 243, 5, 79, 0, 0, 243, 244, 5, 77, 0, 0, 244,
		245, 5, 77, 0, 0, 245, 246, 5, 69, 0, 0, 246, 247, 5, 78, 0, 0, 247, 248,
		5, 84, 0, 0, 248, 249, 5, 39, 0, 0, 249, 253, 1, 0, 0, 0, 250, 252, 8,
		1, 0, 0, 251, 250, 1, 0, 0, 0, 252, 255, 1, 0, 0, 0, 253, 251, 1, 0, 0,
		0, 253, 254, 1, 0, 0, 0, 254, 74, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 256,
		262, 5, 34, 0, 0, 257, 261, 8, 2, 0, 0, 258, 259, 5, 92, 0, 0, 259, 261,
		9, 0, 0, 0, 260, 257, 1, 0, 0, 0, 260, 258, 1, 0, 0, 0, 261, 264, 1, 0,
		0, 0, 262, 260, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 265, 1, 0, 0, 0,
		264, 262, 1, 0, 0, 0, 265, 266, 5, 34, 0, 0, 266, 76, 1, 0, 0, 0, 267,
		271, 7, 3, 0, 0, 268, 270, 7, 4, 0, 0, 269, 268, 1, 0, 0, 0, 270, 273,
		1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 274, 1, 0,
		0, 0, 273, 271, 1, 0, 0, 0, 274, 275, 6, 38, 0, 0, 275, 78, 1, 0, 0, 0,
		276, 280, 7, 5, 0, 0, 277, 279, 7, 6, 0, 0, 278, 277, 1, 0, 0, 0, 279,
		282, 1, 0, 0, 0, 280, 278, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 283,
		1, 0, 0, 0, 282, 280, 1, 0, 0, 0, 283, 284, 6, 39, 1, 0, 284, 80, 1, 0,
		0, 0, 285, 287, 7, 0, 0, 0, 286, 285, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0,
		288, 286, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290,
		291, 6, 40, 2, 0, 291, 82, 1, 0, 0, 0, 8, 0, 129, 253, 260, 262, 271, 280,
		288, 3, 1, 38, 0, 1, 39, 1, 6, 0, 0,
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

// Algol60V2LexerInit initializes any static state used to implement Algol60V2Lexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewAlgol60V2Lexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func Algol60V2LexerInit() {
	staticData := &algol60v2lexerLexerStaticData
	staticData.once.Do(algol60v2lexerLexerInit)
}

// NewAlgol60V2Lexer produces a new lexer instance for the optional input antlr.CharStream.
func NewAlgol60V2Lexer(input antlr.CharStream) *Algol60V2Lexer {
	Algol60V2LexerInit()
	l := new(Algol60V2Lexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &algol60v2lexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Algol60V2.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Algol60V2Lexer tokens.
const (
	Algol60V2LexerT__0          = 1
	Algol60V2LexerT__1          = 2
	Algol60V2LexerT__2          = 3
	Algol60V2LexerT__3          = 4
	Algol60V2LexerT__4          = 5
	Algol60V2LexerT__5          = 6
	Algol60V2LexerT__6          = 7
	Algol60V2LexerT__7          = 8
	Algol60V2LexerT__8          = 9
	Algol60V2LexerT__9          = 10
	Algol60V2LexerT__10         = 11
	Algol60V2LexerT__11         = 12
	Algol60V2LexerT__12         = 13
	Algol60V2LexerT__13         = 14
	Algol60V2LexerT__14         = 15
	Algol60V2LexerT__15         = 16
	Algol60V2LexerT__16         = 17
	Algol60V2LexerT__17         = 18
	Algol60V2LexerT__18         = 19
	Algol60V2LexerT__19         = 20
	Algol60V2LexerT__20         = 21
	Algol60V2LexerASSIGN        = 22
	Algol60V2LexerBEGIN         = 23
	Algol60V2LexerEND           = 24
	Algol60V2LexerIF            = 25
	Algol60V2LexerTHEN          = 26
	Algol60V2LexerELSE          = 27
	Algol60V2LexerFOR           = 28
	Algol60V2LexerDO            = 29
	Algol60V2LexerSTEP          = 30
	Algol60V2LexerUNTIL         = 31
	Algol60V2LexerWHILE         = 32
	Algol60V2LexerPROCEDURE     = 33
	Algol60V2LexerINTEGER       = 34
	Algol60V2LexerARRAY         = 35
	Algol60V2LexerSTRING        = 36
	Algol60V2LexerCOMMENT       = 37
	Algol60V2LexerQUOTED_STRING = 38
	Algol60V2LexerNUMBER        = 39
	Algol60V2LexerIDENTIFIER    = 40
	Algol60V2LexerSKIP_WS       = 41
)

func (l *Algol60V2Lexer) Action(localctx antlr.RuleContext, ruleIndex, actionIndex int) {
	switch ruleIndex {
	case 38:
		l.NUMBER_Action(localctx, actionIndex)

	case 39:
		l.IDENTIFIER_Action(localctx, actionIndex)

	default:
		panic("No registered action for: " + fmt.Sprint(ruleIndex))
	}
}

func (l *Algol60V2Lexer) NUMBER_Action(localctx antlr.RuleContext, actionIndex int) {
	this := l
	_ = this

	switch actionIndex {
	case 0:

		this.SetText(regexp.MustCompile("[ \t\r\n]").
			ReplaceAllLiteralString(this.GetText(), ""))

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *Algol60V2Lexer) IDENTIFIER_Action(localctx antlr.RuleContext, actionIndex int) {
	this := l
	_ = this

	switch actionIndex {
	case 1:

		this.SetText(regexp.MustCompile("[ \t\r\n]").
			ReplaceAllLiteralString(this.GetText(), ""))

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
