package main

import (
	"flag"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"io/ioutil"
)

func main() {
	var filename = flag.String("filename", "test.md", "filename for parsing")
	flag.Parse()

	input, err := ioutil.ReadFile(*filename)
	Check(err)

	yaml, markdown := SplitYaml(input)

	var renderer blackfriday.Renderer
	renderer = blackfriday.HtmlRenderer(commonHTMLFlags, "", "")

	yamlResult, err := HandleYAMLMetaData(yaml)
	Check(err)

	for k, v := range yamlResult {
		fmt.Println("Key: ", k, " Value: ", v)
	}
	fmt.Println()

	unsafe := blackfriday.Markdown(markdown, renderer, extensions)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	fmt.Print(string(html))
}
