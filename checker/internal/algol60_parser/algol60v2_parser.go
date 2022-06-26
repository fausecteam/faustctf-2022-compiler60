// Code generated from Algol60V2.g4 by ANTLR 4.10.1. DO NOT EDIT.

package algol60_parser // Algol60V2
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type Algol60V2Parser struct {
	*antlr.BaseParser
}

var algol60v2ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func algol60v2ParserInit() {
	staticData := &algol60v2ParserStaticData
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
		"program", "block", "block_contents", "declaration", "type_declaration",
		"array_declaration", "array_segment", "string_declaration", "string_segment",
		"procedure_declaration", "formal_parameter_list", "parameter_delimiter",
		"statement", "unconditional_statement", "assignment_statement", "left_part",
		"procedure_statement", "actual_parameter_list", "conditional_statement",
		"for_statement", "for_list", "for_list_element", "expression", "simple_expression",
		"term", "factor", "function_designator", "boolean_expression", "implication",
		"boolean_term", "boolean_factor", "boolean_secondary", "boolean_primary",
		"subscripted_variable", "identifier", "number",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 41, 397, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 1, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 5, 2, 83, 8, 2, 10, 2,
		12, 2, 86, 9, 2, 1, 2, 5, 2, 89, 8, 2, 10, 2, 12, 2, 92, 9, 2, 1, 2, 3,
		2, 95, 8, 2, 1, 2, 5, 2, 98, 8, 2, 10, 2, 12, 2, 101, 9, 2, 1, 2, 1, 2,
		5, 2, 105, 8, 2, 10, 2, 12, 2, 108, 9, 2, 1, 2, 5, 2, 111, 8, 2, 10, 2,
		12, 2, 114, 9, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 121, 8, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 5, 4, 127, 8, 4, 10, 4, 12, 4, 130, 9, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 5, 5, 137, 8, 5, 10, 5, 12, 5, 140, 9, 5, 1, 6, 1, 6,
		1, 6, 5, 6, 145, 8, 6, 10, 6, 12, 6, 148, 9, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		5, 6, 154, 8, 6, 10, 6, 12, 6, 157, 9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 5, 7, 165, 8, 7, 10, 7, 12, 7, 168, 9, 7, 1, 8, 1, 8, 1, 8, 5, 8,
		173, 8, 8, 10, 8, 12, 8, 176, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 3, 9,
		183, 8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 191, 8, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 5, 9, 197, 8, 9, 10, 9, 12, 9, 200, 9, 9, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 10, 1, 10, 3, 10, 208, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 3, 11, 215, 8, 11, 1, 12, 1, 12, 1, 12, 3, 12, 220, 8, 12, 1, 13,
		1, 13, 1, 13, 1, 13, 3, 13, 226, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 15, 1, 15, 3, 15, 235, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		3, 16, 242, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 248, 8, 17, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 256, 8, 18, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 5, 20, 268, 8, 20,
		10, 20, 12, 20, 271, 9, 20, 1, 21, 1, 21, 1, 21, 3, 21, 276, 8, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 284, 8, 21, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 294, 8, 22, 1, 23, 3,
		23, 297, 8, 23, 1, 23, 1, 23, 1, 23, 5, 23, 302, 8, 23, 10, 23, 12, 23,
		305, 9, 23, 1, 24, 1, 24, 1, 24, 5, 24, 310, 8, 24, 10, 24, 12, 24, 313,
		9, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3,
		25, 324, 8, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 339, 8, 27, 1, 28, 1, 28, 1,
		28, 5, 28, 344, 8, 28, 10, 28, 12, 28, 347, 9, 28, 1, 29, 1, 29, 1, 29,
		5, 29, 352, 8, 29, 10, 29, 12, 29, 355, 9, 29, 1, 30, 1, 30, 1, 30, 5,
		30, 360, 8, 30, 10, 30, 12, 30, 363, 9, 30, 1, 31, 3, 31, 366, 8, 31, 1,
		31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32,
		378, 8, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 5, 33, 385, 8, 33, 10, 33,
		12, 33, 388, 9, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1,
		35, 0, 0, 36, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
		68, 70, 0, 3, 1, 0, 8, 9, 1, 0, 10, 11, 1, 0, 16, 21, 410, 0, 72, 1, 0,
		0, 0, 2, 75, 1, 0, 0, 0, 4, 84, 1, 0, 0, 0, 6, 120, 1, 0, 0, 0, 8, 122,
		1, 0, 0, 0, 10, 131, 1, 0, 0, 0, 12, 141, 1, 0, 0, 0, 14, 160, 1, 0, 0,
		0, 16, 169, 1, 0, 0, 0, 18, 182, 1, 0, 0, 0, 20, 203, 1, 0, 0, 0, 22, 214,
		1, 0, 0, 0, 24, 219, 1, 0, 0, 0, 26, 225, 1, 0, 0, 0, 28, 227, 1, 0, 0,
		0, 30, 234, 1, 0, 0, 0, 32, 236, 1, 0, 0, 0, 34, 243, 1, 0, 0, 0, 36, 249,
		1, 0, 0, 0, 38, 257, 1, 0, 0, 0, 40, 264, 1, 0, 0, 0, 42, 272, 1, 0, 0,
		0, 44, 293, 1, 0, 0, 0, 46, 296, 1, 0, 0, 0, 48, 306, 1, 0, 0, 0, 50, 323,
		1, 0, 0, 0, 52, 325, 1, 0, 0, 0, 54, 338, 1, 0, 0, 0, 56, 340, 1, 0, 0,
		0, 58, 348, 1, 0, 0, 0, 60, 356, 1, 0, 0, 0, 62, 365, 1, 0, 0, 0, 64, 377,
		1, 0, 0, 0, 66, 379, 1, 0, 0, 0, 68, 391, 1, 0, 0, 0, 70, 393, 1, 0, 0,
		0, 72, 73, 3, 2, 1, 0, 73, 74, 5, 0, 0, 1, 74, 1, 1, 0, 0, 0, 75, 76, 5,
		23, 0, 0, 76, 77, 3, 4, 2, 0, 77, 78, 5, 24, 0, 0, 78, 3, 1, 0, 0, 0, 79,
		80, 3, 6, 3, 0, 80, 81, 5, 1, 0, 0, 81, 83, 1, 0, 0, 0, 82, 79, 1, 0, 0,
		0, 83, 86, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 90,
		1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 87, 89, 5, 1, 0, 0, 88, 87, 1, 0, 0, 0,
		89, 92, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 94, 1,
		0, 0, 0, 92, 90, 1, 0, 0, 0, 93, 95, 3, 24, 12, 0, 94, 93, 1, 0, 0, 0,
		94, 95, 1, 0, 0, 0, 95, 99, 1, 0, 0, 0, 96, 98, 5, 1, 0, 0, 97, 96, 1,
		0, 0, 0, 98, 101, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0,
		100, 106, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 102, 103, 5, 1, 0, 0, 103, 105,
		3, 24, 12, 0, 104, 102, 1, 0, 0, 0, 105, 108, 1, 0, 0, 0, 106, 104, 1,
		0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 112, 1, 0, 0, 0, 108, 106, 1, 0, 0,
		0, 109, 111, 5, 1, 0, 0, 110, 109, 1, 0, 0, 0, 111, 114, 1, 0, 0, 0, 112,
		110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 5, 1, 0, 0, 0, 114, 112, 1,
		0, 0, 0, 115, 121, 3, 8, 4, 0, 116, 121, 3, 10, 5, 0, 117, 121, 3, 14,
		7, 0, 118, 121, 3, 18, 9, 0, 119, 121, 5, 37, 0, 0, 120, 115, 1, 0, 0,
		0, 120, 116, 1, 0, 0, 0, 120, 117, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120,
		119, 1, 0, 0, 0, 121, 7, 1, 0, 0, 0, 122, 123, 5, 34, 0, 0, 123, 128, 3,
		68, 34, 0, 124, 125, 5, 2, 0, 0, 125, 127, 3, 68, 34, 0, 126, 124, 1, 0,
		0, 0, 127, 130, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0,
		129, 9, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 131, 132, 5, 34, 0, 0, 132, 133,
		5, 35, 0, 0, 133, 138, 3, 12, 6, 0, 134, 135, 5, 2, 0, 0, 135, 137, 3,
		12, 6, 0, 136, 134, 1, 0, 0, 0, 137, 140, 1, 0, 0, 0, 138, 136, 1, 0, 0,
		0, 138, 139, 1, 0, 0, 0, 139, 11, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141,
		146, 3, 68, 34, 0, 142, 143, 5, 2, 0, 0, 143, 145, 3, 68, 34, 0, 144, 142,
		1, 0, 0, 0, 145, 148, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0,
		0, 0, 147, 149, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 149, 150, 5, 3, 0, 0,
		150, 155, 3, 70, 35, 0, 151, 152, 5, 2, 0, 0, 152, 154, 3, 70, 35, 0, 153,
		151, 1, 0, 0, 0, 154, 157, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 155, 156,
		1, 0, 0, 0, 156, 158, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 158, 159, 5, 4,
		0, 0, 159, 13, 1, 0, 0, 0, 160, 161, 5, 36, 0, 0, 161, 166, 3, 16, 8, 0,
		162, 163, 5, 2, 0, 0, 163, 165, 3, 16, 8, 0, 164, 162, 1, 0, 0, 0, 165,
		168, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 15, 1,
		0, 0, 0, 168, 166, 1, 0, 0, 0, 169, 174, 3, 68, 34, 0, 170, 171, 5, 2,
		0, 0, 171, 173, 3, 68, 34, 0, 172, 170, 1, 0, 0, 0, 173, 176, 1, 0, 0,
		0, 174, 172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 177, 1, 0, 0, 0, 176,
		174, 1, 0, 0, 0, 177, 178, 5, 3, 0, 0, 178, 179, 3, 70, 35, 0, 179, 180,
		5, 4, 0, 0, 180, 17, 1, 0, 0, 0, 181, 183, 5, 34, 0, 0, 182, 181, 1, 0,
		0, 0, 182, 183, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 185, 5, 33, 0, 0,
		185, 190, 3, 68, 34, 0, 186, 187, 5, 5, 0, 0, 187, 188, 3, 20, 10, 0, 188,
		189, 5, 6, 0, 0, 189, 191, 1, 0, 0, 0, 190, 186, 1, 0, 0, 0, 190, 191,
		1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 198, 5, 1, 0, 0, 193, 194, 3, 6,
		3, 0, 194, 195, 5, 1, 0, 0, 195, 197, 1, 0, 0, 0, 196, 193, 1, 0, 0, 0,
		197, 200, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199,
		201, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 201, 202, 3, 24, 12, 0, 202, 19,
		1, 0, 0, 0, 203, 207, 3, 68, 34, 0, 204, 205, 3, 22, 11, 0, 205, 206, 3,
		20, 10, 0, 206, 208, 1, 0, 0, 0, 207, 204, 1, 0, 0, 0, 207, 208, 1, 0,
		0, 0, 208, 21, 1, 0, 0, 0, 209, 215, 5, 2, 0, 0, 210, 211, 5, 6, 0, 0,
		211, 212, 5, 40, 0, 0, 212, 213, 5, 7, 0, 0, 213, 215, 5, 5, 0, 0, 214,
		209, 1, 0, 0, 0, 214, 210, 1, 0, 0, 0, 215, 23, 1, 0, 0, 0, 216, 220, 3,
		26, 13, 0, 217, 220, 3, 36, 18, 0, 218, 220, 5, 37, 0, 0, 219, 216, 1,
		0, 0, 0, 219, 217, 1, 0, 0, 0, 219, 218, 1, 0, 0, 0, 220, 25, 1, 0, 0,
		0, 221, 226, 3, 38, 19, 0, 222, 226, 3, 2, 1, 0, 223, 226, 3, 28, 14, 0,
		224, 226, 3, 32, 16, 0, 225, 221, 1, 0, 0, 0, 225, 222, 1, 0, 0, 0, 225,
		223, 1, 0, 0, 0, 225, 224, 1, 0, 0, 0, 226, 27, 1, 0, 0, 0, 227, 228, 3,
		30, 15, 0, 228, 229, 5, 22, 0, 0, 229, 230, 1, 0, 0, 0, 230, 231, 3, 44,
		22, 0, 231, 29, 1, 0, 0, 0, 232, 235, 3, 68, 34, 0, 233, 235, 3, 66, 33,
		0, 234, 232, 1, 0, 0, 0, 234, 233, 1, 0, 0, 0, 235, 31, 1, 0, 0, 0, 236,
		241, 3, 68, 34, 0, 237, 238, 5, 5, 0, 0, 238, 239, 3, 34, 17, 0, 239, 240,
		5, 6, 0, 0, 240, 242, 1, 0, 0, 0, 241, 237, 1, 0, 0, 0, 241, 242, 1, 0,
		0, 0, 242, 33, 1, 0, 0, 0, 243, 247, 3, 44, 22, 0, 244, 245, 3, 22, 11,
		0, 245, 246, 3, 34, 17, 0, 246, 248, 1, 0, 0, 0, 247, 244, 1, 0, 0, 0,
		247, 248, 1, 0, 0, 0, 248, 35, 1, 0, 0, 0, 249, 250, 5, 25, 0, 0, 250,
		251, 3, 54, 27, 0, 251, 252, 5, 26, 0, 0, 252, 255, 3, 26, 13, 0, 253,
		254, 5, 27, 0, 0, 254, 256, 3, 24, 12, 0, 255, 253, 1, 0, 0, 0, 255, 256,
		1, 0, 0, 0, 256, 37, 1, 0, 0, 0, 257, 258, 5, 28, 0, 0, 258, 259, 3, 68,
		34, 0, 259, 260, 5, 22, 0, 0, 260, 261, 3, 40, 20, 0, 261, 262, 5, 29,
		0, 0, 262, 263, 3, 24, 12, 0, 263, 39, 1, 0, 0, 0, 264, 269, 3, 42, 21,
		0, 265, 266, 5, 2, 0, 0, 266, 268, 3, 42, 21, 0, 267, 265, 1, 0, 0, 0,
		268, 271, 1, 0, 0, 0, 269, 267, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270,
		41, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 272, 283, 3, 44, 22, 0, 273, 275,
		5, 30, 0, 0, 274, 276, 5, 8, 0, 0, 275, 274, 1, 0, 0, 0, 275, 276, 1, 0,
		0, 0, 276, 277, 1, 0, 0, 0, 277, 278, 3, 70, 35, 0, 278, 279, 5, 31, 0,
		0, 279, 280, 3, 44, 22, 0, 280, 284, 1, 0, 0, 0, 281, 282, 5, 32, 0, 0,
		282, 284, 3, 54, 27, 0, 283, 273, 1, 0, 0, 0, 283, 281, 1, 0, 0, 0, 283,
		284, 1, 0, 0, 0, 284, 43, 1, 0, 0, 0, 285, 294, 3, 46, 23, 0, 286, 287,
		5, 25, 0, 0, 287, 288, 3, 54, 27, 0, 288, 289, 5, 26, 0, 0, 289, 290, 3,
		46, 23, 0, 290, 291, 5, 27, 0, 0, 291, 292, 3, 44, 22, 0, 292, 294, 1,
		0, 0, 0, 293, 285, 1, 0, 0, 0, 293, 286, 1, 0, 0, 0, 294, 45, 1, 0, 0,
		0, 295, 297, 7, 0, 0, 0, 296, 295, 1, 0, 0, 0, 296, 297, 1, 0, 0, 0, 297,
		298, 1, 0, 0, 0, 298, 303, 3, 48, 24, 0, 299, 300, 7, 0, 0, 0, 300, 302,
		3, 48, 24, 0, 301, 299, 1, 0, 0, 0, 302, 305, 1, 0, 0, 0, 303, 301, 1,
		0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 47, 1, 0, 0, 0, 305, 303, 1, 0, 0,
		0, 306, 311, 3, 50, 25, 0, 307, 308, 7, 1, 0, 0, 308, 310, 3, 50, 25, 0,
		309, 307, 1, 0, 0, 0, 310, 313, 1, 0, 0, 0, 311, 309, 1, 0, 0, 0, 311,
		312, 1, 0, 0, 0, 312, 49, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0, 314, 324, 3,
		70, 35, 0, 315, 324, 5, 38, 0, 0, 316, 324, 3, 68, 34, 0, 317, 324, 3,
		52, 26, 0, 318, 324, 3, 66, 33, 0, 319, 320, 5, 5, 0, 0, 320, 321, 3, 44,
		22, 0, 321, 322, 5, 6, 0, 0, 322, 324, 1, 0, 0, 0, 323, 314, 1, 0, 0, 0,
		323, 315, 1, 0, 0, 0, 323, 316, 1, 0, 0, 0, 323, 317, 1, 0, 0, 0, 323,
		318, 1, 0, 0, 0, 323, 319, 1, 0, 0, 0, 324, 51, 1, 0, 0, 0, 325, 326, 3,
		68, 34, 0, 326, 327, 5, 5, 0, 0, 327, 328, 3, 34, 17, 0, 328, 329, 5, 6,
		0, 0, 329, 53, 1, 0, 0, 0, 330, 339, 3, 56, 28, 0, 331, 332, 5, 25, 0,
		0, 332, 333, 3, 54, 27, 0, 333, 334, 5, 26, 0, 0, 334, 335, 3, 56, 28,
		0, 335, 336, 5, 27, 0, 0, 336, 337, 3, 54, 27, 0, 337, 339, 1, 0, 0, 0,
		338, 330, 1, 0, 0, 0, 338, 331, 1, 0, 0, 0, 339, 55, 1, 0, 0, 0, 340, 345,
		3, 58, 29, 0, 341, 342, 5, 12, 0, 0, 342, 344, 3, 58, 29, 0, 343, 341,
		1, 0, 0, 0, 344, 347, 1, 0, 0, 0, 345, 343, 1, 0, 0, 0, 345, 346, 1, 0,
		0, 0, 346, 57, 1, 0, 0, 0, 347, 345, 1, 0, 0, 0, 348, 353, 3, 60, 30, 0,
		349, 350, 5, 13, 0, 0, 350, 352, 3, 60, 30, 0, 351, 349, 1, 0, 0, 0, 352,
		355, 1, 0, 0, 0, 353, 351, 1, 0, 0, 0, 353, 354, 1, 0, 0, 0, 354, 59, 1,
		0, 0, 0, 355, 353, 1, 0, 0, 0, 356, 361, 3, 62, 31, 0, 357, 358, 5, 14,
		0, 0, 358, 360, 3, 62, 31, 0, 359, 357, 1, 0, 0, 0, 360, 363, 1, 0, 0,
		0, 361, 359, 1, 0, 0, 0, 361, 362, 1, 0, 0, 0, 362, 61, 1, 0, 0, 0, 363,
		361, 1, 0, 0, 0, 364, 366, 5, 15, 0, 0, 365, 364, 1, 0, 0, 0, 365, 366,
		1, 0, 0, 0, 366, 367, 1, 0, 0, 0, 367, 368, 3, 64, 32, 0, 368, 63, 1, 0,
		0, 0, 369, 370, 3, 46, 23, 0, 370, 371, 7, 2, 0, 0, 371, 372, 3, 46, 23,
		0, 372, 378, 1, 0, 0, 0, 373, 374, 5, 5, 0, 0, 374, 375, 3, 54, 27, 0,
		375, 376, 5, 6, 0, 0, 376, 378, 1, 0, 0, 0, 377, 369, 1, 0, 0, 0, 377,
		373, 1, 0, 0, 0, 378, 65, 1, 0, 0, 0, 379, 380, 3, 68, 34, 0, 380, 381,
		5, 3, 0, 0, 381, 386, 3, 44, 22, 0, 382, 383, 5, 2, 0, 0, 383, 385, 3,
		44, 22, 0, 384, 382, 1, 0, 0, 0, 385, 388, 1, 0, 0, 0, 386, 384, 1, 0,
		0, 0, 386, 387, 1, 0, 0, 0, 387, 389, 1, 0, 0, 0, 388, 386, 1, 0, 0, 0,
		389, 390, 5, 4, 0, 0, 390, 67, 1, 0, 0, 0, 391, 392, 5, 40, 0, 0, 392,
		69, 1, 0, 0, 0, 393, 394, 5, 39, 0, 0, 394, 395, 6, 35, -1, 0, 395, 71,
		1, 0, 0, 0, 39, 84, 90, 94, 99, 106, 112, 120, 128, 138, 146, 155, 166,
		174, 182, 190, 198, 207, 214, 219, 225, 234, 241, 247, 255, 269, 275, 283,
		293, 296, 303, 311, 323, 338, 345, 353, 361, 365, 377, 386,
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

// Algol60V2ParserInit initializes any static state used to implement Algol60V2Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAlgol60V2Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Algol60V2ParserInit() {
	staticData := &algol60v2ParserStaticData
	staticData.once.Do(algol60v2ParserInit)
}

