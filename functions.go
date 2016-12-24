package main

import (
	"fmt"
	"regexp"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// Check function for handling errors
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// HandleYAMLMetaData function for parse yaml data
func HandleYAMLMetaData(datum []byte) (map[string]interface{}, error) {
	m := map[string]interface{}{}
	err := yaml.Unmarshal(datum, &m)
	return m, err
}

// CheckFrontMatter function for handling FrontMatter in files
func CheckFrontMatter(datum []byte) {
	re := regexp.MustCompile("(?s)^---\n.*\n---\n")
	if re.Match(datum) {
		fmt.Println("We have YAML")
	} else {
		fmt.Println("We dont have YAML fronmatter")
	}
}

// SplitYaml function for spliting yaml formater and body
func SplitYaml(datum []byte) ([]byte, []byte) {
	re := regexp.MustCompile("(?s)^---\n.*\n---\n")
	yaml := re.FindString(string(datum))
	body := strings.TrimPrefix(string(datum), yaml)
	return []byte(yaml), []byte(body)
}
