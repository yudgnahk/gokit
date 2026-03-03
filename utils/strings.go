package utils

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/yudgnahk/gokit/constants"
)

// support:
// splitting words: Hello World
// snake: hello_world
// camel words: HelloWorld

// Snake ...
func Snake(s string) string {
	var result string
	if strings.Contains(s, constants.Space) {
		words := strings.Split(s, constants.Space)
		for i := range words {
			words[i] = strings.ToLower(words[i])
		}

		result = strings.Join(words, constants.Underscore)
		return result
	}

	if strings.Contains(s, constants.Underscore) {
		words := strings.Split(s, constants.Underscore)
		for i := range words {
			words[i] = strings.ToLower(words[i])
		}

		result = strings.Join(words, constants.Underscore)
		return result
	}

	re := regexp.MustCompile(`[A-Za-z][^A-Z]*`)
	words := re.FindAllString(s, -1)
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}

	result = strings.Join(words, constants.Underscore)
	return result
}

// Camel ...
func Camel(s string, lowerFirst bool) string {
	var result string
	if strings.Contains(s, constants.Space) {
		words := strings.Split(s, constants.Space)
		for i := range words {
			words[i] = strings.Title(strings.ToLower(words[i]))
		}

		result = strings.Join(words, constants.EmptyString)
		if lowerFirst {
			result = LcFirst(result)
		}
		return result
	}

	if strings.Contains(s, constants.Underscore) {
		words := strings.Split(s, constants.Underscore)
		for i := range words {
			words[i] = strings.Title(strings.ToLower(words[i]))
		}

		result = strings.Join(words, constants.EmptyString)
		if lowerFirst {
			result = LcFirst(result)
		}
		return result
	}

	re := regexp.MustCompile(`[A-Za-z][^A-Z]*`)
	words := re.FindAllString(s, -1)
	for i := range words {
		words[i] = strings.Title(strings.ToLower(words[i]))
	}

	result = strings.Join(words, constants.EmptyString)
	if lowerFirst {
		result = LcFirst(result)
	}
	return result
}

// UcFirst ...
func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// LcFirst ...
func LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func GetLastPart(s string) string {
	words := strings.Split(s, constants.Slash)
	return words[len(words)-1]
}
