[![Build Status](https://travis-ci.org/mvbjrn/gosuite.svg?branch=master)](https://travis-ci.org/mvbjrn/gosuite)
[![Coverage](http://gocover.io/_badge/github.com/mvbjrn/gosuite)](http://gocover.io/github.com/mvbjrn/gosuite)
[![GoDoc](https://godoc.org/github.com/mvbjrn/gosuite?status.svg)](https://godoc.org/github.com/mvbjrn/gosuite)

# gosuite



## gofile

This tool allows you to create a go file.

Simply execute gofile in your command-line and add flags and the desired file name.


	$ gofile -with-main main.go

produces a main.go

```go
// licence goes here

package name

// const
const (

)

// var
var (

)

// structs and its functions

// functions

// main function
func main() {

}

```

## gotest

This tool allows you to create a test file for your go file, following the go convention.

Simply execute gotest in your command-line and add the file to be tested.

## gopackdoc

This tool allows you to create a simple doc.go in your package.

Simply execute gopackdoc in your command-line and add one or more paths.
