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
	exitCode = 0
)

func report(err error) {
	scanner.PrintError(os.Stderr, err)
	exitCode = 2
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: gopackdoc [flags] [path ...]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {

	// call gopackdocMain in a separate function
	// so that it can use defer and have them
	// run before the exit.
	gopackdocMain()
	os.Exit(exitCode)

}

func gopackdocMain() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		usage()
	}

	for i := 0; i < flag.NArg(); i++ {
		createDoc(flag.Arg(i))
	}
}

func createDoc(path string) {

	err := os.Chdir(path)
	if err != nil {
		report(err)
	}

	wd, err := os.Getwd()
	if err != nil {
		report(err)
	}

	packageName := filepath.Base(wd)

	var content bytes.Buffer

	content.WriteString("// licence goes here\n\n")
	content.WriteString("// ")
	content.WriteString(packageName)
	content.WriteString("...\n")
	content.WriteString("package ")
	content.WriteString(packageName)
	content.WriteString("\n")

	ioutil.WriteFile("./doc.go", content.Bytes(), 0666)

}
