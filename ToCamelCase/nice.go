package main

import (
	"regexp"
	"strings"
)

func ToCamelCaseNice(s string) string {
	return regexp.MustCompile("[-_](.)").ReplaceAllStringFunc(s, func(w string) string {
		return strings.ToUpper(w[1:])
	})
}