// NewAlgol60V2Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewAlgol60V2Parser(input antlr.TokenStream) *Algol60V2Parser {
	Algol60V2ParserInit()
	this := new(Algol60V2Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &algol60v2ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Algol60V2.g4"

	return this
}

// Algol60V2Parser tokens.
const (
	Algol60V2ParserEOF           = antlr.TokenEOF
	Algol60V2ParserT__0          = 1
	Algol60V2ParserT__1          = 2
	Algol60V2ParserT__2          = 3
	Algol60V2ParserT__3          = 4
	Algol60V2ParserT__4          = 5
	Algol60V2ParserT__5          = 6
	Algol60V2ParserT__6          = 7
	Algol60V2ParserT__7          = 8
	Algol60V2ParserT__8          = 9
	Algol60V2ParserT__9          = 10
	Algol60V2ParserT__10         = 11
	Algol60V2ParserT__11         = 12
	Algol60V2ParserT__12         = 13
	Algol60V2ParserT__13         = 14
	Algol60V2ParserT__14         = 15
	Algol60V2ParserT__15         = 16
	Algol60V2ParserT__16         = 17
	Algol60V2ParserT__17         = 18
	Algol60V2ParserT__18         = 19
	Algol60V2ParserT__19         = 20
	Algol60V2ParserT__20         = 21
	Algol60V2ParserASSIGN        = 22
	Algol60V2ParserBEGIN         = 23
	Algol60V2ParserEND           = 24
	Algol60V2ParserIF            = 25
	Algol60V2ParserTHEN          = 26
	Algol60V2ParserELSE          = 27
	Algol60V2ParserFOR           = 28
	Algol60V2ParserDO            = 29
	Algol60V2ParserSTEP          = 30
	Algol60V2ParserUNTIL         = 31
	Algol60V2ParserWHILE         = 32
	Algol60V2ParserPROCEDURE     = 33
	Algol60V2ParserINTEGER       = 34
	Algol60V2ParserARRAY         = 35
	Algol60V2ParserSTRING        = 36
	Algol60V2ParserCOMMENT       = 37
	Algol60V2ParserQUOTED_STRING = 38
	Algol60V2ParserNUMBER        = 39
	Algol60V2ParserIDENTIFIER    = 40
	Algol60V2ParserSKIP_WS       = 41
)

// Algol60V2Parser rules.
const (
	Algol60V2ParserRULE_program                 = 0
	Algol60V2ParserRULE_block                   = 1
	Algol60V2ParserRULE_block_contents          = 2
	Algol60V2ParserRULE_declaration             = 3
	Algol60V2ParserRULE_type_declaration        = 4
	Algol60V2ParserRULE_array_declaration       = 5
	Algol60V2ParserRULE_array_segment           = 6
	Algol60V2ParserRULE_string_declaration      = 7
	Algol60V2ParserRULE_string_segment          = 8
	Algol60V2ParserRULE_procedure_declaration   = 9
	Algol60V2ParserRULE_formal_parameter_list   = 10
	Algol60V2ParserRULE_parameter_delimiter     = 11
	Algol60V2ParserRULE_statement               = 12
	Algol60V2ParserRULE_unconditional_statement = 13
	Algol60V2ParserRULE_assignment_statement    = 14
	Algol60V2ParserRULE_left_part               = 15
	Algol60V2ParserRULE_procedure_statement     = 16
	Algol60V2ParserRULE_actual_parameter_list   = 17
	Algol60V2ParserRULE_conditional_statement   = 18
	Algol60V2ParserRULE_for_statement           = 19
	Algol60V2ParserRULE_for_list                = 20
	Algol60V2ParserRULE_for_list_element        = 21
	Algol60V2ParserRULE_expression              = 22
	Algol60V2ParserRULE_simple_expression       = 23
	Algol60V2ParserRULE_term                    = 24
	Algol60V2ParserRULE_factor                  = 25
	Algol60V2ParserRULE_function_designator     = 26
	Algol60V2ParserRULE_boolean_expression      = 27
	Algol60V2ParserRULE_implication             = 28
	Algol60V2ParserRULE_boolean_term            = 29
	Algol60V2ParserRULE_boolean_factor          = 30
	Algol60V2ParserRULE_boolean_secondary       = 31
	Algol60V2ParserRULE_boolean_primary         = 32
	Algol60V2ParserRULE_subscripted_variable    = 33
	Algol60V2ParserRULE_identifier              = 34
	Algol60V2ParserRULE_number                  = 35
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Algol60V2ParserRULE_program)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Block()
	}
	{
		p.SetState(73)
		p.Match(Algol60V2ParserEOF)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) BEGIN() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserBEGIN, 0)
}

func (s *BlockContext) Block_contents() IBlock_contentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_contentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_contentsContext)
}

func (s *BlockContext) END() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserEND, 0)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Algol60V2ParserRULE_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(Algol60V2ParserBEGIN)
	}
	{
		p.SetState(76)
		p.Block_contents()
	}
	{
		p.SetState(77)
		p.Match(Algol60V2ParserEND)
	}

	return localctx
}

// IBlock_contentsContext is an interface to support dynamic dispatch.
type IBlock_contentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_declaration returns the _declaration rule contexts.
	Get_declaration() IDeclarationContext

	// Get_statement returns the _statement rule contexts.
	Get_statement() IStatementContext

	// Set_declaration sets the _declaration rule contexts.
	Set_declaration(IDeclarationContext)

	// Set_statement sets the _statement rule contexts.
	Set_statement(IStatementContext)

	// GetDecls returns the decls rule context list.
	GetDecls() []IDeclarationContext

	// GetStmts returns the stmts rule context list.
	GetStmts() []IStatementContext

	// SetDecls sets the decls rule context list.
	SetDecls([]IDeclarationContext)

	// SetStmts sets the stmts rule context list.
	SetStmts([]IStatementContext)

	// IsBlock_contentsContext differentiates from other interfaces.
	IsBlock_contentsContext()
}

type Block_contentsContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	_declaration IDeclarationContext
	decls        []IDeclarationContext
	_statement   IStatementContext
	stmts        []IStatementContext
}

