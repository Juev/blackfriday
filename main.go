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

	CheckFrontMatter(input)

	var renderer blackfriday.Renderer
	renderer = blackfriday.HtmlRenderer(commonHTMLFlags, "", "")

	yaml, err := HandleYAMLMetaData(input)
	Check(err)

	fmt.Print(yaml)
	// Need to remove yaml from file

	unsafe := blackfriday.Markdown(input, renderer, extensions)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	fmt.Print(string(html))
}
