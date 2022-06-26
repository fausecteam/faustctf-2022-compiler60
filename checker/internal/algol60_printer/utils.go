package algol60_printer

import (
	"checker/internal/algol60_parser"
	"checker/internal/utils"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
)

const maxSourceSize = 16 * 1024

type Algol60Printer interface {
	algol60_parser.Algol60V2Visitor

	GetOutput() string
}

// like ConsoleErrorListener but use log instead of fmt
type LogErrorListener struct {
	// ugly hack of embedding a struct, so we only need to "override" the function which doesn't do nothing
	// don't ask antlr uses this everywhere
	*antlr.DefaultErrorListener
}

func (l *LogErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	log.Printf("runPrinter: syntax error line %d:%d %s\n", line, column, msg)
}

func runPrinter(code string, printer Algol60Printer) string {
	input := antlr.NewInputStream(code)
	lexer := algol60_parser.NewAlgol60V2Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := algol60_parser.NewAlgol60V2Parser(stream)
	//p.SetErrorHandler(antlr.NewBailErrorStrategy())
	p.AddErrorListener(&LogErrorListener{})
	// parse program
	program := p.Program()

	printer.Visit(program)

	return printer.GetOutput()
}

func RandPrint(code string, errLine int, deterministic bool) (resultCode string, newErrorLine int) {
	randPrinter := RandomPrintVisitor{
		Rand:                     &utils.CheckerRand,
		ErrorLine:                errLine,
		DeterministicCompilation: deterministic,
	}
	result := runPrinter(code, &randPrinter)

	return result, randPrinter.GetNewErrorLine()
}

func PrettyPrint(code string) string {

	prettyPrinter := PrettyPrintVisitor{}

	return runPrinter(code, &prettyPrinter)
}

func RewriteCode(code string, errLine int) (resultCode string, newErrorLine int) {
	switch utils.CheckerRand.Int31n(10) {
	case 0:
		return code, errLine
	case 1:
		if errLine == -1 {
			// pretty print does not keep ErrorLine - too lazy to implement
			return PrettyPrint(code), -1
		}
	}
	resultCode, newErrorLine = RandPrint(code, errLine, false)
	if len(resultCode) > maxSourceSize {
		return code, errLine
	}
	return
}

func RewriteCodeDeterministic(code string) string {
	switch utils.CheckerRand.Int31n(10) {
	case 0:
		return code
	case 1:
		return PrettyPrint(code)
	}
	resultCode, _ := RandPrint(code, -1, true)
	if len(resultCode) > maxSourceSize {
		return code
	}
	return resultCode
}