func NewEmptyBlock_contentsContext() *Block_contentsContext {
	var p = new(Block_contentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_block_contents
	return p
}

func (*Block_contentsContext) IsBlock_contentsContext() {}

func NewBlock_contentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_contentsContext {
	var p = new(Block_contentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_block_contents

	return p
}

func (s *Block_contentsContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_contentsContext) Get_declaration() IDeclarationContext { return s._declaration }

func (s *Block_contentsContext) Get_statement() IStatementContext { return s._statement }

func (s *Block_contentsContext) Set_declaration(v IDeclarationContext) { s._declaration = v }

func (s *Block_contentsContext) Set_statement(v IStatementContext) { s._statement = v }

func (s *Block_contentsContext) GetDecls() []IDeclarationContext { return s.decls }

func (s *Block_contentsContext) GetStmts() []IStatementContext { return s.stmts }

func (s *Block_contentsContext) SetDecls(v []IDeclarationContext) { s.decls = v }

func (s *Block_contentsContext) SetStmts(v []IStatementContext) { s.stmts = v }

func (s *Block_contentsContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *Block_contentsContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
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

	return t.(IDeclarationContext)
}

func (s *Block_contentsContext) AllStatement() []IStatementContext {
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

func (s *Block_contentsContext) Statement(i int) IStatementContext {
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

func (s *Block_contentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_contentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_contentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitBlock_contents(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Block_contents() (localctx IBlock_contentsContext) {
	this := p
	_ = this

	localctx = NewBlock_contentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Algol60V2ParserRULE_block_contents)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(79)

				var _x = p.Declaration()

				localctx.(*Block_contentsContext)._declaration = _x
			}
			localctx.(*Block_contentsContext).decls = append(localctx.(*Block_contentsContext).decls, localctx.(*Block_contentsContext)._declaration)
			{
				p.SetState(80)
				p.Match(Algol60V2ParserT__0)
			}

		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(87)
				p.Match(Algol60V2ParserT__0)
			}

		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-23)&-(0x1f+1)) == 0 && ((1<<uint((_la-23)))&((1<<(Algol60V2ParserBEGIN-23))|(1<<(Algol60V2ParserIF-23))|(1<<(Algol60V2ParserFOR-23))|(1<<(Algol60V2ParserCOMMENT-23))|(1<<(Algol60V2ParserIDENTIFIER-23)))) != 0 {
		{
			p.SetState(93)

			var _x = p.Statement()

			localctx.(*Block_contentsContext)._statement = _x
		}
		localctx.(*Block_contentsContext).stmts = append(localctx.(*Block_contentsContext).stmts, localctx.(*Block_contentsContext)._statement)

	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(96)
				p.Match(Algol60V2ParserT__0)
			}

		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(102)
				p.Match(Algol60V2ParserT__0)
			}
			{
				p.SetState(103)

				var _x = p.Statement()

				localctx.(*Block_contentsContext)._statement = _x
			}
			localctx.(*Block_contentsContext).stmts = append(localctx.(*Block_contentsContext).stmts, localctx.(*Block_contentsContext)._statement)

		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Algol60V2ParserT__0 {
		{
			p.SetState(109)
			p.Match(Algol60V2ParserT__0)
		}

		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) Type_declaration() IType_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_declarationContext)
}

func (s *DeclarationContext) Array_declaration() IArray_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_declarationContext)
}

func (s *DeclarationContext) String_declaration() IString_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_declarationContext)
}

func (s *DeclarationContext) Procedure_declaration() IProcedure_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProcedure_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProcedure_declarationContext)
}

func (s *DeclarationContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserCOMMENT, 0)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Declaration() (localctx IDeclarationContext) {
	this := p
	_ = this

	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Algol60V2ParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(115)
			p.Type_declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)
			p.Array_declaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(117)
			p.String_declaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(118)
			p.Procedure_declaration()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(119)
			p.Match(Algol60V2ParserCOMMENT)
		}

	}

	return localctx
}

// IType_declarationContext is an interface to support dynamic dispatch.
type IType_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_identifier returns the _identifier rule contexts.
	Get_identifier() IIdentifierContext

	// Set_identifier sets the _identifier rule contexts.
	Set_identifier(IIdentifierContext)

	// GetVars returns the vars rule context list.
	GetVars() []IIdentifierContext

	// SetVars sets the vars rule context list.
	SetVars([]IIdentifierContext)

	// IsType_declarationContext differentiates from other interfaces.
	IsType_declarationContext()
}

type Type_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	_identifier IIdentifierContext
	vars        []IIdentifierContext
}

func NewEmptyType_declarationContext() *Type_declarationContext {
	var p = new(Type_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_type_declaration
	return p
}

func (*Type_declarationContext) IsType_declarationContext() {}

func NewType_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_declarationContext {
	var p = new(Type_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_type_declaration

	return p
}

func (s *Type_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_declarationContext) Get_identifier() IIdentifierContext { return s._identifier }

func (s *Type_declarationContext) Set_identifier(v IIdentifierContext) { s._identifier = v }

func (s *Type_declarationContext) GetVars() []IIdentifierContext { return s.vars }

func (s *Type_declarationContext) SetVars(v []IIdentifierContext) { s.vars = v }

func (s *Type_declarationContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserINTEGER, 0)
}

func (s *Type_declarationContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *Type_declarationContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
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

	return t.(IIdentifierContext)
}

func (s *Type_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitType_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Type_declaration() (localctx IType_declarationContext) {
	this := p
	_ = this

	localctx = NewType_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Algol60V2ParserRULE_type_declaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Match(Algol60V2ParserINTEGER)
	}
	{
		p.SetState(123)

		var _x = p.Identifier()

		localctx.(*Type_declarationContext)._identifier = _x
	}
	localctx.(*Type_declarationContext).vars = append(localctx.(*Type_declarationContext).vars, localctx.(*Type_declarationContext)._identifier)
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Algol60V2ParserT__1 {
		{
			p.SetState(124)
			p.Match(Algol60V2ParserT__1)
		}
		{
			p.SetState(125)

			var _x = p.Identifier()

			localctx.(*Type_declarationContext)._identifier = _x
		}
		localctx.(*Type_declarationContext).vars = append(localctx.(*Type_declarationContext).vars, localctx.(*Type_declarationContext)._identifier)

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IArray_declarationContext is an interface to support dynamic dispatch.
type IArray_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_array_segment returns the _array_segment rule contexts.
	Get_array_segment() IArray_segmentContext

	// Set_array_segment sets the _array_segment rule contexts.
	Set_array_segment(IArray_segmentContext)

	// GetSegments returns the segments rule context list.
	GetSegments() []IArray_segmentContext

	// SetSegments sets the segments rule context list.
	SetSegments([]IArray_segmentContext)

	// IsArray_declarationContext differentiates from other interfaces.
	IsArray_declarationContext()
}

type Array_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	_array_segment IArray_segmentContext
	segments       []IArray_segmentContext
}

func NewEmptyArray_declarationContext() *Array_declarationContext {
	var p = new(Array_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_array_declaration
	return p
}

func (*Array_declarationContext) IsArray_declarationContext() {}

func NewArray_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_declarationContext {
	var p = new(Array_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_array_declaration

	return p
}

func (s *Array_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_declarationContext) Get_array_segment() IArray_segmentContext { return s._array_segment }

func (s *Array_declarationContext) Set_array_segment(v IArray_segmentContext) { s._array_segment = v }

func (s *Array_declarationContext) GetSegments() []IArray_segmentContext { return s.segments }

func (s *Array_declarationContext) SetSegments(v []IArray_segmentContext) { s.segments = v }

func (s *Array_declarationContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserINTEGER, 0)
}

func (s *Array_declarationContext) ARRAY() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserARRAY, 0)
}

func (s *Array_declarationContext) AllArray_segment() []IArray_segmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArray_segmentContext); ok {
			len++
		}
	}

	tst := make([]IArray_segmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArray_segmentContext); ok {
			tst[i] = t.(IArray_segmentContext)
			i++
		}
	}

	return tst
}

func (s *Array_declarationContext) Array_segment(i int) IArray_segmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_segmentContext); ok {
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

	return t.(IArray_segmentContext)
}

func (s *Array_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitArray_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Array_declaration() (localctx IArray_declarationContext) {
	this := p
	_ = this

	localctx = NewArray_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Algol60V2ParserRULE_array_declaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(Algol60V2ParserINTEGER)
	}
	{
		p.SetState(132)
		p.Match(Algol60V2ParserARRAY)
	}
	{
		p.SetState(133)

		var _x = p.Array_segment()

		localctx.(*Array_declarationContext)._array_segment = _x
	}
	localctx.(*Array_declarationContext).segments = append(localctx.(*Array_declarationContext).segments, localctx.(*Array_declarationContext)._array_segment)
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Algol60V2ParserT__1 {
		{
			p.SetState(134)
			p.Match(Algol60V2ParserT__1)
		}
		{
			p.SetState(135)

			var _x = p.Array_segment()

			localctx.(*Array_declarationContext)._array_segment = _x
		}
		localctx.(*Array_declarationContext).segments = append(localctx.(*Array_declarationContext).segments, localctx.(*Array_declarationContext)._array_segment)

		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IArray_segmentContext is an interface to support dynamic dispatch.
type IArray_segmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_identifier returns the _identifier rule contexts.
	Get_identifier() IIdentifierContext

	// Get_number returns the _number rule contexts.
	Get_number() INumberContext

	// Set_identifier sets the _identifier rule contexts.
	Set_identifier(IIdentifierContext)

	// Set_number sets the _number rule contexts.
	Set_number(INumberContext)

	// GetVars returns the vars rule context list.
	GetVars() []IIdentifierContext

	// GetSizes returns the sizes rule context list.
	GetSizes() []INumberContext

	// SetVars sets the vars rule context list.
	SetVars([]IIdentifierContext)

	// SetSizes sets the sizes rule context list.
	SetSizes([]INumberContext)

	// IsArray_segmentContext differentiates from other interfaces.
	IsArray_segmentContext()
}

type Array_segmentContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	_identifier IIdentifierContext
	vars        []IIdentifierContext
	_number     INumberContext
	sizes       []INumberContext
}

func NewEmptyArray_segmentContext() *Array_segmentContext {
	var p = new(Array_segmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_array_segment
	return p
}

func (*Array_segmentContext) IsArray_segmentContext() {}

func NewArray_segmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_segmentContext {
	var p = new(Array_segmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_array_segment

	return p
}

func (s *Array_segmentContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_segmentContext) Get_identifier() IIdentifierContext { return s._identifier }

func (s *Array_segmentContext) Get_number() INumberContext { return s._number }

func (s *Array_segmentContext) Set_identifier(v IIdentifierContext) { s._identifier = v }

func (s *Array_segmentContext) Set_number(v INumberContext) { s._number = v }

func (s *Array_segmentContext) GetVars() []IIdentifierContext { return s.vars }

func (s *Array_segmentContext) GetSizes() []INumberContext { return s.sizes }

func (s *Array_segmentContext) SetVars(v []IIdentifierContext) { s.vars = v }

func (s *Array_segmentContext) SetSizes(v []INumberContext) { s.sizes = v }

func (s *Array_segmentContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *Array_segmentContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
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

	return t.(IIdentifierContext)
}

func (s *Array_segmentContext) AllNumber() []INumberContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumberContext); ok {
			len++
		}
	}

	tst := make([]INumberContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumberContext); ok {
			tst[i] = t.(INumberContext)
			i++
		}
	}

	return tst
}

func (s *Array_segmentContext) Number(i int) INumberContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
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

	return t.(INumberContext)
}

func (s *Array_segmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_segmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_segmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitArray_segment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Array_segment() (localctx IArray_segmentContext) {
	this := p
	_ = this

	localctx = NewArray_segmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, Algol60V2ParserRULE_array_segment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)

		var _x = p.Identifier()

		localctx.(*Array_segmentContext)._identifier = _x
	}
	localctx.(*Array_segmentContext).vars = append(localctx.(*Array_segmentContext).vars, localctx.(*Array_segmentContext)._identifier)
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Algol60V2ParserT__1 {
		{
			p.SetState(142)
			p.Match(Algol60V2ParserT__1)
		}
		{
			p.SetState(143)

			var _x = p.Identifier()

			localctx.(*Array_segmentContext)._identifier = _x
		}
		localctx.(*Array_segmentContext).vars = append(localctx.(*Array_segmentContext).vars, localctx.(*Array_segmentContext)._identifier)

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(149)
		p.Match(Algol60V2ParserT__2)
	}
	{
		p.SetState(150)

		var _x = p.Number()

		localctx.(*Array_segmentContext)._number = _x
	}
	localctx.(*Array_segmentContext).sizes = append(localctx.(*Array_segmentContext).sizes, localctx.(*Array_segmentContext)._number)
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Algol60V2ParserT__1 {
		{
			p.SetState(151)
			p.Match(Algol60V2ParserT__1)
		}
		{
			p.SetState(152)

			var _x = p.Number()

			localctx.(*Array_segmentContext)._number = _x
		}
		localctx.(*Array_segmentContext).sizes = append(localctx.(*Array_segmentContext).sizes, localctx.(*Array_segmentContext)._number)

		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(158)
		p.Match(Algol60V2ParserT__3)
	}

	return localctx
}

// IString_declarationContext is an interface to support dynamic dispatch.
type IString_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_string_segment returns the _string_segment rule contexts.
	Get_string_segment() IString_segmentContext

	// Set_string_segment sets the _string_segment rule contexts.
	Set_string_segment(IString_segmentContext)

	// GetSegments returns the segments rule context list.
	GetSegments() []IString_segmentContext

	// SetSegments sets the segments rule context list.
	SetSegments([]IString_segmentContext)

	// IsString_declarationContext differentiates from other interfaces.
	IsString_declarationContext()
}

type String_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	_string_segment IString_segmentContext
	segments        []IString_segmentContext
}

func NewEmptyString_declarationContext() *String_declarationContext {
	var p = new(String_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_string_declaration
	return p
}

func (*String_declarationContext) IsString_declarationContext() {}

func NewString_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_declarationContext {
	var p = new(String_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_string_declaration

	return p
}

func (s *String_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *String_declarationContext) Get_string_segment() IString_segmentContext {
	return s._string_segment
}

func (s *String_declarationContext) Set_string_segment(v IString_segmentContext) {
	s._string_segment = v
}

func (s *String_declarationContext) GetSegments() []IString_segmentContext { return s.segments }

func (s *String_declarationContext) SetSegments(v []IString_segmentContext) { s.segments = v }

func (s *String_declarationContext) STRING() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserSTRING, 0)
}

func (s *String_declarationContext) AllString_segment() []IString_segmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IString_segmentContext); ok {
			len++
		}
	}

	tst := make([]IString_segmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IString_segmentContext); ok {
			tst[i] = t.(IString_segmentContext)
			i++
		}
	}

	return tst
}

