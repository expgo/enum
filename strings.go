package main

import "strings"

func capitalize(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func IndefiniteArticle(word string) string {
	lc := strings.ToLower(word[:1])
	switch lc {
	case "a", "e", "i", "o", "u":
		return "an " + word
	default:
		return "a " + word
	}
}
