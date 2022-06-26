package main

import (
	"checker/internal/algol60_printer"
	"checker/internal/utils"
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const testCasesPath = "../compiler/src/test/resources/"

func main() {
	utils.CheckerRand = *rand.New(rand.NewSource(time.Now().UnixNano()))

	var testCases []string
	err := filepath.WalkDir(testCasesPath,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() && strings.HasSuffix(path, ".a60") {
				testCases = append(testCases, path)
			}
			return nil
		})
	if err != nil {
		log.Panicf("Err reading testcases %s\n", err)
	}

	for _, testCasePath := range testCases {
		filename := filepath.Base(testCasePath)
		code, err := os.ReadFile(testCasePath)
		if err != nil {
			log.Panicf("Err failed to read testcase[%s]: %s\n", filename, err)
		}

		expectedErrorLine := -1
		if idx := strings.Index(filename, "-L"); idx > 0 {
			expectedErrorLineStr := filename[idx+2 : len(filename)-len(".a60")]
			expectedErrorLine, err = strconv.Atoi(expectedErrorLineStr)
			if err != nil {
				log.Panicf("Err parsing line in testcase filename[%s]: %s\n", filename, err)
			}
		}

		randCode, newErrorLine := algol60_printer.RandPrint(string(code), expectedErrorLine, false)
		if false {
			fmt.Printf("NewErrLine: %d\n", newErrorLine)
		}

		// check double prettyPrint equal
		pretty := algol60_printer.PrettyPrint(randCode)
		if pretty != algol60_printer.PrettyPrint(pretty) {
			fmt.Printf("File: %s\n", filename)
			os.WriteFile("./foo.a60", []byte(randCode), os.FileMode(0644))
			log.Panicf("rand fail %s\n", pretty)
		}
	}
}