func (s *String_declarationContext) String_segment(i int) IString_segmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_segmentContext); ok {
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

	return t.(IString_segmentContext)
}

func (s *String_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitString_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) String_declaration() (localctx IString_declarationContext) {
	this := p
	_ = this

	localctx = NewString_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, Algol60V2ParserRULE_string_declaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(Algol60V2ParserSTRING)
	}
	{
		p.SetState(161)

		var _x = p.String_segment()

		localctx.(*String_declarationContext)._string_segment = _x
	}
	localctx.(*String_declarationContext).segments = append(localctx.(*String_declarationContext).segments, localctx.(*String_declarationContext)._string_segment)
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Algol60V2ParserT__1 {
		{
			p.SetState(162)
			p.Match(Algol60V2ParserT__1)
		}
		{
			p.SetState(163)

			var _x = p.String_segment()

			localctx.(*String_declarationContext)._string_segment = _x
		}
		localctx.(*String_declarationContext).segments = append(localctx.(*String_declarationContext).segments, localctx.(*String_declarationContext)._string_segment)

		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IString_segmentContext is an interface to support dynamic dispatch.
type IString_segmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_identifier returns the _identifier rule contexts.
	Get_identifier() IIdentifierContext

	// GetSize returns the size rule contexts.
	GetSize() INumberContext

	// Set_identifier sets the _identifier rule contexts.
	Set_identifier(IIdentifierContext)

	// SetSize sets the size rule contexts.
	SetSize(INumberContext)

	// GetVars returns the vars rule context list.
	GetVars() []IIdentifierContext

	// SetVars sets the vars rule context list.
	SetVars([]IIdentifierContext)

	// IsString_segmentContext differentiates from other interfaces.
	IsString_segmentContext()
}

type String_segmentContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	_identifier IIdentifierContext
	vars        []IIdentifierContext
	size        INumberContext
}

func NewEmptyString_segmentContext() *String_segmentContext {
	var p = new(String_segmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_string_segment
	return p
}

func (*String_segmentContext) IsString_segmentContext() {}

func NewString_segmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_segmentContext {
	var p = new(String_segmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_string_segment

	return p
}

func (s *String_segmentContext) GetParser() antlr.Parser { return s.parser }

func (s *String_segmentContext) Get_identifier() IIdentifierContext { return s._identifier }

func (s *String_segmentContext) GetSize() INumberContext { return s.size }

func (s *String_segmentContext) Set_identifier(v IIdentifierContext) { s._identifier = v }

func (s *String_segmentContext) SetSize(v INumberContext) { s.size = v }

func (s *String_segmentContext) GetVars() []IIdentifierContext { return s.vars }

func (s *String_segmentContext) SetVars(v []IIdentifierContext) { s.vars = v }

func (s *String_segmentContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *String_segmentContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
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

	return t.(IIdentifierContext)
}

func (s *String_segmentContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *String_segmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_segmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_segmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitString_segment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) String_segment() (localctx IString_segmentContext) {
	this := p
	_ = this

	localctx = NewString_segmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, Algol60V2ParserRULE_string_segment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)

		var _x = p.Identifier()

		localctx.(*String_segmentContext)._identifier = _x
	}
	localctx.(*String_segmentContext).vars = append(localctx.(*String_segmentContext).vars, localctx.(*String_segmentContext)._identifier)
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Algol60V2ParserT__1 {
		{
			p.SetState(170)
			p.Match(Algol60V2ParserT__1)
		}
		{
			p.SetState(171)

			var _x = p.Identifier()

			localctx.(*String_segmentContext)._identifier = _x
		}
		localctx.(*String_segmentContext).vars = append(localctx.(*String_segmentContext).vars, localctx.(*String_segmentContext)._identifier)

		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(177)
		p.Match(Algol60V2ParserT__2)
	}
	{
		p.SetState(178)

		var _x = p.Number()

		localctx.(*String_segmentContext).size = _x
	}
	{
		p.SetState(179)
		p.Match(Algol60V2ParserT__3)
	}

	return localctx
}

// IProcedure_declarationContext is an interface to support dynamic dispatch.
type IProcedure_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetReturnType returns the returnType token.
	GetReturnType() antlr.Token

	// SetReturnType sets the returnType token.
	SetReturnType(antlr.Token)

	// GetParams returns the params rule contexts.
	GetParams() IFormal_parameter_listContext

	// Get_declaration returns the _declaration rule contexts.
	Get_declaration() IDeclarationContext

	// SetParams sets the params rule contexts.
	SetParams(IFormal_parameter_listContext)

	// Set_declaration sets the _declaration rule contexts.
	Set_declaration(IDeclarationContext)

	// GetDecls returns the decls rule context list.
	GetDecls() []IDeclarationContext

	// SetDecls sets the decls rule context list.
	SetDecls([]IDeclarationContext)

	// IsProcedure_declarationContext differentiates from other interfaces.
	IsProcedure_declarationContext()
}

type Procedure_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	returnType   antlr.Token
	params       IFormal_parameter_listContext
	_declaration IDeclarationContext
	decls        []IDeclarationContext
}

func NewEmptyProcedure_declarationContext() *Procedure_declarationContext {
	var p = new(Procedure_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_procedure_declaration
	return p
}

func (*Procedure_declarationContext) IsProcedure_declarationContext() {}

func NewProcedure_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Procedure_declarationContext {
	var p = new(Procedure_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_procedure_declaration

	return p
}

func (s *Procedure_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Procedure_declarationContext) GetReturnType() antlr.Token { return s.returnType }

func (s *Procedure_declarationContext) SetReturnType(v antlr.Token) { s.returnType = v }

func (s *Procedure_declarationContext) GetParams() IFormal_parameter_listContext { return s.params }

func (s *Procedure_declarationContext) Get_declaration() IDeclarationContext { return s._declaration }

func (s *Procedure_declarationContext) SetParams(v IFormal_parameter_listContext) { s.params = v }

func (s *Procedure_declarationContext) Set_declaration(v IDeclarationContext) { s._declaration = v }

func (s *Procedure_declarationContext) GetDecls() []IDeclarationContext { return s.decls }

func (s *Procedure_declarationContext) SetDecls(v []IDeclarationContext) { s.decls = v }

func (s *Procedure_declarationContext) PROCEDURE() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserPROCEDURE, 0)
}

func (s *Procedure_declarationContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Procedure_declarationContext) Statement() IStatementContext {
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

func (s *Procedure_declarationContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserINTEGER, 0)
}

func (s *Procedure_declarationContext) Formal_parameter_list() IFormal_parameter_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormal_parameter_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormal_parameter_listContext)
}

func (s *Procedure_declarationContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *Procedure_declarationContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
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

	return t.(IDeclarationContext)
}

func (s *Procedure_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Procedure_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Procedure_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitProcedure_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Procedure_declaration() (localctx IProcedure_declarationContext) {
	this := p
	_ = this

	localctx = NewProcedure_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, Algol60V2ParserRULE_procedure_declaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Algol60V2ParserINTEGER {
		{
			p.SetState(181)

			var _m = p.Match(Algol60V2ParserINTEGER)

			localctx.(*Procedure_declarationContext).returnType = _m
		}

	}
	{
		p.SetState(184)
		p.Match(Algol60V2ParserPROCEDURE)
	}
	{
		p.SetState(185)
		p.Identifier()
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Algol60V2ParserT__4 {
		{
			p.SetState(186)
			p.Match(Algol60V2ParserT__4)
		}
		{
			p.SetState(187)

			var _x = p.Formal_parameter_list()

			localctx.(*Procedure_declarationContext).params = _x
		}
		{
			p.SetState(188)
			p.Match(Algol60V2ParserT__5)
		}

	}
	{
		p.SetState(192)
		p.Match(Algol60V2ParserT__0)
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(193)

				var _x = p.Declaration()

				localctx.(*Procedure_declarationContext)._declaration = _x
			}
			localctx.(*Procedure_declarationContext).decls = append(localctx.(*Procedure_declarationContext).decls, localctx.(*Procedure_declarationContext)._declaration)
			{
				p.SetState(194)
				p.Match(Algol60V2ParserT__0)
			}

		}
		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}
	{
		p.SetState(201)
		p.Statement()
	}

	return localctx
}

// IFormal_parameter_listContext is an interface to support dynamic dispatch.
type IFormal_parameter_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormal_parameter_listContext differentiates from other interfaces.
	IsFormal_parameter_listContext()
}

type Formal_parameter_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormal_parameter_listContext() *Formal_parameter_listContext {
	var p = new(Formal_parameter_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_formal_parameter_list
	return p
}

func (*Formal_parameter_listContext) IsFormal_parameter_listContext() {}

func NewFormal_parameter_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Formal_parameter_listContext {
	var p = new(Formal_parameter_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_formal_parameter_list

	return p
}

func (s *Formal_parameter_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Formal_parameter_listContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Formal_parameter_listContext) Parameter_delimiter() IParameter_delimiterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameter_delimiterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameter_delimiterContext)
}

func (s *Formal_parameter_listContext) Formal_parameter_list() IFormal_parameter_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormal_parameter_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormal_parameter_listContext)
}

func (s *Formal_parameter_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Formal_parameter_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Formal_parameter_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitFormal_parameter_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Formal_parameter_list() (localctx IFormal_parameter_listContext) {
	this := p
	_ = this

	localctx = NewFormal_parameter_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, Algol60V2ParserRULE_formal_parameter_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Identifier()
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(204)
			p.Parameter_delimiter()
		}
		{
			p.SetState(205)
			p.Formal_parameter_list()
		}

	}

	return localctx
}

// IParameter_delimiterContext is an interface to support dynamic dispatch.
type IParameter_delimiterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameter_delimiterContext differentiates from other interfaces.
	IsParameter_delimiterContext()
}

type Parameter_delimiterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameter_delimiterContext() *Parameter_delimiterContext {
	var p = new(Parameter_delimiterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_parameter_delimiter
	return p
}

func (*Parameter_delimiterContext) IsParameter_delimiterContext() {}

func NewParameter_delimiterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parameter_delimiterContext {
	var p = new(Parameter_delimiterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_parameter_delimiter

	return p
}

func (s *Parameter_delimiterContext) GetParser() antlr.Parser { return s.parser }

func (s *Parameter_delimiterContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserIDENTIFIER, 0)
}

func (s *Parameter_delimiterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parameter_delimiterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Parameter_delimiterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitParameter_delimiter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Parameter_delimiter() (localctx IParameter_delimiterContext) {
	this := p
	_ = this

	localctx = NewParameter_delimiterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, Algol60V2ParserRULE_parameter_delimiter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(214)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Algol60V2ParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(209)
			p.Match(Algol60V2ParserT__1)
		}

	case Algol60V2ParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(210)
			p.Match(Algol60V2ParserT__5)
		}
		{
			p.SetState(211)
			p.Match(Algol60V2ParserIDENTIFIER)
		}
		{
			p.SetState(212)
			p.Match(Algol60V2ParserT__6)
		}
		{
			p.SetState(213)
			p.Match(Algol60V2ParserT__4)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Unconditional_statement() IUnconditional_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnconditional_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnconditional_statementContext)
}

func (s *StatementContext) Conditional_statement() IConditional_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditional_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditional_statementContext)
}

func (s *StatementContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserCOMMENT, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, Algol60V2ParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(219)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Algol60V2ParserBEGIN, Algol60V2ParserFOR, Algol60V2ParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(216)
			p.Unconditional_statement()
		}

	case Algol60V2ParserIF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.Conditional_statement()
		}

	case Algol60V2ParserCOMMENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(218)
			p.Match(Algol60V2ParserCOMMENT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IUnconditional_statementContext is an interface to support dynamic dispatch.
type IUnconditional_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnconditional_statementContext differentiates from other interfaces.
	IsUnconditional_statementContext()
}

type Unconditional_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnconditional_statementContext() *Unconditional_statementContext {
	var p = new(Unconditional_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_unconditional_statement
	return p
}

func (*Unconditional_statementContext) IsUnconditional_statementContext() {}

func NewUnconditional_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unconditional_statementContext {
	var p = new(Unconditional_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_unconditional_statement

	return p
}

func (s *Unconditional_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Unconditional_statementContext) For_statement() IFor_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_statementContext)
}

func (s *Unconditional_statementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Unconditional_statementContext) Assignment_statement() IAssignment_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignment_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignment_statementContext)
}

func (s *Unconditional_statementContext) Procedure_statement() IProcedure_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProcedure_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProcedure_statementContext)
}

func (s *Unconditional_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unconditional_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unconditional_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitUnconditional_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Unconditional_statement() (localctx IUnconditional_statementContext) {
	this := p
	_ = this

	localctx = NewUnconditional_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, Algol60V2ParserRULE_unconditional_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(221)
			p.For_statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(222)
			p.Block()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(223)
			p.Assignment_statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(224)
			p.Procedure_statement()
		}

	}

	return localctx
}

// IAssignment_statementContext is an interface to support dynamic dispatch.
type IAssignment_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTarget returns the target rule contexts.
	GetTarget() ILeft_partContext

	// SetTarget sets the target rule contexts.
	SetTarget(ILeft_partContext)

	// IsAssignment_statementContext differentiates from other interfaces.
	IsAssignment_statementContext()
}

type Assignment_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	target ILeft_partContext
}

func NewEmptyAssignment_statementContext() *Assignment_statementContext {
	var p = new(Assignment_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_assignment_statement
	return p
}

func (*Assignment_statementContext) IsAssignment_statementContext() {}

func NewAssignment_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assignment_statementContext {
	var p = new(Assignment_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_assignment_statement

	return p
}

func (s *Assignment_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Assignment_statementContext) GetTarget() ILeft_partContext { return s.target }

func (s *Assignment_statementContext) SetTarget(v ILeft_partContext) { s.target = v }

func (s *Assignment_statementContext) Expression() IExpressionContext {
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

func (s *Assignment_statementContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserASSIGN, 0)
}

func (s *Assignment_statementContext) Left_part() ILeft_partContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeft_partContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeft_partContext)
}

