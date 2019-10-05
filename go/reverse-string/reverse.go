package reverse

import "strings"

func Reverse(input string) string {
	entry := []rune(input)
	var output strings.Builder

	for n := len(entry) - 1; n >= 0; n-- {
		output.WriteRune(entry[n])
	}

	return output.String()
}
