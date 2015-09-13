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
)

var (
	mainFlag    = flag.Bool("with-main", false, "insert main function")
	constFlag   = flag.Bool("without-const", false, "insert const")
	varFlag     = flag.Bool("without-var", false, "insert var")
	packageFlag = flag.String("package", "", "insert alternate package")
	exitCode    = 0
)

func report(err error) {
	scanner.PrintError(os.Stderr, err)
	exitCode = 2
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: gofile [flags] filepath\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {

	// call gofileMain in a separate function
	// so that it can use defer and have them
	// run before the exit.
	gofileMain()
	os.Exit(exitCode)

}

func gofileMain() {

	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 || flag.NArg() > 1 {
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

	filePath := fmt.Sprintf("%s/%s", dir, file)

	if _, err := os.Stat(filePath); err == nil {
		fmt.Fprintf(os.Stderr, "file already exists\n")
		os.Exit(2)
	}

	packageName := filepath.Base(dir)

	if *packageFlag != "" {
		packageName = *packageFlag
	}

	var content bytes.Buffer

	content.WriteString("// licence goes here\n\n")
	content.WriteString("package ")
	content.WriteString(packageName)
	content.WriteString("\n\n")

	if !*constFlag {
		content.WriteString("// const\n")
		content.WriteString("const (\n\t\n)\n\n")
	}

	if !*varFlag {
		content.WriteString("// var\n")
		content.WriteString("var (\n\t\n)\n\n")
	}

	content.WriteString("// structs and its functions\n\n")
	content.WriteString("// functions\n\n")

	if *mainFlag {
		content.WriteString("// main function\n")
		content.WriteString("func main() {\n\t\n}\n")
	}

	ioutil.WriteFile(filePath, content.Bytes(), 0666)

}