func (s *Assignment_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assignment_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assignment_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitAssignment_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Assignment_statement() (localctx IAssignment_statementContext) {
	this := p
	_ = this

	localctx = NewAssignment_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, Algol60V2ParserRULE_assignment_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)

		var _x = p.Left_part()

		localctx.(*Assignment_statementContext).target = _x
	}
	{
		p.SetState(228)
		p.Match(Algol60V2ParserASSIGN)
	}

	{
		p.SetState(230)
		p.Expression()
	}

	return localctx
}

// ILeft_partContext is an interface to support dynamic dispatch.
type ILeft_partContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLeft_partContext differentiates from other interfaces.
	IsLeft_partContext()
}

type Left_partContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeft_partContext() *Left_partContext {
	var p = new(Left_partContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_left_part
	return p
}

func (*Left_partContext) IsLeft_partContext() {}

func NewLeft_partContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Left_partContext {
	var p = new(Left_partContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_left_part

	return p
}

func (s *Left_partContext) GetParser() antlr.Parser { return s.parser }

func (s *Left_partContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Left_partContext) Subscripted_variable() ISubscripted_variableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubscripted_variableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubscripted_variableContext)
}

func (s *Left_partContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Left_partContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Left_partContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitLeft_part(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Left_part() (localctx ILeft_partContext) {
	this := p
	_ = this

	localctx = NewLeft_partContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, Algol60V2ParserRULE_left_part)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(232)
			p.Identifier()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(233)
			p.Subscripted_variable()
		}

	}

	return localctx
}

// IProcedure_statementContext is an interface to support dynamic dispatch.
type IProcedure_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetProc_name returns the proc_name rule contexts.
	GetProc_name() IIdentifierContext

	// GetArgs returns the args rule contexts.
	GetArgs() IActual_parameter_listContext

	// SetProc_name sets the proc_name rule contexts.
	SetProc_name(IIdentifierContext)

	// SetArgs sets the args rule contexts.
	SetArgs(IActual_parameter_listContext)

	// IsProcedure_statementContext differentiates from other interfaces.
	IsProcedure_statementContext()
}

type Procedure_statementContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	proc_name IIdentifierContext
	args      IActual_parameter_listContext
}

func NewEmptyProcedure_statementContext() *Procedure_statementContext {
	var p = new(Procedure_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_procedure_statement
	return p
}

func (*Procedure_statementContext) IsProcedure_statementContext() {}

func NewProcedure_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Procedure_statementContext {
	var p = new(Procedure_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_procedure_statement

	return p
}

func (s *Procedure_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Procedure_statementContext) GetProc_name() IIdentifierContext { return s.proc_name }

func (s *Procedure_statementContext) GetArgs() IActual_parameter_listContext { return s.args }

func (s *Procedure_statementContext) SetProc_name(v IIdentifierContext) { s.proc_name = v }

func (s *Procedure_statementContext) SetArgs(v IActual_parameter_listContext) { s.args = v }

func (s *Procedure_statementContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Procedure_statementContext) Actual_parameter_list() IActual_parameter_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActual_parameter_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActual_parameter_listContext)
}

func (s *Procedure_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Procedure_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Procedure_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitProcedure_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Procedure_statement() (localctx IProcedure_statementContext) {
	this := p
	_ = this

	localctx = NewProcedure_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, Algol60V2ParserRULE_procedure_statement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)

		var _x = p.Identifier()

		localctx.(*Procedure_statementContext).proc_name = _x
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Algol60V2ParserT__4 {
		{
			p.SetState(237)
			p.Match(Algol60V2ParserT__4)
		}
		{
			p.SetState(238)

			var _x = p.Actual_parameter_list()

			localctx.(*Procedure_statementContext).args = _x
		}
		{
			p.SetState(239)
			p.Match(Algol60V2ParserT__5)
		}

	}

	return localctx
}

// IActual_parameter_listContext is an interface to support dynamic dispatch.
type IActual_parameter_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActual_parameter_listContext differentiates from other interfaces.
	IsActual_parameter_listContext()
}

type Actual_parameter_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActual_parameter_listContext() *Actual_parameter_listContext {
	var p = new(Actual_parameter_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_actual_parameter_list
	return p
}

func (*Actual_parameter_listContext) IsActual_parameter_listContext() {}

func NewActual_parameter_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Actual_parameter_listContext {
	var p = new(Actual_parameter_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_actual_parameter_list

	return p
}

func (s *Actual_parameter_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Actual_parameter_listContext) Expression() IExpressionContext {
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

func (s *Actual_parameter_listContext) Parameter_delimiter() IParameter_delimiterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameter_delimiterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameter_delimiterContext)
}

func (s *Actual_parameter_listContext) Actual_parameter_list() IActual_parameter_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActual_parameter_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActual_parameter_listContext)
}

func (s *Actual_parameter_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Actual_parameter_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Actual_parameter_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitActual_parameter_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Actual_parameter_list() (localctx IActual_parameter_listContext) {
	this := p
	_ = this

	localctx = NewActual_parameter_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, Algol60V2ParserRULE_actual_parameter_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Expression()
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(244)
			p.Parameter_delimiter()
		}
		{
			p.SetState(245)
			p.Actual_parameter_list()
		}

	}

	return localctx
}

// IConditional_statementContext is an interface to support dynamic dispatch.
type IConditional_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetThenBlock returns the thenBlock rule contexts.
	GetThenBlock() IUnconditional_statementContext

	// GetElseBlock returns the elseBlock rule contexts.
	GetElseBlock() IStatementContext

	// SetThenBlock sets the thenBlock rule contexts.
	SetThenBlock(IUnconditional_statementContext)

	// SetElseBlock sets the elseBlock rule contexts.
	SetElseBlock(IStatementContext)

	// IsConditional_statementContext differentiates from other interfaces.
	IsConditional_statementContext()
}

type Conditional_statementContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	thenBlock IUnconditional_statementContext
	elseBlock IStatementContext
}

func NewEmptyConditional_statementContext() *Conditional_statementContext {
	var p = new(Conditional_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_conditional_statement
	return p
}

func (*Conditional_statementContext) IsConditional_statementContext() {}

func NewConditional_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Conditional_statementContext {
	var p = new(Conditional_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_conditional_statement

	return p
}

func (s *Conditional_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Conditional_statementContext) GetThenBlock() IUnconditional_statementContext {
	return s.thenBlock
}

func (s *Conditional_statementContext) GetElseBlock() IStatementContext { return s.elseBlock }

func (s *Conditional_statementContext) SetThenBlock(v IUnconditional_statementContext) {
	s.thenBlock = v
}

func (s *Conditional_statementContext) SetElseBlock(v IStatementContext) { s.elseBlock = v }

func (s *Conditional_statementContext) IF() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserIF, 0)
}

func (s *Conditional_statementContext) Boolean_expression() IBoolean_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolean_expressionContext)
}

func (s *Conditional_statementContext) THEN() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserTHEN, 0)
}

func (s *Conditional_statementContext) Unconditional_statement() IUnconditional_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnconditional_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnconditional_statementContext)
}

func (s *Conditional_statementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserELSE, 0)
}

func (s *Conditional_statementContext) Statement() IStatementContext {
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

func (s *Conditional_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Conditional_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Conditional_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitConditional_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Conditional_statement() (localctx IConditional_statementContext) {
	this := p
	_ = this

	localctx = NewConditional_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, Algol60V2ParserRULE_conditional_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.Match(Algol60V2ParserIF)
	}
	{
		p.SetState(250)
		p.Boolean_expression()
	}
	{
		p.SetState(251)
		p.Match(Algol60V2ParserTHEN)
	}
	{
		p.SetState(252)

		var _x = p.Unconditional_statement()

		localctx.(*Conditional_statementContext).thenBlock = _x
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(253)
			p.Match(Algol60V2ParserELSE)
		}
		{
			p.SetState(254)

			var _x = p.Statement()

			localctx.(*Conditional_statementContext).elseBlock = _x
		}

	}

	return localctx
}

// IFor_statementContext is an interface to support dynamic dispatch.
type IFor_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLoopVar returns the loopVar rule contexts.
	GetLoopVar() IIdentifierContext

	// GetBody returns the body rule contexts.
	GetBody() IStatementContext

	// SetLoopVar sets the loopVar rule contexts.
	SetLoopVar(IIdentifierContext)

	// SetBody sets the body rule contexts.
	SetBody(IStatementContext)

	// IsFor_statementContext differentiates from other interfaces.
	IsFor_statementContext()
}

type For_statementContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	loopVar IIdentifierContext
	body    IStatementContext
}

func NewEmptyFor_statementContext() *For_statementContext {
	var p = new(For_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_for_statement
	return p
}

func (*For_statementContext) IsFor_statementContext() {}

func NewFor_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_statementContext {
	var p = new(For_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_for_statement

	return p
}

func (s *For_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *For_statementContext) GetLoopVar() IIdentifierContext { return s.loopVar }

func (s *For_statementContext) GetBody() IStatementContext { return s.body }

func (s *For_statementContext) SetLoopVar(v IIdentifierContext) { s.loopVar = v }

func (s *For_statementContext) SetBody(v IStatementContext) { s.body = v }

func (s *For_statementContext) FOR() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserFOR, 0)
}

func (s *For_statementContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserASSIGN, 0)
}

func (s *For_statementContext) For_list() IFor_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_listContext)
}

func (s *For_statementContext) DO() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserDO, 0)
}

func (s *For_statementContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *For_statementContext) Statement() IStatementContext {
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

func (s *For_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitFor_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) For_statement() (localctx IFor_statementContext) {
	this := p
	_ = this

	localctx = NewFor_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, Algol60V2ParserRULE_for_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		p.Match(Algol60V2ParserFOR)
	}
	{
		p.SetState(258)

		var _x = p.Identifier()

		localctx.(*For_statementContext).loopVar = _x
	}
	{
		p.SetState(259)
		p.Match(Algol60V2ParserASSIGN)
	}
	{
		p.SetState(260)
		p.For_list()
	}
	{
		p.SetState(261)
		p.Match(Algol60V2ParserDO)
	}
	{
		p.SetState(262)

		var _x = p.Statement()

		localctx.(*For_statementContext).body = _x
	}

	return localctx
}

// IFor_listContext is an interface to support dynamic dispatch.
type IFor_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_for_list_element returns the _for_list_element rule contexts.
	Get_for_list_element() IFor_list_elementContext

	// Set_for_list_element sets the _for_list_element rule contexts.
	Set_for_list_element(IFor_list_elementContext)

	// GetElems returns the elems rule context list.
	GetElems() []IFor_list_elementContext

	// SetElems sets the elems rule context list.
	SetElems([]IFor_list_elementContext)

	// IsFor_listContext differentiates from other interfaces.
	IsFor_listContext()
}

type For_listContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	_for_list_element IFor_list_elementContext
	elems             []IFor_list_elementContext
}

func NewEmptyFor_listContext() *For_listContext {
	var p = new(For_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_for_list
	return p
}

func (*For_listContext) IsFor_listContext() {}

func NewFor_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_listContext {
	var p = new(For_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_for_list

	return p
}

func (s *For_listContext) GetParser() antlr.Parser { return s.parser }

func (s *For_listContext) Get_for_list_element() IFor_list_elementContext { return s._for_list_element }

func (s *For_listContext) Set_for_list_element(v IFor_list_elementContext) { s._for_list_element = v }

func (s *For_listContext) GetElems() []IFor_list_elementContext { return s.elems }

func (s *For_listContext) SetElems(v []IFor_list_elementContext) { s.elems = v }

func (s *For_listContext) AllFor_list_element() []IFor_list_elementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFor_list_elementContext); ok {
			len++
		}
	}

	tst := make([]IFor_list_elementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFor_list_elementContext); ok {
			tst[i] = t.(IFor_list_elementContext)
			i++
		}
	}

	return tst
}

func (s *For_listContext) For_list_element(i int) IFor_list_elementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_list_elementContext); ok {
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

	return t.(IFor_list_elementContext)
}

func (s *For_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitFor_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) For_list() (localctx IFor_listContext) {
	this := p
	_ = this

	localctx = NewFor_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, Algol60V2ParserRULE_for_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)

		var _x = p.For_list_element()

		localctx.(*For_listContext)._for_list_element = _x
	}
	localctx.(*For_listContext).elems = append(localctx.(*For_listContext).elems, localctx.(*For_listContext)._for_list_element)
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Algol60V2ParserT__1 {
		{
			p.SetState(265)
			p.Match(Algol60V2ParserT__1)
		}
		{
			p.SetState(266)

			var _x = p.For_list_element()

			localctx.(*For_listContext)._for_list_element = _x
		}
		localctx.(*For_listContext).elems = append(localctx.(*For_listContext).elems, localctx.(*For_listContext)._for_list_element)

		p.SetState(271)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFor_list_elementContext is an interface to support dynamic dispatch.
type IFor_list_elementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStepNeg returns the stepNeg token.
	GetStepNeg() antlr.Token

	// SetStepNeg sets the stepNeg token.
	SetStepNeg(antlr.Token)

	// GetFirstExpr returns the firstExpr rule contexts.
	GetFirstExpr() IExpressionContext

	// GetStep returns the step rule contexts.
	GetStep() INumberContext

	// GetUntil returns the until rule contexts.
	GetUntil() IExpressionContext

	// GetWhileCond returns the whileCond rule contexts.
	GetWhileCond() IBoolean_expressionContext

	// SetFirstExpr sets the firstExpr rule contexts.
	SetFirstExpr(IExpressionContext)

	// SetStep sets the step rule contexts.
	SetStep(INumberContext)

	// SetUntil sets the until rule contexts.
	SetUntil(IExpressionContext)

	// SetWhileCond sets the whileCond rule contexts.
	SetWhileCond(IBoolean_expressionContext)

	// IsFor_list_elementContext differentiates from other interfaces.
	IsFor_list_elementContext()
}

