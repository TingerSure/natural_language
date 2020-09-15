package nl_string

import (
	"strings"
	"unicode/utf8"
)

func Len(info string) int {
	return utf8.RuneCountInString(info)
}

func SubString(info string, from, to int) string {
	return string([]rune(info)[from:to])
}

func SubStringFrom(info string, from int) string {
	return string([]rune(info)[from:])
}

func SubStringTo(info string, to int) string {
	return string([]rune(info)[:to])
}

func Index(info string, sub string) int {
	var index = strings.Index(info, sub)
	if index <= 0 {
		return index
	}
	return Len(info[:index])
}
