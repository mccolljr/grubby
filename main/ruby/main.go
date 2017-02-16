package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/grubby/grubby/interpreter/vm"
	"github.com/grubby/grubby/parser"
)

var verboseFlag = flag.Bool("verbose", false, "enables verbose mode")
var versionFlag = flag.Bool("version", false, "print version and exit")

func init() {
	flag.BoolVar(versionFlag, "v", false, "print version and exit")
	flag.BoolVar(verboseFlag, "V", false, "enables verbose mode")
}

func main() {
	flag.Parse()

	// exit early if we've been asked for the version
	if *versionFlag {
		fmt.Println("grubby")
		return
	}

	var err error
	var file *os.File
	var fname string

	if len(flag.Args()) == 0 {
		file = os.Stdin
		fname = "STDIN"
	} else {
		fname = flag.Args()[0]
		file, err = os.Open(fname)
		if err != nil {
			fmt.Printf("can't open file %s, aborting\n", fname)
			os.Exit(1)
		}
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("can't read file %s, aborting\n", fname)
		os.Exit(1)
	}

	// TODO is this a good idea? Should we use standard ruby environment variables?
	home := os.Getenv("HOME")
	grubbyHome := filepath.Join(home, ".grubby")

	rubyVM := vm.NewVM(grubbyHome, fname)
	defer rubyVM.Exit()

	_, err = rubyVM.Run(string(bytes))

	switch err.(type) {
	case *vm.ParseError:
		offendingFilename := err.(*vm.ParseError).Filename
		println(fmt.Sprintf("Error parsing ruby script %s", offendingFilename))
		println("last 20 statements from the parser:")
		println("")

		debugStatements := []string{}
		for _, d := range parser.DebugStatements {
			debugStatements = append(debugStatements, d)
		}

		threshold := 61
		debugCount := len(debugStatements)
		if debugCount <= threshold {
			for _, stmt := range debugStatements {
				fmt.Printf("\t%s\n", stmt)
			}
		} else {
			for _, stmt := range debugStatements[debugCount-threshold:] {
				fmt.Printf("\t%s\n", stmt)
			}
		}

		os.Exit(1)
	case nil:
	case error:
		panic(err.Error())
	default:
		panic(fmt.Sprintf("%#v", err))
	}
}