type For_list_elementContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	firstExpr IExpressionContext
	stepNeg   antlr.Token
	step      INumberContext
	until     IExpressionContext
	whileCond IBoolean_expressionContext
}

func NewEmptyFor_list_elementContext() *For_list_elementContext {
	var p = new(For_list_elementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_for_list_element
	return p
}

func (*For_list_elementContext) IsFor_list_elementContext() {}

func NewFor_list_elementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_list_elementContext {
	var p = new(For_list_elementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_for_list_element

	return p
}

func (s *For_list_elementContext) GetParser() antlr.Parser { return s.parser }

func (s *For_list_elementContext) GetStepNeg() antlr.Token { return s.stepNeg }

func (s *For_list_elementContext) SetStepNeg(v antlr.Token) { s.stepNeg = v }

func (s *For_list_elementContext) GetFirstExpr() IExpressionContext { return s.firstExpr }

func (s *For_list_elementContext) GetStep() INumberContext { return s.step }

func (s *For_list_elementContext) GetUntil() IExpressionContext { return s.until }

func (s *For_list_elementContext) GetWhileCond() IBoolean_expressionContext { return s.whileCond }

func (s *For_list_elementContext) SetFirstExpr(v IExpressionContext) { s.firstExpr = v }

func (s *For_list_elementContext) SetStep(v INumberContext) { s.step = v }

func (s *For_list_elementContext) SetUntil(v IExpressionContext) { s.until = v }

func (s *For_list_elementContext) SetWhileCond(v IBoolean_expressionContext) { s.whileCond = v }

func (s *For_list_elementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *For_list_elementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *For_list_elementContext) STEP() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserSTEP, 0)
}

func (s *For_list_elementContext) UNTIL() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserUNTIL, 0)
}

func (s *For_list_elementContext) WHILE() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserWHILE, 0)
}

func (s *For_list_elementContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *For_list_elementContext) Boolean_expression() IBoolean_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolean_expressionContext)
}

func (s *For_list_elementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_list_elementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_list_elementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitFor_list_element(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) For_list_element() (localctx IFor_list_elementContext) {
	this := p
	_ = this

	localctx = NewFor_list_elementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, Algol60V2ParserRULE_for_list_element)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)

		var _x = p.Expression()

		localctx.(*For_list_elementContext).firstExpr = _x
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Algol60V2ParserSTEP:
		{
			p.SetState(273)
			p.Match(Algol60V2ParserSTEP)
		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == Algol60V2ParserT__7 {
			{
				p.SetState(274)

				var _m = p.Match(Algol60V2ParserT__7)

				localctx.(*For_list_elementContext).stepNeg = _m
			}

		}
		{
			p.SetState(277)

			var _x = p.Number()

			localctx.(*For_list_elementContext).step = _x
		}
		{
			p.SetState(278)
			p.Match(Algol60V2ParserUNTIL)
		}
		{
			p.SetState(279)

			var _x = p.Expression()

			localctx.(*For_list_elementContext).until = _x
		}

	case Algol60V2ParserWHILE:
		{
			p.SetState(281)
			p.Match(Algol60V2ParserWHILE)
		}
		{
			p.SetState(282)

			var _x = p.Boolean_expression()

			localctx.(*For_list_elementContext).whileCond = _x
		}

	case Algol60V2ParserT__1, Algol60V2ParserDO:

	default:
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCond returns the cond rule contexts.
	GetCond() IBoolean_expressionContext

	// GetTrueCase returns the trueCase rule contexts.
	GetTrueCase() ISimple_expressionContext

	// GetFalseCase returns the falseCase rule contexts.
	GetFalseCase() IExpressionContext

	// SetCond sets the cond rule contexts.
	SetCond(IBoolean_expressionContext)

	// SetTrueCase sets the trueCase rule contexts.
	SetTrueCase(ISimple_expressionContext)

	// SetFalseCase sets the falseCase rule contexts.
	SetFalseCase(IExpressionContext)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	cond      IBoolean_expressionContext
	trueCase  ISimple_expressionContext
	falseCase IExpressionContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetCond() IBoolean_expressionContext { return s.cond }

func (s *ExpressionContext) GetTrueCase() ISimple_expressionContext { return s.trueCase }

func (s *ExpressionContext) GetFalseCase() IExpressionContext { return s.falseCase }

func (s *ExpressionContext) SetCond(v IBoolean_expressionContext) { s.cond = v }

func (s *ExpressionContext) SetTrueCase(v ISimple_expressionContext) { s.trueCase = v }

func (s *ExpressionContext) SetFalseCase(v IExpressionContext) { s.falseCase = v }

func (s *ExpressionContext) Simple_expression() ISimple_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_expressionContext)
}

func (s *ExpressionContext) IF() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserIF, 0)
}

func (s *ExpressionContext) THEN() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserTHEN, 0)
}

func (s *ExpressionContext) ELSE() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserELSE, 0)
}

func (s *ExpressionContext) Boolean_expression() IBoolean_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolean_expressionContext)
}

func (s *ExpressionContext) Expression() IExpressionContext {
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

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, Algol60V2ParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(293)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Algol60V2ParserT__4, Algol60V2ParserT__7, Algol60V2ParserT__8, Algol60V2ParserQUOTED_STRING, Algol60V2ParserNUMBER, Algol60V2ParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(285)
			p.Simple_expression()
		}

	case Algol60V2ParserIF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(286)
			p.Match(Algol60V2ParserIF)
		}
		{
			p.SetState(287)

			var _x = p.Boolean_expression()

			localctx.(*ExpressionContext).cond = _x
		}
		{
			p.SetState(288)
			p.Match(Algol60V2ParserTHEN)
		}
		{
			p.SetState(289)

			var _x = p.Simple_expression()

			localctx.(*ExpressionContext).trueCase = _x
		}
		{
			p.SetState(290)
			p.Match(Algol60V2ParserELSE)
		}
		{
			p.SetState(291)

			var _x = p.Expression()

			localctx.(*ExpressionContext).falseCase = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISimple_expressionContext is an interface to support dynamic dispatch.
type ISimple_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetS9 returns the s9 token.
	GetS9() antlr.Token

	// GetS8 returns the s8 token.
	GetS8() antlr.Token

	// Get_tset542 returns the _tset542 token.
	Get_tset542() antlr.Token

	// Get_tset562 returns the _tset562 token.
	Get_tset562() antlr.Token

	// SetS9 sets the s9 token.
	SetS9(antlr.Token)

	// SetS8 sets the s8 token.
	SetS8(antlr.Token)

	// Set_tset542 sets the _tset542 token.
	Set_tset542(antlr.Token)

	// Set_tset562 sets the _tset562 token.
	Set_tset562(antlr.Token)

	// GetOperators returns the operators token list.
	GetOperators() []antlr.Token

	// SetOperators sets the operators token list.
	SetOperators([]antlr.Token)

	// Get_term returns the _term rule contexts.
	Get_term() ITermContext

	// Set_term sets the _term rule contexts.
	Set_term(ITermContext)

	// GetOperands returns the operands rule context list.
	GetOperands() []ITermContext

	// SetOperands sets the operands rule context list.
	SetOperands([]ITermContext)

	// IsSimple_expressionContext differentiates from other interfaces.
	IsSimple_expressionContext()
}

type Simple_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	s9        antlr.Token
	operators []antlr.Token
	s8        antlr.Token
	_tset542  antlr.Token
	_term     ITermContext
	operands  []ITermContext
	_tset562  antlr.Token
}

func NewEmptySimple_expressionContext() *Simple_expressionContext {
	var p = new(Simple_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_simple_expression
	return p
}

func (*Simple_expressionContext) IsSimple_expressionContext() {}

func NewSimple_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_expressionContext {
	var p = new(Simple_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_simple_expression

	return p
}

func (s *Simple_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_expressionContext) GetS9() antlr.Token { return s.s9 }

func (s *Simple_expressionContext) GetS8() antlr.Token { return s.s8 }

func (s *Simple_expressionContext) Get_tset542() antlr.Token { return s._tset542 }

func (s *Simple_expressionContext) Get_tset562() antlr.Token { return s._tset562 }

func (s *Simple_expressionContext) SetS9(v antlr.Token) { s.s9 = v }

func (s *Simple_expressionContext) SetS8(v antlr.Token) { s.s8 = v }

func (s *Simple_expressionContext) Set_tset542(v antlr.Token) { s._tset542 = v }

func (s *Simple_expressionContext) Set_tset562(v antlr.Token) { s._tset562 = v }

func (s *Simple_expressionContext) GetOperators() []antlr.Token { return s.operators }

func (s *Simple_expressionContext) SetOperators(v []antlr.Token) { s.operators = v }

func (s *Simple_expressionContext) Get_term() ITermContext { return s._term }

func (s *Simple_expressionContext) Set_term(v ITermContext) { s._term = v }

func (s *Simple_expressionContext) GetOperands() []ITermContext { return s.operands }

func (s *Simple_expressionContext) SetOperands(v []ITermContext) { s.operands = v }

func (s *Simple_expressionContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *Simple_expressionContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
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

	return t.(ITermContext)
}

func (s *Simple_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitSimple_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Simple_expression() (localctx ISimple_expressionContext) {
	this := p
	_ = this

	localctx = NewSimple_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, Algol60V2ParserRULE_simple_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Algol60V2ParserT__7 || _la == Algol60V2ParserT__8 {
		{
			p.SetState(295)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Simple_expressionContext)._tset542 = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == Algol60V2ParserT__7 || _la == Algol60V2ParserT__8) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Simple_expressionContext)._tset542 = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		localctx.(*Simple_expressionContext).operators = append(localctx.(*Simple_expressionContext).operators, localctx.(*Simple_expressionContext)._tset542)

	}
	{
		p.SetState(298)

		var _x = p.Term()

		localctx.(*Simple_expressionContext)._term = _x
	}
	localctx.(*Simple_expressionContext).operands = append(localctx.(*Simple_expressionContext).operands, localctx.(*Simple_expressionContext)._term)
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Algol60V2ParserT__7 || _la == Algol60V2ParserT__8 {
		{
			p.SetState(299)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Simple_expressionContext)._tset562 = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == Algol60V2ParserT__7 || _la == Algol60V2ParserT__8) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Simple_expressionContext)._tset562 = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		localctx.(*Simple_expressionContext).operators = append(localctx.(*Simple_expressionContext).operators, localctx.(*Simple_expressionContext)._tset562)
		{
			p.SetState(300)

			var _x = p.Term()

			localctx.(*Simple_expressionContext)._term = _x
		}
		localctx.(*Simple_expressionContext).operands = append(localctx.(*Simple_expressionContext).operands, localctx.(*Simple_expressionContext)._term)

		p.SetState(305)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetS10 returns the s10 token.
	GetS10() antlr.Token

	// GetS11 returns the s11 token.
	GetS11() antlr.Token

	// Get_tset592 returns the _tset592 token.
	Get_tset592() antlr.Token

	// SetS10 sets the s10 token.
	SetS10(antlr.Token)

	// SetS11 sets the s11 token.
	SetS11(antlr.Token)

	// Set_tset592 sets the _tset592 token.
	Set_tset592(antlr.Token)

	// GetOperators returns the operators token list.
	GetOperators() []antlr.Token

	// SetOperators sets the operators token list.
	SetOperators([]antlr.Token)

	// Get_factor returns the _factor rule contexts.
	Get_factor() IFactorContext

	// Set_factor sets the _factor rule contexts.
	Set_factor(IFactorContext)

	// GetOperands returns the operands rule context list.
	GetOperands() []IFactorContext

	// SetOperands sets the operands rule context list.
	SetOperands([]IFactorContext)

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	_factor   IFactorContext
	operands  []IFactorContext
	s10       antlr.Token
	operators []antlr.Token
	s11       antlr.Token
	_tset592  antlr.Token
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) GetS10() antlr.Token { return s.s10 }

func (s *TermContext) GetS11() antlr.Token { return s.s11 }

func (s *TermContext) Get_tset592() antlr.Token { return s._tset592 }

func (s *TermContext) SetS10(v antlr.Token) { s.s10 = v }

func (s *TermContext) SetS11(v antlr.Token) { s.s11 = v }

func (s *TermContext) Set_tset592(v antlr.Token) { s._tset592 = v }

func (s *TermContext) GetOperators() []antlr.Token { return s.operators }

func (s *TermContext) SetOperators(v []antlr.Token) { s.operators = v }

func (s *TermContext) Get_factor() IFactorContext { return s._factor }

func (s *TermContext) Set_factor(v IFactorContext) { s._factor = v }

func (s *TermContext) GetOperands() []IFactorContext { return s.operands }

func (s *TermContext) SetOperands(v []IFactorContext) { s.operands = v }

func (s *TermContext) AllFactor() []IFactorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFactorContext); ok {
			len++
		}
	}

	tst := make([]IFactorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFactorContext); ok {
			tst[i] = t.(IFactorContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) Factor(i int) IFactorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
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

	return t.(IFactorContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, Algol60V2ParserRULE_term)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)

		var _x = p.Factor()

		localctx.(*TermContext)._factor = _x
	}
	localctx.(*TermContext).operands = append(localctx.(*TermContext).operands, localctx.(*TermContext)._factor)
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Algol60V2ParserT__9 || _la == Algol60V2ParserT__10 {
		{
			p.SetState(307)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*TermContext)._tset592 = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == Algol60V2ParserT__9 || _la == Algol60V2ParserT__10) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*TermContext)._tset592 = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		localctx.(*TermContext).operators = append(localctx.(*TermContext).operators, localctx.(*TermContext)._tset592)
		{
			p.SetState(308)

			var _x = p.Factor()

			localctx.(*TermContext)._factor = _x
		}
		localctx.(*TermContext).operands = append(localctx.(*TermContext).operands, localctx.(*TermContext)._factor)

		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *FactorContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserQUOTED_STRING, 0)
}

