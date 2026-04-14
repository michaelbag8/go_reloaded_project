package main

import (
	"strconv"
	"strings"
	"unicode"
)

func Capitalize(word string) string {
	words := strings.Fields(word)
	for i, currentWord := range words {
		runes := []rune(strings.ToLower(currentWord))
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}
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

func markerDetector(word, mode string) string {
	switch mode {
	case "low":
		return strings.ToLower(word)
	case "up":
		return strings.ToUpper(word)
	case "cap":
		return Capitalize(word)
	}
	return word
}

func applyLastN(word string) string {
	word = strings.ReplaceAll(word, ", ", ",")
	words := strings.Fields(word)

	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(") && strings.HasSuffix(words[i], ")") {
			inner := strings.Trim(words[i], "()")
			extract := strings.Split(inner, ",")
			mode := extract[0]

			if len(extract) == 1 {
				if i > 0 {
					words[i-1] = markerDetector(words[i-1], mode)
				}
			} else {
				num, err := strconv.Atoi(extract[1])
				if err == nil {
					start := i - num
					if start < 0 {
						start = 0
					}
					for j := start; j < i; j++ {
						words[j] = markerDetector(words[j], mode)
					}
				}

			}
			words = append(words[:i], words[i+1:]...)
			i--

		}
	}
	return strings.Join(words, " ")
}
