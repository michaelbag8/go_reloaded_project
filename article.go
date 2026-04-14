package main

import (
	"strings"
)

func fixArticle(word string) string {
	words := strings.Fields(word)
	vowels := "aeiuohAEIUOH"
	for i := 0; i < len(words); i++ {
		if (words[i] == "A" || words[i] == "a") && strings.ContainsRune(vowels, rune(words[i+1][0])) {
			if words[i] == "A" {
				words[i] = "An"
			} else {
				words[i] = "an"
			}
		}
	}
	return strings.Join(words, " ")
}