func (s *FactorContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *FactorContext) Function_designator() IFunction_designatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_designatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_designatorContext)
}

func (s *FactorContext) Subscripted_variable() ISubscripted_variableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubscripted_variableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubscripted_variableContext)
}

func (s *FactorContext) Expression() IExpressionContext {
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

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitFactor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Factor() (localctx IFactorContext) {
	this := p
	_ = this

	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, Algol60V2ParserRULE_factor)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(314)
			p.Number()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(315)
			p.Match(Algol60V2ParserQUOTED_STRING)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(316)
			p.Identifier()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(317)
			p.Function_designator()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(318)
			p.Subscripted_variable()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(319)
			p.Match(Algol60V2ParserT__4)
		}
		{
			p.SetState(320)
			p.Expression()
		}
		{
			p.SetState(321)
			p.Match(Algol60V2ParserT__5)
		}

	}

	return localctx
}

// IFunction_designatorContext is an interface to support dynamic dispatch.
type IFunction_designatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetProc_name returns the proc_name rule contexts.
	GetProc_name() IIdentifierContext

	// GetArgs returns the args rule contexts.
	GetArgs() IActual_parameter_listContext

	// SetProc_name sets the proc_name rule contexts.
	SetProc_name(IIdentifierContext)

	// SetArgs sets the args rule contexts.
	SetArgs(IActual_parameter_listContext)

	// IsFunction_designatorContext differentiates from other interfaces.
	IsFunction_designatorContext()
}

type Function_designatorContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	proc_name IIdentifierContext
	args      IActual_parameter_listContext
}

func NewEmptyFunction_designatorContext() *Function_designatorContext {
	var p = new(Function_designatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_function_designator
	return p
}

func (*Function_designatorContext) IsFunction_designatorContext() {}

func NewFunction_designatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_designatorContext {
	var p = new(Function_designatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_function_designator

	return p
}

func (s *Function_designatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_designatorContext) GetProc_name() IIdentifierContext { return s.proc_name }

func (s *Function_designatorContext) GetArgs() IActual_parameter_listContext { return s.args }

func (s *Function_designatorContext) SetProc_name(v IIdentifierContext) { s.proc_name = v }

func (s *Function_designatorContext) SetArgs(v IActual_parameter_listContext) { s.args = v }

func (s *Function_designatorContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Function_designatorContext) Actual_parameter_list() IActual_parameter_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActual_parameter_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActual_parameter_listContext)
}

func (s *Function_designatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_designatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_designatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitFunction_designator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Function_designator() (localctx IFunction_designatorContext) {
	this := p
	_ = this

	localctx = NewFunction_designatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, Algol60V2ParserRULE_function_designator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(325)

		var _x = p.Identifier()

		localctx.(*Function_designatorContext).proc_name = _x
	}
	{
		p.SetState(326)
		p.Match(Algol60V2ParserT__4)
	}
	{
		p.SetState(327)

		var _x = p.Actual_parameter_list()

		localctx.(*Function_designatorContext).args = _x
	}
	{
		p.SetState(328)
		p.Match(Algol60V2ParserT__5)
	}

	return localctx
}

// IBoolean_expressionContext is an interface to support dynamic dispatch.
type IBoolean_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCond returns the cond rule contexts.
	GetCond() IBoolean_expressionContext

	// GetThenValue returns the thenValue rule contexts.
	GetThenValue() IImplicationContext

	// GetElseValue returns the elseValue rule contexts.
	GetElseValue() IBoolean_expressionContext

	// SetCond sets the cond rule contexts.
	SetCond(IBoolean_expressionContext)

	// SetThenValue sets the thenValue rule contexts.
	SetThenValue(IImplicationContext)

	// SetElseValue sets the elseValue rule contexts.
	SetElseValue(IBoolean_expressionContext)

	// IsBoolean_expressionContext differentiates from other interfaces.
	IsBoolean_expressionContext()
}

type Boolean_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	cond      IBoolean_expressionContext
	thenValue IImplicationContext
	elseValue IBoolean_expressionContext
}

func NewEmptyBoolean_expressionContext() *Boolean_expressionContext {
	var p = new(Boolean_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_boolean_expression
	return p
}

func (*Boolean_expressionContext) IsBoolean_expressionContext() {}

func NewBoolean_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_expressionContext {
	var p = new(Boolean_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_boolean_expression

	return p
}

func (s *Boolean_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_expressionContext) GetCond() IBoolean_expressionContext { return s.cond }

func (s *Boolean_expressionContext) GetThenValue() IImplicationContext { return s.thenValue }

func (s *Boolean_expressionContext) GetElseValue() IBoolean_expressionContext { return s.elseValue }

func (s *Boolean_expressionContext) SetCond(v IBoolean_expressionContext) { s.cond = v }

func (s *Boolean_expressionContext) SetThenValue(v IImplicationContext) { s.thenValue = v }

func (s *Boolean_expressionContext) SetElseValue(v IBoolean_expressionContext) { s.elseValue = v }

func (s *Boolean_expressionContext) Implication() IImplicationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImplicationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImplicationContext)
}

func (s *Boolean_expressionContext) IF() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserIF, 0)
}

func (s *Boolean_expressionContext) THEN() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserTHEN, 0)
}

func (s *Boolean_expressionContext) ELSE() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserELSE, 0)
}

func (s *Boolean_expressionContext) AllBoolean_expression() []IBoolean_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoolean_expressionContext); ok {
			len++
		}
	}

	tst := make([]IBoolean_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoolean_expressionContext); ok {
			tst[i] = t.(IBoolean_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Boolean_expressionContext) Boolean_expression(i int) IBoolean_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_expressionContext); ok {
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

	return t.(IBoolean_expressionContext)
}

func (s *Boolean_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Boolean_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitBoolean_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Boolean_expression() (localctx IBoolean_expressionContext) {
	this := p
	_ = this

	localctx = NewBoolean_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, Algol60V2ParserRULE_boolean_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(338)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Algol60V2ParserT__4, Algol60V2ParserT__7, Algol60V2ParserT__8, Algol60V2ParserT__14, Algol60V2ParserQUOTED_STRING, Algol60V2ParserNUMBER, Algol60V2ParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(330)
			p.Implication()
		}

	case Algol60V2ParserIF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(331)
			p.Match(Algol60V2ParserIF)
		}
		{
			p.SetState(332)

			var _x = p.Boolean_expression()

			localctx.(*Boolean_expressionContext).cond = _x
		}
		{
			p.SetState(333)
			p.Match(Algol60V2ParserTHEN)
		}
		{
			p.SetState(334)

			var _x = p.Implication()

			localctx.(*Boolean_expressionContext).thenValue = _x
		}
		{
			p.SetState(335)
			p.Match(Algol60V2ParserELSE)
		}
		{
			p.SetState(336)

			var _x = p.Boolean_expression()

			localctx.(*Boolean_expressionContext).elseValue = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IImplicationContext is an interface to support dynamic dispatch.
type IImplicationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetS12 returns the s12 token.
	GetS12() antlr.Token

	// SetS12 sets the s12 token.
	SetS12(antlr.Token)

	// GetOperators returns the operators token list.
	GetOperators() []antlr.Token

	// SetOperators sets the operators token list.
	SetOperators([]antlr.Token)

	// Get_boolean_term returns the _boolean_term rule contexts.
	Get_boolean_term() IBoolean_termContext

	// Set_boolean_term sets the _boolean_term rule contexts.
	Set_boolean_term(IBoolean_termContext)

	// GetOperands returns the operands rule context list.
	GetOperands() []IBoolean_termContext

	// SetOperands sets the operands rule context list.
	SetOperands([]IBoolean_termContext)

	// IsImplicationContext differentiates from other interfaces.
	IsImplicationContext()
}

type ImplicationContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	_boolean_term IBoolean_termContext
	operands      []IBoolean_termContext
	s12           antlr.Token
	operators     []antlr.Token
}

func NewEmptyImplicationContext() *ImplicationContext {
	var p = new(ImplicationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_implication
	return p
}

func (*ImplicationContext) IsImplicationContext() {}

func NewImplicationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImplicationContext {
	var p = new(ImplicationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_implication

	return p
}

func (s *ImplicationContext) GetParser() antlr.Parser { return s.parser }

func (s *ImplicationContext) GetS12() antlr.Token { return s.s12 }

func (s *ImplicationContext) SetS12(v antlr.Token) { s.s12 = v }

func (s *ImplicationContext) GetOperators() []antlr.Token { return s.operators }

func (s *ImplicationContext) SetOperators(v []antlr.Token) { s.operators = v }

func (s *ImplicationContext) Get_boolean_term() IBoolean_termContext { return s._boolean_term }

func (s *ImplicationContext) Set_boolean_term(v IBoolean_termContext) { s._boolean_term = v }

func (s *ImplicationContext) GetOperands() []IBoolean_termContext { return s.operands }

func (s *ImplicationContext) SetOperands(v []IBoolean_termContext) { s.operands = v }

func (s *ImplicationContext) AllBoolean_term() []IBoolean_termContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoolean_termContext); ok {
			len++
		}
	}

	tst := make([]IBoolean_termContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoolean_termContext); ok {
			tst[i] = t.(IBoolean_termContext)
			i++
		}
	}

	return tst
}

func (s *ImplicationContext) Boolean_term(i int) IBoolean_termContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_termContext); ok {
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

	return t.(IBoolean_termContext)
}

func (s *ImplicationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImplicationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImplicationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitImplication(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Implication() (localctx IImplicationContext) {
	this := p
	_ = this

	localctx = NewImplicationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, Algol60V2ParserRULE_implication)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(340)

		var _x = p.Boolean_term()

		localctx.(*ImplicationContext)._boolean_term = _x
	}
	localctx.(*ImplicationContext).operands = append(localctx.(*ImplicationContext).operands, localctx.(*ImplicationContext)._boolean_term)
	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Algol60V2ParserT__11 {
		{
			p.SetState(341)

			var _m = p.Match(Algol60V2ParserT__11)

			localctx.(*ImplicationContext).s12 = _m
		}
		localctx.(*ImplicationContext).operators = append(localctx.(*ImplicationContext).operators, localctx.(*ImplicationContext).s12)
		{
			p.SetState(342)

			var _x = p.Boolean_term()

			localctx.(*ImplicationContext)._boolean_term = _x
		}
		localctx.(*ImplicationContext).operands = append(localctx.(*ImplicationContext).operands, localctx.(*ImplicationContext)._boolean_term)

		p.SetState(347)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBoolean_termContext is an interface to support dynamic dispatch.
type IBoolean_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetS13 returns the s13 token.
	GetS13() antlr.Token

	// SetS13 sets the s13 token.
	SetS13(antlr.Token)

	// GetOperators returns the operators token list.
	GetOperators() []antlr.Token

	// SetOperators sets the operators token list.
	SetOperators([]antlr.Token)

	// Get_boolean_factor returns the _boolean_factor rule contexts.
	Get_boolean_factor() IBoolean_factorContext

	// Set_boolean_factor sets the _boolean_factor rule contexts.
	Set_boolean_factor(IBoolean_factorContext)

	// GetOperands returns the operands rule context list.
	GetOperands() []IBoolean_factorContext

	// SetOperands sets the operands rule context list.
	SetOperands([]IBoolean_factorContext)

	// IsBoolean_termContext differentiates from other interfaces.
	IsBoolean_termContext()
}

type Boolean_termContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	_boolean_factor IBoolean_factorContext
	operands        []IBoolean_factorContext
	s13             antlr.Token
	operators       []antlr.Token
}

func NewEmptyBoolean_termContext() *Boolean_termContext {
	var p = new(Boolean_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_boolean_term
	return p
}

func (*Boolean_termContext) IsBoolean_termContext() {}

func NewBoolean_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_termContext {
	var p = new(Boolean_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_boolean_term

	return p
}

func (s *Boolean_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_termContext) GetS13() antlr.Token { return s.s13 }

func (s *Boolean_termContext) SetS13(v antlr.Token) { s.s13 = v }

func (s *Boolean_termContext) GetOperators() []antlr.Token { return s.operators }

func (s *Boolean_termContext) SetOperators(v []antlr.Token) { s.operators = v }

func (s *Boolean_termContext) Get_boolean_factor() IBoolean_factorContext { return s._boolean_factor }

func (s *Boolean_termContext) Set_boolean_factor(v IBoolean_factorContext) { s._boolean_factor = v }

func (s *Boolean_termContext) GetOperands() []IBoolean_factorContext { return s.operands }

func (s *Boolean_termContext) SetOperands(v []IBoolean_factorContext) { s.operands = v }

func (s *Boolean_termContext) AllBoolean_factor() []IBoolean_factorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoolean_factorContext); ok {
			len++
		}
	}

	tst := make([]IBoolean_factorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoolean_factorContext); ok {
			tst[i] = t.(IBoolean_factorContext)
			i++
		}
	}

	return tst
}

func (s *Boolean_termContext) Boolean_factor(i int) IBoolean_factorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_factorContext); ok {
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

	return t.(IBoolean_factorContext)
}

