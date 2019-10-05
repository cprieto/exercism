package acronym

import (
	"regexp"
	"strings"
)

var words = regexp.MustCompile(`(?m)([a-zA-Z']+(?:_[a-zA-Z]+)*)`)

func Abbreviate(s string) string {
	var builder strings.Builder
	for _, word := range words.FindAllString(s, -1) {
		builder.WriteRune([]rune(strings.ToUpper(word))[0])
	}
	return builder.String()
}
