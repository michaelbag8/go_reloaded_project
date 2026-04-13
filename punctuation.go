package main

import (
	"strings"
)

func fixPunctuation(text string) string {
	punctuation := []string{".", ",", "!", "?", ":", ";"}
	for _, p := range punctuation {
		text = strings.ReplaceAll(text, " "+p, p)
	}
	return text
}

func fixQuotes(text string) string {
	text = strings.ReplaceAll(text, "' ", "'")
	text = strings.ReplaceAll(text, " '", "'")
	return text
}