func (s *Boolean_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Boolean_termContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitBoolean_term(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Boolean_term() (localctx IBoolean_termContext) {
	this := p
	_ = this

	localctx = NewBoolean_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, Algol60V2ParserRULE_boolean_term)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(348)

		var _x = p.Boolean_factor()

		localctx.(*Boolean_termContext)._boolean_factor = _x
	}
	localctx.(*Boolean_termContext).operands = append(localctx.(*Boolean_termContext).operands, localctx.(*Boolean_termContext)._boolean_factor)
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Algol60V2ParserT__12 {
		{
			p.SetState(349)

			var _m = p.Match(Algol60V2ParserT__12)

			localctx.(*Boolean_termContext).s13 = _m
		}
		localctx.(*Boolean_termContext).operators = append(localctx.(*Boolean_termContext).operators, localctx.(*Boolean_termContext).s13)
		{
			p.SetState(350)

			var _x = p.Boolean_factor()

			localctx.(*Boolean_termContext)._boolean_factor = _x
		}
		localctx.(*Boolean_termContext).operands = append(localctx.(*Boolean_termContext).operands, localctx.(*Boolean_termContext)._boolean_factor)

		p.SetState(355)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBoolean_factorContext is an interface to support dynamic dispatch.
type IBoolean_factorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetS14 returns the s14 token.
	GetS14() antlr.Token

	// SetS14 sets the s14 token.
	SetS14(antlr.Token)

	// GetOperators returns the operators token list.
	GetOperators() []antlr.Token

	// SetOperators sets the operators token list.
	SetOperators([]antlr.Token)

	// Get_boolean_secondary returns the _boolean_secondary rule contexts.
	Get_boolean_secondary() IBoolean_secondaryContext

	// Set_boolean_secondary sets the _boolean_secondary rule contexts.
	Set_boolean_secondary(IBoolean_secondaryContext)

	// GetOperands returns the operands rule context list.
	GetOperands() []IBoolean_secondaryContext

	// SetOperands sets the operands rule context list.
	SetOperands([]IBoolean_secondaryContext)

	// IsBoolean_factorContext differentiates from other interfaces.
	IsBoolean_factorContext()
}

type Boolean_factorContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	_boolean_secondary IBoolean_secondaryContext
	operands           []IBoolean_secondaryContext
	s14                antlr.Token
	operators          []antlr.Token
}

func NewEmptyBoolean_factorContext() *Boolean_factorContext {
	var p = new(Boolean_factorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_boolean_factor
	return p
}

func (*Boolean_factorContext) IsBoolean_factorContext() {}

func NewBoolean_factorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_factorContext {
	var p = new(Boolean_factorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_boolean_factor

	return p
}

func (s *Boolean_factorContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_factorContext) GetS14() antlr.Token { return s.s14 }

func (s *Boolean_factorContext) SetS14(v antlr.Token) { s.s14 = v }

func (s *Boolean_factorContext) GetOperators() []antlr.Token { return s.operators }

func (s *Boolean_factorContext) SetOperators(v []antlr.Token) { s.operators = v }

func (s *Boolean_factorContext) Get_boolean_secondary() IBoolean_secondaryContext {
	return s._boolean_secondary
}

func (s *Boolean_factorContext) Set_boolean_secondary(v IBoolean_secondaryContext) {
	s._boolean_secondary = v
}

func (s *Boolean_factorContext) GetOperands() []IBoolean_secondaryContext { return s.operands }

func (s *Boolean_factorContext) SetOperands(v []IBoolean_secondaryContext) { s.operands = v }

func (s *Boolean_factorContext) AllBoolean_secondary() []IBoolean_secondaryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoolean_secondaryContext); ok {
			len++
		}
	}

	tst := make([]IBoolean_secondaryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoolean_secondaryContext); ok {
			tst[i] = t.(IBoolean_secondaryContext)
			i++
		}
	}

	return tst
}

func (s *Boolean_factorContext) Boolean_secondary(i int) IBoolean_secondaryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_secondaryContext); ok {
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

	return t.(IBoolean_secondaryContext)
}

func (s *Boolean_factorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_factorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Boolean_factorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitBoolean_factor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Boolean_factor() (localctx IBoolean_factorContext) {
	this := p
	_ = this

	localctx = NewBoolean_factorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, Algol60V2ParserRULE_boolean_factor)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(356)

		var _x = p.Boolean_secondary()

		localctx.(*Boolean_factorContext)._boolean_secondary = _x
	}
	localctx.(*Boolean_factorContext).operands = append(localctx.(*Boolean_factorContext).operands, localctx.(*Boolean_factorContext)._boolean_secondary)
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Algol60V2ParserT__13 {
		{
			p.SetState(357)

			var _m = p.Match(Algol60V2ParserT__13)

			localctx.(*Boolean_factorContext).s14 = _m
		}
		localctx.(*Boolean_factorContext).operators = append(localctx.(*Boolean_factorContext).operators, localctx.(*Boolean_factorContext).s14)
		{
			p.SetState(358)

			var _x = p.Boolean_secondary()

			localctx.(*Boolean_factorContext)._boolean_secondary = _x
		}
		localctx.(*Boolean_factorContext).operands = append(localctx.(*Boolean_factorContext).operands, localctx.(*Boolean_factorContext)._boolean_secondary)

		p.SetState(363)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBoolean_secondaryContext is an interface to support dynamic dispatch.
type IBoolean_secondaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOperator returns the operator token.
	GetOperator() antlr.Token

	// SetOperator sets the operator token.
	SetOperator(antlr.Token)

	// GetOperand returns the operand rule contexts.
	GetOperand() IBoolean_primaryContext

	// SetOperand sets the operand rule contexts.
	SetOperand(IBoolean_primaryContext)

	// IsBoolean_secondaryContext differentiates from other interfaces.
	IsBoolean_secondaryContext()
}

type Boolean_secondaryContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	operator antlr.Token
	operand  IBoolean_primaryContext
}

func NewEmptyBoolean_secondaryContext() *Boolean_secondaryContext {
	var p = new(Boolean_secondaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_boolean_secondary
	return p
}

func (*Boolean_secondaryContext) IsBoolean_secondaryContext() {}

func NewBoolean_secondaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_secondaryContext {
	var p = new(Boolean_secondaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_boolean_secondary

	return p
}

func (s *Boolean_secondaryContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_secondaryContext) GetOperator() antlr.Token { return s.operator }

func (s *Boolean_secondaryContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *Boolean_secondaryContext) GetOperand() IBoolean_primaryContext { return s.operand }

func (s *Boolean_secondaryContext) SetOperand(v IBoolean_primaryContext) { s.operand = v }

func (s *Boolean_secondaryContext) Boolean_primary() IBoolean_primaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_primaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolean_primaryContext)
}

func (s *Boolean_secondaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_secondaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Boolean_secondaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitBoolean_secondary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Boolean_secondary() (localctx IBoolean_secondaryContext) {
	this := p
	_ = this

	localctx = NewBoolean_secondaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, Algol60V2ParserRULE_boolean_secondary)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Algol60V2ParserT__14 {
		{
			p.SetState(364)

			var _m = p.Match(Algol60V2ParserT__14)

			localctx.(*Boolean_secondaryContext).operator = _m
		}

	}
	{
		p.SetState(367)

		var _x = p.Boolean_primary()

		localctx.(*Boolean_secondaryContext).operand = _x
	}

	return localctx
}

// IBoolean_primaryContext is an interface to support dynamic dispatch.
type IBoolean_primaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOperator returns the operator token.
	GetOperator() antlr.Token

	// SetOperator sets the operator token.
	SetOperator(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() ISimple_expressionContext

	// GetRight returns the right rule contexts.
	GetRight() ISimple_expressionContext

	// SetLeft sets the left rule contexts.
	SetLeft(ISimple_expressionContext)

	// SetRight sets the right rule contexts.
	SetRight(ISimple_expressionContext)

	// IsBoolean_primaryContext differentiates from other interfaces.
	IsBoolean_primaryContext()
}

type Boolean_primaryContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	left     ISimple_expressionContext
	operator antlr.Token
	right    ISimple_expressionContext
}

func NewEmptyBoolean_primaryContext() *Boolean_primaryContext {
	var p = new(Boolean_primaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_boolean_primary
	return p
}

func (*Boolean_primaryContext) IsBoolean_primaryContext() {}

func NewBoolean_primaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_primaryContext {
	var p = new(Boolean_primaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_boolean_primary

	return p
}

func (s *Boolean_primaryContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_primaryContext) GetOperator() antlr.Token { return s.operator }

func (s *Boolean_primaryContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *Boolean_primaryContext) GetLeft() ISimple_expressionContext { return s.left }

func (s *Boolean_primaryContext) GetRight() ISimple_expressionContext { return s.right }

func (s *Boolean_primaryContext) SetLeft(v ISimple_expressionContext) { s.left = v }

func (s *Boolean_primaryContext) SetRight(v ISimple_expressionContext) { s.right = v }

func (s *Boolean_primaryContext) AllSimple_expression() []ISimple_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimple_expressionContext); ok {
			len++
		}
	}

	tst := make([]ISimple_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimple_expressionContext); ok {
			tst[i] = t.(ISimple_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Boolean_primaryContext) Simple_expression(i int) ISimple_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_expressionContext); ok {
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

	return t.(ISimple_expressionContext)
}

func (s *Boolean_primaryContext) Boolean_expression() IBoolean_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolean_expressionContext)
}

func (s *Boolean_primaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_primaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Boolean_primaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitBoolean_primary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Boolean_primary() (localctx IBoolean_primaryContext) {
	this := p
	_ = this

	localctx = NewBoolean_primaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, Algol60V2ParserRULE_boolean_primary)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(377)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(369)

			var _x = p.Simple_expression()

			localctx.(*Boolean_primaryContext).left = _x
		}
		{
			p.SetState(370)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Boolean_primaryContext).operator = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Algol60V2ParserT__15)|(1<<Algol60V2ParserT__16)|(1<<Algol60V2ParserT__17)|(1<<Algol60V2ParserT__18)|(1<<Algol60V2ParserT__19)|(1<<Algol60V2ParserT__20))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Boolean_primaryContext).operator = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(371)

			var _x = p.Simple_expression()

			localctx.(*Boolean_primaryContext).right = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(373)
			p.Match(Algol60V2ParserT__4)
		}
		{
			p.SetState(374)
			p.Boolean_expression()
		}
		{
			p.SetState(375)
			p.Match(Algol60V2ParserT__5)
		}

	}

	return localctx
}

// ISubscripted_variableContext is an interface to support dynamic dispatch.
type ISubscripted_variableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetArrayVar returns the arrayVar rule contexts.
	GetArrayVar() IIdentifierContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// SetArrayVar sets the arrayVar rule contexts.
	SetArrayVar(IIdentifierContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetIdxExprs returns the idxExprs rule context list.
	GetIdxExprs() []IExpressionContext

	// SetIdxExprs sets the idxExprs rule context list.
	SetIdxExprs([]IExpressionContext)

	// IsSubscripted_variableContext differentiates from other interfaces.
	IsSubscripted_variableContext()
}

type Subscripted_variableContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	arrayVar    IIdentifierContext
	_expression IExpressionContext
	idxExprs    []IExpressionContext
}

func NewEmptySubscripted_variableContext() *Subscripted_variableContext {
	var p = new(Subscripted_variableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_subscripted_variable
	return p
}

func (*Subscripted_variableContext) IsSubscripted_variableContext() {}

func NewSubscripted_variableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Subscripted_variableContext {
	var p = new(Subscripted_variableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_subscripted_variable

	return p
}

func (s *Subscripted_variableContext) GetParser() antlr.Parser { return s.parser }

func (s *Subscripted_variableContext) GetArrayVar() IIdentifierContext { return s.arrayVar }

func (s *Subscripted_variableContext) Get_expression() IExpressionContext { return s._expression }

func (s *Subscripted_variableContext) SetArrayVar(v IIdentifierContext) { s.arrayVar = v }

func (s *Subscripted_variableContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Subscripted_variableContext) GetIdxExprs() []IExpressionContext { return s.idxExprs }

func (s *Subscripted_variableContext) SetIdxExprs(v []IExpressionContext) { s.idxExprs = v }

func (s *Subscripted_variableContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Subscripted_variableContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *Subscripted_variableContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *Subscripted_variableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Subscripted_variableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Subscripted_variableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitSubscripted_variable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Subscripted_variable() (localctx ISubscripted_variableContext) {
	this := p
	_ = this

	localctx = NewSubscripted_variableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, Algol60V2ParserRULE_subscripted_variable)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(379)

		var _x = p.Identifier()

		localctx.(*Subscripted_variableContext).arrayVar = _x
	}
	{
		p.SetState(380)
		p.Match(Algol60V2ParserT__2)
	}
	{
		p.SetState(381)

		var _x = p.Expression()

		localctx.(*Subscripted_variableContext)._expression = _x
	}
	localctx.(*Subscripted_variableContext).idxExprs = append(localctx.(*Subscripted_variableContext).idxExprs, localctx.(*Subscripted_variableContext)._expression)
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Algol60V2ParserT__1 {
		{
			p.SetState(382)
			p.Match(Algol60V2ParserT__1)
		}
		{
			p.SetState(383)

			var _x = p.Expression()

			localctx.(*Subscripted_variableContext)._expression = _x
		}
		localctx.(*Subscripted_variableContext).idxExprs = append(localctx.(*Subscripted_variableContext).idxExprs, localctx.(*Subscripted_variableContext)._expression)

		p.SetState(388)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(389)
		p.Match(Algol60V2ParserT__3)
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserIDENTIFIER, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Identifier() (localctx IIdentifierContext) {
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, Algol60V2ParserRULE_identifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(391)
		p.Match(Algol60V2ParserIDENTIFIER)
	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value attribute.
	GetValue() int64

	// SetValue sets the value attribute.
	SetValue(int64)

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  int64
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Algol60V2ParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Algol60V2ParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) GetValue() int64 { return s.value }

func (s *NumberContext) SetValue(v int64) { s.value = v }

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(Algol60V2ParserNUMBER, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Algol60V2Visitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Algol60V2Parser) Number() (localctx INumberContext) {
	this := p
	_ = this

	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, Algol60V2ParserRULE_number)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(393)
		p.Match(Algol60V2ParserNUMBER)
	}

	val, err := strconv.ParseInt(p.GetTokenStream().GetTextFromTokens(localctx.GetStart(), p.GetTokenStream().LT(-1)), 10, 64)
	if err != nil {
		panic(err)
	}
	localctx.(*NumberContext).value = val

	return localctx
}
