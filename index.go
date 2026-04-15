package main
import(
	"strings"
	"strconv"
	"unicode"
)
func capitalize(word string)string{
	words := strings.Fields(word)
	for i, currentWord := range words{
		runes := []rune(strings.ToLower(currentWord))
		if len(runes) > 0{
			runes[0] = unicode.ToUpper(runes[0])
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}

func Case(s string)string{
	s = strings.ReplaceAll(s, ", ",",")
	word := strings.Fields(s)
	for i := 0; i < len(word); i++ {
		if word[i] == "(up)" && i > 0{
			word[i-1] = strings.ToUpper(word[i-1])
			word = append(word[:i], word[i+1:]...)
			i--
		}
		if word[i] == "(low)" && i > 0{
			word[i-1] = strings.ToLower(word[i-1])
			word = append(word[:i], word[i+1:]...)
			i--
		}
		if word[i] == "(cap)" && i > 0{
			word[i-1] = capitalize(word[i-1])
			word = append(word[:i], word[i+1:]...)
			i--
		}
		if strings.HasPrefix(word[i],"(") && strings.HasSuffix(word[i], ")"){
			word[i] = strings.Trim(word[i], "()")
			result := strings.Split(word[i], ",")
			num, err := strconv.Atoi(result[1])
			if err == nil{
				start := i - num
				if start < 0{
					start = 0
				}
				for j := start; j < i; j++{
					switch result[0] {
					case "up":
						word[j] = strings.ToUpper(word[j])
					case "low":
						word[j] = strings.ToLower(word[j])
					case "cap":
						word[j] = capitalize(word[j])
					}
					
				}
			}
			word = append(word[:i], word[i+1:]...)
			i--
		}
}
return strings.Join(word, " ")
}