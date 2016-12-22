package main

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func HandleYAMLMetaData(datum []byte) (interface{}, error) {
	m := map[string]interface{}{}
	err := yaml.Unmarshal(datum, &m)
	return m, err
}

func main() {
	input, err := ioutil.ReadFile("test.md")
	check(err)

	extensions := 0 |
		blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_HEADER_IDS |
		blackfriday.EXTENSION_BACKSLASH_LINE_BREAK |
		blackfriday.EXTENSION_DEFINITION_LISTS |
		blackfriday.EXTENSION_FOOTNOTES

	commonHtmlFlags := 0 |
		blackfriday.HTML_USE_XHTML |
		blackfriday.HTML_USE_SMARTYPANTS |
		blackfriday.HTML_SMARTYPANTS_FRACTIONS |
		blackfriday.HTML_SMARTYPANTS_DASHES |
		blackfriday.HTML_SMARTYPANTS_LATEX_DASHES

	var renderer blackfriday.Renderer
	renderer = blackfriday.HtmlRenderer(commonHtmlFlags, "", "")

	yaml, err := HandleYAMLMetaData(input)
	if err != nil {
		panic(err)
	}

	fmt.Print(yaml)
	// Need to remove yaml from file

	unsafe := blackfriday.Markdown(input, renderer, extensions)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	fmt.Print(string(html))
}
