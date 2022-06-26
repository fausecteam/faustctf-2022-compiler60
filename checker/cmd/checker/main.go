package main

import (
	"checker/internal/checker"
	"github.com/fausecteam/ctf-gameserver/go/checkerlib"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if os.Getenv("CTF_CHECKERSCRIPT") != "" {
		// Checker Runner has working dir '/' but we need to access local files
		self, _ := os.Executable()
		selfDir := filepath.Dir(self)
		_ = os.Chdir(selfDir)
	}
	pwd, _ := os.Getwd()
	log.Printf("Running checker in '%s' with args: ip=%s team=%s tick=%s\n", pwd, os.Args[1], os.Args[2], os.Args[3])

	defer func() {
		if r := recover(); r != nil {
			// make sure to log panics
			log.Panicf("checker paniced: %v\n", r)
		}
	}()

	checkerlib.RunCheck(&checker.Compiler60Checker{})
}
