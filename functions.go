package main

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

// Check function for handling errors
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// HandleYAMLMetaData function for parse yaml data
func HandleYAMLMetaData(datum []byte) (interface{}, error) {
	m := map[string]interface{}{}
	err := yaml.Unmarshal(datum, &m)
	return m, err
}

// CheckFrontMatter function for handling FrontMatter in files
func CheckFrontMatter(datum []byte) {
	if datum[0] == '-' {
		fmt.Print("We have YAML (maybe)")
	}
}
