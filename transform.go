package main

import (
	"strconv"
	"strings"
	"unicode"
)

func capitalize(word string) string {
	words := strings.Fields(word)
	for i, char := range words {
		runes := []rune(strings.ToLower(char))
		runes[0] = unicode.ToUpper(runes[0])
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}


func ConvertNumbers(word string) string {
	words := strings.Fields(word)
	for i := 0; i < len(words); i++ {
		switch words[i] {
		case "(hex)":
			if i > 0 {
				num, err := strconv.ParseInt(words[i-1], 16, 64)
				if err == nil {
					words[i-1] = strconv.FormatInt(num, 10)
				}
			}
			words = append(words[:i], words[i+1:]...)
			i--

		case "(bin)":
			if i > 0 {
				num, err := strconv.ParseInt(words[i-1], 2, 64)
				if err == nil {
					words[i-1] = strconv.FormatInt(num, 10)
				}
			}
			words = append(words[:i], words[i+1:]...)
			i--

		}

	}

	return strings.Join(words, " ")
}

func ApplyMarkers(word string) string {
	words := strings.Fields(word)
	for i := 0; i < len(words); i++ {
		switch words[i] {
		case "(low)":
			if i > 0 {
				words[i-1] = strings.ToLower(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--

		case "(up)":
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--
		case "(cap)":
			if i > 0 {
				capitalize(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--

		}

	}

	return strings.Join(words, " ")
}
