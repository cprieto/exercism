package isogram

import (
	"strings"
	"unicode"
)


func IsIsogram(input string) bool {
	letters := make(map[rune]int)
	for idx, letter := range []rune(strings.ToUpper(input)) {
		if _, ok := letters[letter]; ok && unicode.IsLetter(letter)  {
			return false
		}
		letters[letter] = idx
	}
	return true
}
