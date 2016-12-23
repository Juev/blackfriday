package main

import (
	"fmt"
	"io/ioutil"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func main() {
	input, err := ioutil.ReadFile("test.md")
	Check(err)

	yaml, markdown := SplitYaml(input)

	var renderer blackfriday.Renderer
	renderer = blackfriday.HtmlRenderer(commonHTMLFlags, "", "")

	yamlResult, err := HandleYAMLMetaData(yaml)
	Check(err)

	fmt.Println("type: " + yamlResult["type"])
	fmt.Println("title: " + yamlResult["title"])
	fmt.Println()

	unsafe := blackfriday.Markdown(markdown, renderer, extensions)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	fmt.Print(string(html))
}
