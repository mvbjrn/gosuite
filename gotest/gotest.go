// Copyright 2013 r00ky. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/scanner"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

var (
	functions = flag.Int("functions", 1, "insert a specific number of test functions")
	exitCode  = 0
)

func report(err error) {
	scanner.PrintError(os.Stderr, err)
	exitCode = 2
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: gotest [flags] [path to file]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {

	// call gotestMain in a separate function
	// so that it can use defer and have them
	// run before the exit.
	gotestMain()
	os.Exit(exitCode)

}

func gotestMain() {

	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		usage()
	}

	dir, file := filepath.Split(flag.Arg(0))
	var err error

	if dir == "" || dir == "./" || dir == "../" {
		dir, err = os.Getwd()
		if err != nil {
			report(err)
		}
	}

	regex := regexp.MustCompile("(?P<filename>.*)\\.go")
	match := regex.FindStringSubmatch(file)

	if len(match) == 0 {
		fmt.Fprintf(os.Stderr, "file is not a go file\n")
		os.Exit(2)
	}

	testfilePath := fmt.Sprintf("%s/%s_test.go", dir, match[1])

	if _, err := os.Stat(testfilePath); err == nil {
		fmt.Fprintf(os.Stderr, "test file already exists\n")
		os.Exit(2)
	}

	var content bytes.Buffer

	content.WriteString("// licence goes here\n\n")
	content.WriteString("package ")
	// TODO: get package name from file
	content.WriteString(filepath.Base(dir))
	content.WriteString("\n\n")
	content.WriteString("import (\n\t\"testing\"\n)\n\n")
	content.WriteString("// function pattern: func TestFunctionname(t *testing.T) {}\n\n")

	for i := 0; i < *functions; i++ {

		content.WriteString("// TODO: documentation\n")
		content.WriteString("func Test")
		content.WriteString(strconv.Itoa(i + 1))
		content.WriteString("(t *testing.T) {\n\t// TODO:\n}\n\n")
	}

	ioutil.WriteFile(testfilePath, content.Bytes(), 0666)

}
