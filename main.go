package main

import (
	"flag"
	"fmt"
)

var filename string

func init() {
	const (
		defaultFilename = "test.md"
		usage           = "filename for parsing"
	)
	flag.StringVar(&filename, "filename", defaultFilename, usage)
	flag.StringVar(&filename, "f", defaultFilename, usage+" (shorthand)")
}

func main() {
	flag.Parse()

	yamlResult, html := ParseFile(filename)

	for k, v := range yamlResult {
		fmt.Println("Key: ", k, " Value: ", v)
	}
	fmt.Println()

	fmt.Print(string(html))
}
