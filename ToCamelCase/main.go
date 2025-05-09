package main

import "strings"

func ToCamelCase(s string) string {
	b := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if s[i] != '_' && s[i] != '-' {
			b.WriteByte(s[i])
			continue
		}
		if i+1 < len(s) {
			b.WriteByte(strings.ToUpper(string(s[i+1]))[0])
			i++
		}
	}

	return b.String()
}
